package member

import "context"

// Member struct
type Member struct {
	Name    string
	Email   string
	Phone   string
	Address string
}

// Service represents Member use case contract
type Service interface {
	Index(ctx context.Context, query map[string]interface{}, page, size int) ([]Member, error)
	Insert(ctx context.Context, data Member) (Member, error)
}

// Repo represents member repo contract
type Repo interface {
	Index(ctx context.Context, query map[string]interface{}, page, size int) ([]Member, error)
	Insert(ctx context.Context, data Member) (Member, error)
}
