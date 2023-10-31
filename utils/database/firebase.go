package database

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"firebase.google.com/go/messaging"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/option"
)

func InitFirebaseApp(sdk string, projectID string, url string) (*firebase.App, *db.Client, *messaging.Client) {
	conf := &firebase.Config{
		DatabaseURL: url,
	}

	file, err := os.Open("credentials.json")
	if err != nil {
		log.Fatalln("Error reading credentials file:", err)
	}

	opt := option.WithCredentialsFile(file.Name())

	app, err := firebase.NewApp(context.Background(), conf, opt)
	if err != nil {
		log.Fatalln("Error initializing app:", err)
	}

	fcm, err := app.Messaging(context.Background())
	if err != nil {
		log.Fatalln(err)
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

	return app, client, fcm
}
