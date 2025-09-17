// Package user provides a repository for users
package user

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"pollparlor/internal/domain"
)

// MongoRepo is a repository for users in mongoDB
type MongoRepo struct {
	col       *mongo.Collection
	opTimeout time.Duration
}

var _ domain.UserRepository = (*MongoRepo)(nil)

// NewMongoRepo creates a new repository for users in mongoDB
func NewMongoRepo(db *mongo.Database, opTimeout time.Duration) *MongoRepo {
	return &MongoRepo{
		col:       db.Collection("users"),
		opTimeout: opTimeout,
	}
}

// List returns all users in the repository
func (r *MongoRepo) List(limit, skip int64) ([]domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.opTimeout)
	defer cancel()

	opts := options.Find()
	if limit != 0 {
		opts.SetLimit(limit)
	}
	if skip != 0 {
		opts.SetSkip(skip)
	}

	cur, err := r.col.Find(ctx, bson.D{}, opts)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var out []domain.User
	if err := cur.All(ctx, &out); err != nil {
		return nil, err
	}

	return out, nil
}

// GetByID returns a user by its ID
func (r *MongoRepo) GetByID(id string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.opTimeout)
	defer cancel()

	var u domain.User
	err := r.col.FindOne(ctx, bson.M{"id": id}).Decode(&u)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	return &u, nil
}

// Create adds a user to the repository
func (r *MongoRepo) Create(u domain.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.opTimeout)
	defer cancel()

	_, err := r.col.InsertOne(ctx, u)
	return err
}
