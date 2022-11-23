package user

//go:generate mockgen -source=${GOFILE} -destination=mocks_${GOFILE} -package=${GOPACKAGE}

import (
	"context"
	"time"

	"github.com/konstellation-io/kdl-server/app/api/entity"
)

// Repository interface to retrieve and persists users.
type Repository interface {
	EnsureIndexes() error
	Get(ctx context.Context, id string) (entity.User, error)
	GetByUsername(ctx context.Context, username string) (entity.User, error)
	GetByEmail(ctx context.Context, email string) (entity.User, error)
	Create(ctx context.Context, user entity.User) (string, error)
	UpdateAccessLevel(ctx context.Context, userIDs []string, level entity.AccessLevel) error
	UpdateSSHKey(ctx context.Context, username string, SSHKey entity.SSHKey) error
	FindAll(ctx context.Context, includeDeleted bool) ([]entity.User, error)
	FindByIDs(ctx context.Context, userIDs []string) ([]entity.User, error)
	UpdateEmail(ctx context.Context, userID, email string) error
	UpdateDeleted(ctx context.Context, userID string, deleted bool) error
}

// UseCase interface to manage all operations related with users.
type UseCase interface {
	Create(ctx context.Context, email, username string, accessLevel entity.AccessLevel) (entity.User, error)
	CreateAdminUser(username, email string) error
	UpdateAccessLevel(ctx context.Context, userIds []string, level entity.AccessLevel) ([]entity.User, error)
	FindAll(ctx context.Context) ([]entity.User, error)
	GetByUsername(ctx context.Context, username string) (entity.User, error)
	StartTools(ctx context.Context, username string, runtimeID *string, capabilitiesId *string) (entity.User, error)
	StopTools(ctx context.Context, username string) (entity.User, error)
	AreToolsRunning(ctx context.Context, username string) (bool, error)
	IsKubeconfigActive() bool
	FindByIDs(ctx context.Context, userIDs []string) ([]entity.User, error)
	GetByID(ctx context.Context, userID string) (entity.User, error)
	RegenerateSSHKeys(ctx context.Context, user entity.User) (entity.User, error)
	ScheduleUsersSyncJob(interval time.Duration) error
	RunSyncUsersCronJob()
	GetKubeconfig(ctx context.Context, username string) (string, error)
	SynchronizeServiceAccountsForUsers() error
}
