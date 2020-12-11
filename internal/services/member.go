package services

import (
	"clean-arch/model"
)

type memberService struct {
	repo model.MemberRepo
}

func NewMemberService(repo model.MemberRepo) model.MemberService {
	return &NewMemberService{
		repo: repo,
	}
}