package poll

import (
	"context"
	"errors"
	"math/rand/v2"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
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
func (r *MongoRepo) List(limit, skip int64) ([]domain.PollWithAuthor, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.opTimeout)
	defer cancel()

	pipeline := mongo.Pipeline{
		// {{Key: "$sort", Value: bson.D{{"createdAt", -1}}}},
	}
	if skip > 0 {
		pipeline = append(pipeline, bson.D{{Key: "$skip", Value: skip}})
	}
	if limit > 0 {
		pipeline = append(pipeline, bson.D{{Key: "$limit", Value: limit}})
	}
	pipeline = append(pipeline,
		bson.D{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "users"},
			{Key: "localField", Value: "authorId"},
			{Key: "foreignField", Value: "uuid"},
			{Key: "as", Value: "author"},
		}}},
		bson.D{{Key: "$unwind", Value: bson.D{
			{Key: "path", Value: "$author"},
			{Key: "preserveNullAndEmptyArrays", Value: true},
		}}},
		bson.D{{Key: "$project", Value: bson.D{
			{Key: "_id", Value: 1},
			{Key: "title", Value: 1},
			// {Key: "authorId", Value: 1},
			{Key: "likes", Value: 1},
			{Key: "createdAt", Value: 1},
			{Key: "updatedAt", Value: 1},
			{Key: "author", Value: 1},
		}}},
	)

	cur, err := r.pollsCol.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var out []domain.PollWithAuthor
	if err := cur.All(ctx, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// GetByID returns a poll by its ID
func (r *MongoRepo) GetByID(id bson.ObjectID) (*domain.PollWithAuthor, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.opTimeout)
	defer cancel()

	pipeline := mongo.Pipeline{
		{{
			Key: "$match", Value: bson.D{
				{Key: "_id", Value: id},
			},
		}},
		{{
			Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "users"},
				{Key: "localField", Value: "authorId"},
				{Key: "foreignField", Value: "uuid"},
				{Key: "as", Value: "author"},
			},
		}},
		{{
			Key: "$unwind", Value: bson.D{
				{Key: "path", Value: "$author"},
				{Key: "preserveNullAndEmptyArrays", Value: true},
			},
		}},
		{{
			Key: "$project", Value: bson.D{
				{Key: "_id", Value: 1},
				{Key: "title", Value: 1},
				{Key: "authorId", Value: 1},
				{Key: "likes", Value: 1},
				{Key: "createdAt", Value: 1},
				{Key: "updatedAt", Value: 1},
				{Key: "author", Value: 1},
			},
		}},
	}

	cur, err := r.pollsCol.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var out []domain.PollWithAuthor
	if err := cur.All(ctx, &out); err != nil {
		return nil, err
	}

	return &out[0], nil
}

// Create adds a poll to the repository
func (r *MongoRepo) Create(p domain.Poll) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.opTimeout)
	defer cancel()

	_, err := r.pollsCol.InsertOne(ctx, p)
	return err
}

// Delete deletes a poll from the repository
func (r *MongoRepo) Delete(id bson.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.opTimeout)
	defer cancel()

	_, err := r.pollsCol.DeleteOne(ctx, bson.M{"_id": id})
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
	if errors.Is(err, mongo.ErrNoDocuments) {
		return []domain.PollPair{}, nil
	}
	if err != nil {
		return nil, err
	}

	cands := items.Cands
	rand.Shuffle(len(cands), func(i, j int) {
		cands[i], cands[j] = cands[j], cands[i]
	})

	pairs := make([]domain.PollPair, 0, len(cands)/2)
	for i := 0; i < len(cands)-1; i += 2 {
		pairs = append(pairs, domain.PollPair{
			Left:  cands[i],
			Right: cands[i+1],
		})
	}

	return pairs, nil
}

// DeleteItems deletes items from the repository
func (r *MongoRepo) DeleteItems(pollID bson.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.opTimeout)
	defer cancel()

	_, err := r.itemsCol.DeleteOne(ctx, bson.M{"pollId": pollID})
	return err
}
