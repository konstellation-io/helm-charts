package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/konstellation-io/kdl-server/app/api/entity"
	"github.com/konstellation-io/kdl-server/app/api/infrastructure/giteaservice"
	"github.com/konstellation-io/kdl-server/app/api/pkg/clock"
	"github.com/konstellation-io/kdl-server/app/api/pkg/k8s"
	"github.com/konstellation-io/kdl-server/app/api/pkg/logging"
	"github.com/konstellation-io/kdl-server/app/api/pkg/sshhelper"
)

// Interactor object implements the UseCase interface.
type Interactor struct {
	logger       logging.Logger
	repo         Repository
	sshGenerator sshhelper.SSHKeyGenerator
	clock        clock.Clock
	giteaService giteaservice.GiteaClient
	k8sClient    k8s.K8sClient
}

// NewInteractor factory function.
func NewInteractor(
	logger logging.Logger,
	repo Repository,
	sshGenerator sshhelper.SSHKeyGenerator,
	c clock.Clock,
	giteaService giteaservice.GiteaClient,
	k8sClient k8s.K8sClient,
) *Interactor {
	return &Interactor{
		logger:       logger,
		repo:         repo,
		sshGenerator: sshGenerator,
		clock:        c,
		giteaService: giteaService,
		k8sClient:    k8sClient,
	}
}

// Create add a new user to the server.
// - If the user already exists (email and username must be unique) returns entity.ErrDuplicatedUser.
// - Generates a new SSH public/private keys.
// - Creates the user into Gitea.
// - Adds the public SSH key to the user in Gitea.
// - Add the user to the KDL team.
// - Stores the user and ssh keys into the DB.
// - Creates a new secret in Kubernetes with the generated SSH keys.
func (i Interactor) Create(ctx context.Context, email, username, password string, accessLevel entity.AccessLevel) (entity.User, error) {
	i.logger.Infof("Creating user \"%s\" with email \"%s\"", username, email)

	// Check if the user already exists
	_, err := i.repo.GetByUsername(ctx, username)
	if err == nil {
		return entity.User{}, entity.ErrDuplicatedUser
	}

	if !errors.Is(err, entity.ErrUserNotFound) {
		return entity.User{}, err
	}

	_, err = i.repo.GetByEmail(ctx, email)
	if err == nil {
		return entity.User{}, entity.ErrDuplicatedUser
	}

	if !errors.Is(err, entity.ErrUserNotFound) {
		return entity.User{}, err
	}

	// Create SSH public and private keys
	keys, err := i.sshGenerator.NewKeys()
	if err != nil {
		return entity.User{}, err
	}

	// Creates the user into Gitea.
	err = i.giteaService.CreateUser(email, username, password)
	if err != nil {
		return entity.User{}, err
	}

	// Adds the public SSH key to the user in Gitea.
	err = i.giteaService.AddSSHKey(username, keys.Public)
	if err != nil {
		return entity.User{}, err
	}

	// Add the user to the KDL team.
	err = i.giteaService.AddTeamMember(username, accessLevel)
	if err != nil {
		return entity.User{}, err
	}

	// Stores the user and ssh keys into the DB.
	user := entity.User{
		Username:     username,
		Email:        email,
		AccessLevel:  accessLevel,
		CreationDate: i.clock.Now(),
		SSHKey:       keys,
	}

	insertedID, err := i.repo.Create(ctx, user)
	if err != nil {
		return entity.User{}, err
	}

	i.logger.Infof("The user \"%s\" (%s) was created with ID \"%s\"", user.Username, user.Email, insertedID)

	secretName := fmt.Sprintf("%s-ssh-keys", user.Username)
	err = i.k8sClient.CreateSecret(secretName, map[string]string{
		"KDL_USER_PUBLIC_SSH_KEY":  keys.Public,
		"KDL_USER_PRIVATE_SSH_KEY": keys.Private,
	})

	if err != nil {
		return entity.User{}, err
	}

	return i.repo.Get(ctx, insertedID)
}

// FindAll returns all users existing in the server.
func (i Interactor) FindAll(ctx context.Context) ([]entity.User, error) {
	i.logger.Info("Finding all users in the server")
	return i.repo.FindAll(ctx)
}

// GetByEmail returns the user with the desired email or returns entity.ErrUserNotFound if the user doesn't exist.
func (i Interactor) GetByEmail(ctx context.Context, email string) (entity.User, error) {
	i.logger.Infof("Getting user by email \"%s\"", email)
	return i.repo.GetByEmail(ctx, email)
}