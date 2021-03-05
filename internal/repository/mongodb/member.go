package mongodb

import (
	"context"
	"fmt"

	member "github.com/goclean/internal/model"
	"github.com/goclean/internal/repository"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoMemberRepo struct {
	client *mongo.Collection
}

func NewMemberRepo(client *mongo.Collection) repository.Repo {
	return &mongoMemberRepo{client}
}

func (m *mongoMemberRepo) Index(ctx context.Context, query map[string]interface{}, page, size int) (members []member.Member, err error) {
	fmt.Println("repo")
	cur, err := m.client.Find(ctx, query)
	if err != nil {
		return members, err
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result member.Member
		err := cur.Decode(&result)
		if err != nil {
			log.Error(err.Error())
		}
		members = append(members, result)
	}
	fmt.Println(members)
	return members, nil
}

func (m *mongoMemberRepo) Insert(ctx context.Context, data member.Member) (member.Member, error) {
	newID, err := m.client.InsertOne(ctx, data)
	if err != nil {
		return member.Member{}, err
	}
	singRes := m.client.FindOne(ctx, newID.InsertedID)
	var newMember member.Member
	err = singRes.Decode(&newMember)
	if err != nil {
		return member.Member{}, err
	}
	return newMember, nil
}
