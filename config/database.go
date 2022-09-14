package config

import (
	"context"
	// "fmt"
	"log"
	// "os"
	"time"

	// "github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"cloud.google.com/go/firestore"
)

type MongoInstance struct {
	Client *mongo.Client
	DB     *mongo.Database
}
type FirebaseInstance struct {
	Client *firestore.Client
	Err     error
}

var MI MongoInstance
var FI FirebaseInstance


func ConnectDatabase() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	sa := option.WithCredentialsFile("./webgolang.json")
	db, err := firebase.NewApp(ctx, nil, sa)
	FI.Client, FI.Err = db.Firestore(ctx)
	if FI.Err != nil {
		log.Fatalln(FI.Err)
	}
	if err != nil {
		log.Fatalln(err)
	}

	// iter:= FI.Client.Collection("posts").Documents(ctx)
	// for {
	// 	doc, err := iter.Next()
	// 	if err == iterator.Done {
	// 		break
	// 	}
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		break
	// 	}
	// 	fmt.Println(doc.Data())
	// }

	defer cancel()
}
