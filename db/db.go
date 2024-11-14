package db

import(
	"context"
    "firebase.google.com/go"
    "google.golang.org/api/option"
    "log"
    // "fmt"
    "os"
    "cloud.google.com/go/firestore"
)

var client *firestore.Client

func InitDatabase() error {
	// set-up the configs of the db
	ctx := context.Background()
	opt := option.WithCredentialsFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))

	// connect to the firebase
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return err
	}

	// init firestore client
	client, err = app.Firestore(ctx)
	if err != nil {
		return err
	}

	log.Println("Firestore database initialized successfully")
	return nil
}

func SaveURL(original string, shortened string) error {
	_, err := client.Collection("urls").Doc(shortened).Set(context.Background(), map[string]interface{}{
		"original": original,
	})
	return err
}

func GetURL(shortened string) (string, error) {
	doc, err := client.Collection("urls").Doc(shortened).Get(context.Background())
	if err != nil {
		return "", err
	}

	original, err := doc.DataAt("original")
	if err != nil {
		return "", err
	}
	return original.(string), nil
}