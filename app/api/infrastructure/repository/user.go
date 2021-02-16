package repository

import (
	"context"
	"errors"
	"time"

	"github.com/konstellation-io/kdl-server/app/api/usecase/user"

	"github.com/konstellation-io/kdl-server/app/api/pkg/mongodb"

	"github.com/konstellation-io/kdl-server/app/api/entity"
	"github.com/konstellation-io/kdl-server/app/api/pkg/logging"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	userCollName         = "users"
	ensureIndexesTimeout = 20 * time.Second
)

type userDTO struct {
	ID            primitive.ObjectID `bson:"_id"`
	Username      string             `bson:"username"`
	Email         string             `bson:"email"`
	CreationDate  time.Time          `bson:"creation_date"`
	AccessLevel   string             `bson:"access_level"`
	PublicSSHKey  string             `bson:"public_ssh_key"`
	PrivateSSHKey string             `bson:"private_ssh_key"`
}

type userMongoDBRepo struct {
	logger     logging.Logger
	collection *mongo.Collection
}

// NewUserMongoDBRepo implements user.Repository interface.
func NewUserMongoDBRepo(logger logging.Logger, client *mongo.Client, dbName string) user.Repository {
	collection := client.Database(dbName).Collection(userCollName)
	return &userMongoDBRepo{logger, collection}
}

// EnsureIndexes creates if not exists the indexes for the user collection.
func (m *userMongoDBRepo) EnsureIndexes() error {
	m.logger.Infof("Creating index for \"%s\" collection...", userCollName)

	ctx, cancel := context.WithTimeout(context.Background(), ensureIndexesTimeout)
	defer cancel()

	unique := true
	result, err := m.collection.Indexes().CreateMany(ctx, []mongo.IndexModel{
		{
			Keys: bson.M{"email": 1},
			Options: &options.IndexOptions{
				Unique: &unique,
			},
		},
		{
			Keys: bson.M{"username": 1},
			Options: &options.IndexOptions{
				Unique: &unique,
			},
		},
	})

	m.logger.Infof("Indexes %s created for \"%s\" collection", result, userCollName)

	return err
}

// Get retrieves the user using the identifier.
func (m *userMongoDBRepo) Get(ctx context.Context, id string) (entity.User, error) {
	idFromHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return entity.User{}, err
	}

	return m.findOne(ctx, bson.M{"_id": idFromHex})
}

// GetByUsername retrieves the user using their user name.
func (m *userMongoDBRepo) GetByUsername(ctx context.Context, username string) (entity.User, error) {
	return m.findOne(ctx, bson.M{"username": username})
}

// GetByUsername retrieves the user using their user email.
func (m *userMongoDBRepo) GetByEmail(ctx context.Context, email string) (entity.User, error) {
	return m.findOne(ctx, bson.M{"email": email})
}

// FindAll retrieves all the existing users.
func (m *userMongoDBRepo) FindAll(ctx context.Context) ([]entity.User, error) {
	return m.find(ctx, bson.M{})
}

// Create inserts into the database a new entity.
func (m *userMongoDBRepo) Create(ctx context.Context, u entity.User) (string, error) {
	m.logger.Debugf("Inserting a new user \"%s\" into %s collection...", u.Email, userCollName)

	dto, err := m.entityToDTO(u)
	if err != nil {
		return "", err
	}

	dto.ID = primitive.NewObjectID()

	result, err := m.collection.InsertOne(ctx, dto)
	if err != nil {
		if mongodb.IsDuplKeyError(err) {
			return "", entity.ErrDuplicatedUser
		}

		return "", err
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (m *userMongoDBRepo) findOne(ctx context.Context, filters bson.M) (entity.User, error) {
	m.logger.Debugf("Finding one user by \"%s\" from database...", filters)

	dto := userDTO{}

	err := m.collection.FindOne(ctx, filters).Decode(&dto)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return entity.User{}, entity.ErrUserNotFound
	}

	return m.dtoToEntity(dto), err
}

func (m *userMongoDBRepo) find(ctx context.Context, filters bson.M) ([]entity.User, error) {
	m.logger.Debugf("Finding users with filters \"%s\"...", filters)

	var dtos []userDTO

	err := mongodb.Find(ctx, filters, m.collection, &dtos)
	if err != nil {
		return nil, err
	}

	return m.dtosToEntities(dtos), nil
}

func (m *userMongoDBRepo) entityToDTO(u entity.User) (userDTO, error) {
	dto := userDTO{
		Username:      u.Username,
		Email:         u.Email,
		AccessLevel:   string(u.AccessLevel),
		PrivateSSHKey: u.SSHKey.Private,
		PublicSSHKey:  u.SSHKey.Public,
		CreationDate:  u.CreationDate,
	}

	if u.ID != "" {
		idFromHex, err := primitive.ObjectIDFromHex(u.ID)
		if err != nil {
			return dto, err
		}

		dto.ID = idFromHex
	}

	return dto, nil
}

func (m *userMongoDBRepo) dtoToEntity(dto userDTO) entity.User {
	return entity.User{
		ID:           dto.ID.Hex(),
		Username:     dto.Username,
		Email:        dto.Email,
		AccessLevel:  entity.AccessLevel(dto.AccessLevel),
		CreationDate: dto.CreationDate,
		SSHKey: entity.SSHKey{
			Public:  dto.PublicSSHKey,
			Private: dto.PrivateSSHKey,
		},
	}
}

func (m *userMongoDBRepo) dtosToEntities(dtos []userDTO) []entity.User {
	result := make([]entity.User, len(dtos))

	for i, dto := range dtos {
		result[i] = m.dtoToEntity(dto)
	}

	return result
}