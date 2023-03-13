package firebase

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go/v4"
)

func InitAppDefault(ctx context.Context) *firebase.App {
	id, found := os.LookupEnv("GOOGLE_PROJECT_ID")
	if !found {
		log.Fatalln("unable to find GOOGLE_PROJECT_ID environment variable")
	}

	config := firebase.Config{
		ProjectID:     id,
		StorageBucket: id + ".appspot.com",
	}
	app, err := firebase.NewApp(ctx, &config)
	if err != nil {
		log.Fatalf("error initializing firebase FirebaseApp: %v\n", err)
	}

	log.Printf("using Firebase default app id = %s\n", id)

	return app
}
