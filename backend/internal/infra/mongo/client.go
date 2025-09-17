// Package mongox provides mongoDB client and database
package mongox

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// NewClient creates a new mongoDB client
func NewClient(ctx context.Context, uri string) (*mongo.Client, error) {
	return mongo.Connect(options.Client().ApplyURI(uri))
}

// NewDatabase creates a new mongoDB database
func NewDatabase(client *mongo.Client, name string) *mongo.Database {
	return client.Database(name)
}
