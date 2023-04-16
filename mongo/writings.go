package mongo

import (
	"blog-api/writings"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewWritingsRepo(db *mongo.Database) WritingsRepo {
	return WritingsRepo{db}
}

type WritingsRepo struct {
	db *mongo.Database
}

func (wr *WritingsRepo) Index(ctx context.Context) ([]writings.Writing, error) {
	fmt.Printf("Writings -> INDEX.")

	cur, err := wr.db.Collection("writings").Find(ctx, bson.D{})
	if err != nil {
		fmt.Println("Finding all documents ERROR:", err)
		defer cur.Close(ctx)
	}

	result := make([]writings.Writing, 0, cur.RemainingBatchLength())

	var i writings.Writing
	for cur.Next(ctx) {
		// Declare a result BSON object
		if err := cur.Decode(&i); err != nil {
			log.Fatal(err)
		}

		result = append(result, i)
	}

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (wr *WritingsRepo) Create(ctx context.Context, entry writings.NewWriting) (*writings.Writing, error) {
	fmt.Printf("Writings -> CREATE.")

	res, err := wr.db.Collection("writings").InsertOne(ctx, entry)
	if err != nil {
		return nil, err
	}

	filter := bson.M{
		"_id": res.InsertedID,
	}

	var result writings.Writing
	err = wr.db.Collection("writings").FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (wr *WritingsRepo) Read(ctx context.Context, id string) (*writings.Writing, error) {
	fmt.Printf("Writings -> READ. ID:: %+v \n", id)

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{
		"_id": objID,
	}

	var result writings.Writing
	err = wr.db.Collection("writings").FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
