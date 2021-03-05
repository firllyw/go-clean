package services

import (
	"context"

	member "github.com/goclean/internal/model"
)

// Service represents Member use case contract
type Service interface {
	Index(ctx context.Context, query map[string]interface{}, page, size int) ([]member.Member, error)
	Insert(ctx context.Context, data member.Member) (member.Member, error)
	Update(ctx context.Context, id string, data member.Member) (member.Member, error)
}
