package db

import (
	"cloud.google.com/go/firestore"
	"context"
	"log"

	firebase "firebase.google.com/go"
)

func CreateFirestoreClient(ctx context.Context) *firestore.Client {
	// Use a service account
	//sa := option.WithCredentialsFile("serviceAccount.json")
	conf := &firebase.Config{ProjectID: "go-api-test-a7aad"}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	return client
}
