package member

import "context"

// Member struct
type Member struct {
	Name    string
	Email   string
	Phone   string
	Address string
}

// MemberService represents Member use case contract
type MemberService interface {
	Index(ctx context.Context, query map[string]interface{}, page, size int) ([]Member, map[string]int, error)
}

// MemberRepo represents member repo contract
type MemberRepo interface {
	Index(ctx context.Context, query map[string]interface{}, page, size int) ([]Member, error)
}
