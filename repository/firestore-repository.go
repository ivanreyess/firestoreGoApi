package repository

import (
	"context"
	"log"

	"../entity"
	"cloud.google.com/go/firestore"
	it "google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

type repo struct{}

const (
	projectID      string = "fir-6e15a"
	collectionName string = "posts"
)

//NewPostRepository return an instance of repo structure
func NewFirestoreRepository() PostRepository {
	return &repo{}
}

func (*repo) Save(post *entity.Post) (*entity.Post, error) {
	opt := option.WithCredentialsFile("/home/ivan/ServiceAccounts/fir-6e15a-firebase-adminsdk-vxoaf-4e478fd55f.json")
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID, opt)

	if err != nil {
		log.Printf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}

	defer client.Close()

	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"ID":    post.ID,
		"Title": post.Title,
		"Text":  post.Text,
	})

	if err != nil {
		log.Printf("Failed to create a new post: %v ", err)
		return nil, err
	}

	return post, nil

}

func (*repo) FindAll() ([]entity.Post, error) {
	opt := option.WithCredentialsFile("/home/ivan/ServiceAccounts/fir-6e15a-firebase-adminsdk-vxoaf-4e478fd55f.json")
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID, opt)

	if err != nil {
		log.Printf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}

	defer client.Close()

	var posts []entity.Post
	iterator := client.Collection(collectionName).Documents(ctx)

	for {
		doc, err := iterator.Next()

		if err == it.Done {
			break
		}

		if err != nil {
			log.Printf("Failed to create a Firestore Client: %v", err)
			return nil, err
		}

		post := entity.Post{
			ID:    doc.Data()["ID"].(int64),
			Title: doc.Data()["Title"].(string),
			Text:  doc.Data()["Text"].(string),
		}

		posts = append(posts, post)
	}

	return posts, nil

}
