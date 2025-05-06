package infrastructure

import (
	"context"

	"curs1_boilerplate/model"

	"github.com/google/uuid"
)

type BidRepository interface {
	GetBidsForPost(ctx context.Context, postId uuid.UUID) ([]model.Bid, error)
}
