package firebase

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func GetApp(ctx context.Context) (*firebase.App, error) {
	opt := option.WithCredentialsFile("./pkg/firebase/serviceAccountKey.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("Error initializing firebase app: %v", err)
		return nil, err
	}
	return app, nil
}

func GetClient(ctx *context.Context) (*firestore.Client, error) {
	app, err := GetApp(*ctx)
	if err != nil {
		return nil, err
	}
	client, err := app.Firestore(*ctx)
	if err != nil {
		return nil, err
	}
	return client, err
}
