package model

import (
	"context"
	"log"
	"time"
	"fmt"
	"encoding/json"

	"github.com/DucGiDay/go-fiber-restapi-firebase/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/api/iterator"
)

type User struct {
	// ID       primitive.ObjectID `json:"id" bson:"_id"`
	ID			 int								`json:"id"`
	Username string             `json:"username"`
	Email    string             `json:"email"`
	Age      int                `json:"age"`
}

func GetAllUsers() ([]User, error) {
	var FI config.FirebaseInstance = config.FI
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var users []User
	
	iter:= FI.Client.Collection("users").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}

		data := doc.Data()
		log.Println(data)

		//convert map[string]interface{} to json string
		jsonStrData, err := json.Marshal(data)
    if err != nil {
			fmt.Println(err)
    }

		// Convert json string to struct
		var user User
    if err := json.Unmarshal(jsonStrData, &user); err != nil {
			fmt.Println(err)
    }
		users = append(users, user)
	}
	fmt.Println("users: ", users)

	return users, nil
}

func GetUser(userId primitive.ObjectID) (User, error) {
	var MI config.MongoInstance = config.MI

	var user User
	collection := MI.DB.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	findResult := collection.FindOne(ctx, bson.M{"_id": userId})
	if err := findResult.Err(); err != nil {
		return user, err
	}

	err := findResult.Decode(&user)

	if err != nil {
		return user, err
	}
	return user, nil
}

func CreateUser(user User) (User, error) {
	var MI config.MongoInstance = config.MI
	collection := MI.DB.Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
	}

	return user, nil
}

func UpdateUser(userId primitive.ObjectID, user User) (User, error) {
	var MI config.MongoInstance = config.MI

	userMap := map[string]interface{}{
		"username": user.Username,
		"email":    user.Email,
		"age":      user.Age,
	}

	collection := MI.DB.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.UpdateOne(ctx, bson.M{"_id": userId}, bson.M{"$set": userMap})
	if err != nil {
		return user, err
	}
	return user, nil
}
