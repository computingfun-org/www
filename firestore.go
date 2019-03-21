package main

import (
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"gitlab.com/zacc/autocertcache"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
)

// NewFirestoreClient ...
func NewFirestoreClient(ctx context.Context, credentials string) (*firestore.Client, error) {
	options := option.WithCredentialsFile(credentials)
	app, err := firebase.NewApp(ctx, nil, options)
	if err != nil {
		return nil, err
	}
	return app.Firestore(ctx)
}

// NewFirestoreClientFatal ...
func NewFirestoreClientFatal(ctx context.Context, credentials string) *firestore.Client {
	fsc, err := NewFirestoreClient(ctx, credentials)
	if err != nil {
		log.Fatalln(err)
	}
	return fsc
}

// NewFirestoreCache ...
func NewFirestoreCache(client *firestore.Client, collection string) *autocertcache.Firestore {
	return autocertcache.NewFirestoreFromClient(client, collection)
}

// NewFirestoreCacheFatal ...
func NewFirestoreCacheFatal(client *firestore.Client, collection string) *autocertcache.Firestore {
	fsCache := NewFirestoreCache(client, collection)
	if fsCache == nil {
		log.Fatalln("NewFirestoreCache returned nil in NewFirestoreCacheFatal.")
	}
	return fsCache
}
