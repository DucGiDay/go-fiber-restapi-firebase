package config

import (
	"context"
	// "fmt"
	"log"
	// "os"
	"time"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"cloud.google.com/go/firestore"
)

//Cái này là dùng với mongodb
// type MongoInstance struct {
// 	Client *mongo.Client
// 	DB     *mongo.Database
// }
// var MI MongoInstance

type FirebaseInstance struct {
	Client *firestore.Client
	Err     error
}
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

	defer cancel()
}
