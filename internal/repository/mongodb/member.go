package mongodb

import (
	"context"

	member "github.com/goclean/internal/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoMemberRepo struct {
	client *mongo.Collection
}

func NewMemberRepo(client *mongo.Collection) member.Repo {
	return &mongoMemberRepo{client}
}

func (m *mongoMemberRepo) Index(ctx context.Context, query map[string]interface{}, page, size int) (members []member.Member, err error) {
	cur, err := m.client.Find(ctx, nil)
	if err != nil {
		return members, err
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result member.Member
		cur.Decode(&result)
		members = append(members, result)
	}
	return members, nil
}

func (m *mongoMemberRepo) Insert(ctx context.Context, data member.Member) (member.Member, error) {
	newID, err := m.client.InsertOne(ctx, data)
	if err != nil {
		return member.Member{}, err
	}
	singRes := m.client.FindOne(ctx, newID)
	var newMember member.Member
	err = singRes.Decode(&newMember)
	if err != nil {
		return member.Member{}, err
	}
	return newMember, nil
}
