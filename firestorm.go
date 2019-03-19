package main

import (
	"golang.org/x/crypto/acme/autocert"
	"golang.org/x/net/context"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

// NewDataStore ...
func NewDataStore(ctx context.Context, credentials string) (*firestore.Client, error) {
	co := option.WithCredentialsFile(credentials)
	app, err := firebase.NewApp(ctx, nil, co)
	if err != nil {
		return nil, err
	}
	return app.Firestore(ctx)
}

// AutoCertFireStorm ...
type AutoCertFireStorm struct {
	Client     *firestore.Client
	Collection string
}

var _ autocert.Cache = AutoCertFireStorm{}

type fireStormData struct {
	Data []byte
}

// Get ...
func (acfs AutoCertFireStorm) Get(ctx context.Context, key string) ([]byte, error) {
	dsnap, err := acfs.Client.Collection(acfs.Collection).Doc(key).Get(ctx)
	if err != nil {
		return nil, autocert.ErrCacheMiss
	}
	var bytes []byte
	if dsnap.DataTo(&bytes) != nil {
		return nil, autocert.ErrCacheMiss
	}
	return bytes, nil
}

// Put ...
func (acfs AutoCertFireStorm) Put(ctx context.Context, key string, data []byte) error {
	_, err := acfs.Client.Collection(acfs.Collection).Doc(key).Set(ctx, fireStormData{Data: data})
	return err
}

// Delete ...
func (acfs AutoCertFireStorm) Delete(ctx context.Context, key string) error {
	_, err := acfs.Client.Collection(acfs.Collection).Doc(key).Update(ctx, []firestore.Update{
		{Path: key, Value: firestore.Delete},
	})
	return err
}
