package database

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/option"
)

func InitFirebaseApp(sdk string, projectID string, url string) (*firebase.App, *db.Client) {
	// opt := option.WithCredentialsFile(sdk)
	// app, err := firebase.NewApp(context.Background(), nil, opt)
	// if err != nil {
	// 	log.Fatalf("Error initializing Firebase App: %v", err)
	// 	return nil
	// }

	// return app
	////////////////////////////////////////////////////////////////////////
	conf := &firebase.Config{
		DatabaseURL: url,
	}
	opt := option.WithCredentialsFile(sdk)

	app, err := firebase.NewApp(context.Background(), conf, opt)
	if err != nil {
		log.Fatalln("Error initializing app:", err)
	}

	client, err := app.DatabaseWithURL(context.Background(), url)
	if err != nil {
		logrus.Info("url :", url)
		log.Fatalln("Error initializing database client:", err)
	}

	ref := client.NewRef("restricted_access/secret_document")
	var data map[string]interface{}
	if err := ref.Get(context.Background(), &data); err != nil {
		log.Fatalln("Error reading from database:", err)
	}

	return app, client
}
