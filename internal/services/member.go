package services

import (
	"context"

	member "github.com/goclean/internal/model"
)

type memberService struct {
	repo member.Repo
}

// NewMemberService initiate new Member Service based on Contract
func NewMemberService(repo member.Repo) member.Service {
	return &memberService{
		repo: repo,
	}
}

func (m *memberService) Index(ctx context.Context, query map[string]interface{}, page, size int) (members []member.Member, err error) {
	return m.repo.Index(ctx, query, page, size)
}

func (m *memberService) Insert(ctx context.Context, data member.Member) (member.Member, error) {
	return m.repo.Insert(ctx, data)
}
