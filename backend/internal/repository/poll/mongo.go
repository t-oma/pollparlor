package poll

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"pollparlor/internal/domain"
)

// MongoRepo is a repository for polls in mongoDB
type MongoRepo struct {
	pollsCol  *mongo.Collection
	itemsCol  *mongo.Collection
	opTimeout time.Duration
}

var _ domain.PollRepository = (*MongoRepo)(nil)

// NewMongoRepo creates a new repository for polls in mongoDB
func NewMongoRepo(db *mongo.Database, opTimeout time.Duration) *MongoRepo {
	return &MongoRepo{
		pollsCol:  db.Collection("polls"),
		itemsCol:  db.Collection("pollItems"),
		opTimeout: opTimeout,
	}
}

// List returns all polls in the repository
func (r *MongoRepo) List(limit, skip int64) ([]domain.Poll, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.opTimeout)
	defer cancel()

	opts := options.Find()
	if limit != 0 {
		opts.SetLimit(limit)
	}
	if skip != 0 {
		opts.SetSkip(skip)
	}

	cur, err := r.pollsCol.Find(ctx, bson.D{}, opts)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return []domain.Poll{}, nil
	}
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var out []domain.Poll
	if err := cur.All(ctx, &out); err != nil {
		return nil, err
	}

	return out, nil
}

// GetByID returns a poll by its ID
func (r *MongoRepo) GetByID(id bson.ObjectID) (*domain.Poll, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.opTimeout)
	defer cancel()

	var p domain.Poll
	err := r.pollsCol.FindOne(ctx, bson.M{"id": id}).Decode(&p)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

// Create adds a poll to the repository
func (r *MongoRepo) Create(p domain.Poll) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.opTimeout)
	defer cancel()

	_, err := r.pollsCol.InsertOne(ctx, p)
	return err
}

// GetItems returns all items in the repository
func (r *MongoRepo) GetItems(pollID bson.ObjectID) ([]domain.Candidate, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.opTimeout)
	defer cancel()

	var items domain.PollItems
	err := r.itemsCol.FindOne(ctx, bson.M{"pollId": pollID}).Decode(&items)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return []domain.Candidate{}, nil
	}
	if err != nil {
		return nil, err
	}

	return items.Cands, nil
}

// CreateItems adds items to the repository
func (r *MongoRepo) CreateItems(pollID bson.ObjectID, items []domain.Candidate) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.opTimeout)
	defer cancel()

	_, err := r.itemsCol.InsertOne(ctx, domain.PollItems{
		ID:     bson.NewObjectID(),
		PollID: pollID,
		Cands:  items,
	})
	return err
}

// GetPairs returns all pairs in the repository
func (r *MongoRepo) GetPairs(pollID bson.ObjectID) ([]domain.PollPair, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.opTimeout)
	defer cancel()

	var items domain.PollItems
	err := r.itemsCol.FindOne(ctx, bson.M{"pollId": pollID}).Decode(&items)
	if err != nil {
		return nil, err
	}

	// TODO: shuffle items and return pairs of candidates

	return nil, nil
}
