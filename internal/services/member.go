package services

import (
	"context"
	"fmt"

	member "github.com/goclean/internal/model"
	"github.com/goclean/internal/repository"
)

type memberService struct {
	repo repository.Repo
}

// NewMemberService initiate new Member Service based on Contract
func NewMemberService(repo repository.Repo) Service {
	return &memberService{
		repo: repo,
	}
}

func (m *memberService) Index(ctx context.Context, query map[string]interface{}, page, size int) (members []member.Member, err error) {
	fmt.Println("service")
	// write business use case here
	return m.repo.Index(ctx, query, page, size)
}

func (m *memberService) Insert(ctx context.Context, data member.Member) (member.Member, error) {
	return m.repo.Insert(ctx, data)
}
