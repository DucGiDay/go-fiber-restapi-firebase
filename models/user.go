package models

import (
	"context"
	"log"
	"time"
	"fmt"
	// "encoding/json"

	"github.com/DucGiDay/go-fiber-restapi-firebase/config"
	"google.golang.org/api/iterator"
	"github.com/google/uuid"
	"cloud.google.com/go/firestore"
)

type User struct {
	// ID       primitive.ObjectID `json:"id"`
	ID			 uuid.UUID					`json:"id"`
	UserID   string							`json:"userId"`
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
		var user User
		data := doc.DataTo(&user) //convert thành struct và lưu vào user
		log.Println(data, user)

		// Phần convert này tạm thời ko dùng đến. Đã convert ở trên
		// //convert map[string]interface{} to json string
		// jsonStrData, err := json.Marshal(data)
    // if err != nil {
		// 	fmt.Println(err)
    // }

		// // Convert json string to struct
		// var user User
    // if err := json.Unmarshal(jsonStrData, &user); err != nil {
		// 	fmt.Println(err)
    // }
		users = append(users, user)
	}

	return users, nil
}

func GetUser(userId string) (User, error) {
	var FI config.FirebaseInstance = config.FI
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var user User
	dsnap, err := FI.Client.Collection("users").Doc(userId).Get(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	data := dsnap.DataTo(&user) ///convert from map[string]interface{} to struct type
	log.Println(data)

	return user, nil
}

func CreateUser(user User) (User, error) {
	var FI config.FirebaseInstance = config.FI
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	iter, temp, err:= FI.Client.Collection("users").Add(ctx, user)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(iter, temp)

	return user, nil
}

func UpdateUser(userId string, user User) (User, error) {
	var FI config.FirebaseInstance = config.FI
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := FI.Client.Collection("users").Doc(userId).Set(ctx, user, firestore.MergeAll)
	if err != nil {
		log.Fatalln(err)
	}
	return user, nil
}

func DeleteUser(userId string) (error) {
	var FI config.FirebaseInstance = config.FI
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := FI.Client.Collection("users").Doc(userId).Delete(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	return nil
}
