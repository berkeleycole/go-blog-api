package mockdb

import (
	"blog-api/writings"
	"context"
	"errors"
)

type WritingsRepo struct{}

func New() WritingsRepo {
	return WritingsRepo{}
}

func (mr *WritingsRepo) Create(ctx context.Context) (*writings.Writing, error) {
	return nil, errors.New("unimplemented")
}
