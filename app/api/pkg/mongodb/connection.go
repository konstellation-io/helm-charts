package mongodb

import (
	"context"
	"time"

	"github.com/konstellation-io/kdl-server/app/api/pkg/logging"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const timeout = 20 * time.Second

// MongoDB will manage the connection with the database.
type MongoDB struct {
	logger logging.Logger
	client *mongo.Client
}

// NewMongoDB is a constructor function.
func NewMongoDB(logger logging.Logger) *MongoDB {
	return &MongoDB{
		logger,
		nil,
	}
}

// Connect open a database connection and check if it is connecting using ping.
func (m *MongoDB) Connect(uri string) (*mongo.Client, error) {
	m.logger.Info("MongoDB connecting...")

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	m.logger.Info("MongoDB ping...")

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	m.logger.Info("MongoDB connected")
	m.client = client

	return client, nil
}

// Disconnect closes the connection with the database.
func (m *MongoDB) Disconnect() {
	m.logger.Info("MongoDB disconnecting...")

	if m.client == nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	err := m.client.Disconnect(ctx)
	if err != nil {
		m.logger.Errorf("Error disconnecting from MongoDB: %s", err)
		return
	}

	m.logger.Info("Connection to MongoDB closed.")
}
