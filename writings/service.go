package writings

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// third column is the bson tag, in this case the name of the property in mongo
type Writing struct {
	ID               primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Group            string             `json:"group" bson:"group"`
	Title            string             `json:"title" bson:"title"`
	Content          string             `json:"content" bson:"content"`
	Copyright        string             `json:"copyright" bson:"copyright"`
	Acknowledgements string             `json:"acknowledgements" bson:"acknowledgements"`
	Kind             string             `json:"kind" bson:"kind"`
	Genre            string             `json:"genre" bson:"genre"`
	Tag              string             `json:"tag" bson:"tag"`
}

type NewWriting struct {
	Group            string `bson:"group"`
	Title            string `bson:"title"`
	Content          string `bson:"content"`
	Copyright        string `bson:"copyright"`
	Acknowledgements string `bson:"acknowledgements"`
	Kind             string `bson:"kind"`
	Genre            string `bson:"genre"`
	Tag              string `bson:"tag"`
}

// We could require that a service took an instance of the mongo db
// But abstracting to requiring an interface means that we can work with a variety of DB implementations
// So long as they have the create method
type WritingsRepo interface {
	Index(ctx context.Context) ([]Writing, error)
	Create(ctx context.Context, entry NewWriting) (*Writing, error)
	Read(ctx context.Context, id string) (*Writing, error)
}

// Return a new instance of Service
func New(db WritingsRepo) Service {
	return Service{db}
}

// Service struct takes a DB instance
type Service struct {
	db WritingsRepo
}

func (s *Service) Index(ctx context.Context) ([]Writing, error) {
	return s.db.Index(ctx)
}

func (s *Service) Create(ctx context.Context, entry NewWriting) (*Writing, error) {
	return s.db.Create(ctx, entry)
}

func (s *Service) Show(ctx context.Context, id string) (*Writing, error) {
	return s.db.Read(ctx, id)
}
