package post

import (
	"context"
	"go-api-test/model"
	"go-api-test/persist/db"
	"google.golang.org/api/iterator"
	"log"
)

type PostService interface {
	GetAllPosts() []model.Post
	SavePost(post *model.Post) *model.Post
}

type PostServiceImpl struct {
}

func NewPostService() PostService {
	return &PostServiceImpl{}
}

func (s *PostServiceImpl) GetAllPosts() []model.Post {
	log.Print("Getting all posts")

	ctx := context.Background()
	client := db.CreateFirestoreClient(ctx)
	defer client.Close()

	iter := client.Collection("posts").Documents(ctx)
	var posts []model.Post
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}

		var post model.Post
		doc.DataTo(&post)
		posts = append(posts, post)
	}

	return posts
}

func (s *PostServiceImpl) SavePost(post *model.Post) *model.Post {
	log.Printf("Saving post: %s", post)

	ctx := context.Background()
	client := db.CreateFirestoreClient(ctx)
	defer client.Close()

	_, err := client.Collection("posts").Doc(post.Id).Set(ctx, post)
	if err != nil {
		log.Fatalf("Failed adding post: %v", err)
	}

	return post
}
