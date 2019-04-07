package firestoredb

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var (
	// AutoCertCache ...
	AutoCertCache *firestore.CollectionRef

	// User ...
	User *firestore.CollectionRef

	// Admin ...
	Admin *firestore.CollectionRef

	// Author ...
	Author *firestore.CollectionRef

	// Article ...
	Article *firestore.CollectionRef
)

// InitFatal ...
func InitFatal(ctx context.Context, credentialsFile string) {
	if err := Init(ctx, credentialsFile); err != nil {
		log.Fatalln("🔑  " + err.Error())
	}
}

// Init ...
func Init(ctx context.Context, credentialsFile string) error {
	options := option.WithCredentialsFile(credentialsFile)
	app, err := firebase.NewApp(ctx, nil, options)
	if err != nil {
		return err
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		return err
	}

	AutoCertCache = client.Collection("certs")
	User = client.Collection("users")
	Admin = client.Collection("admins")
	Author = client.Collection("authors")
	Article = client.Collection("articles")

	return nil
}
