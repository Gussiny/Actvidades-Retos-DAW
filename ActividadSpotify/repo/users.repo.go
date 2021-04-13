package repo

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/Gussiny/apiRest/entity"
	"google.golang.org/api/iterator"
)

type UserRepository interface {
	FindAll() ([]entity.User, error)
}

type repo struct{}

func NewUserRepository() UserRepository {
	return &repo{}
}

const (
	projectId      string = "spotifypirata-b35dd"
	collectionName string = "users"
)

func (*repo) FindAll() ([]entity.User, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}

	defer client.Close()

	var users []entity.User
	iter := client.Collection(collectionName).Documents(ctx)
	for {
		doc, err := iter.Next()

		if err == iterator.Done {
			break
		}

		if err != nil {
			log.Fatalf("Failed to iterate the list of users : %v ", err)
			return nil, err
		}
		user := entity.User{
			Username: doc.Data()["username"].(string),
		}
		users = append(users, user)

	}

	return users, nil

}
