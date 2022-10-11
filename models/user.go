package models

import (
	"context"
	"fmt"
	"log"
	"time"

	// "encoding/json"

	helper "github.com/DucGiDay/go-fiber-restapi-firebase/helper"
	"cloud.google.com/go/firestore"
	"github.com/DucGiDay/go-fiber-restapi-firebase/config"
	"google.golang.org/api/iterator"
)

type User struct {
	Username string    `json:"Username"`
	Email    string    `json:"Email"`
	Age      int       `json:"Age"`
	Role	 string	   `ison:"Role"`
}

func GetAllUsers() ([]User, []string, error) {
	var FI config.FirebaseInstance = config.FI
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var users []User
	var UId []string

	iter := FI.Client.Collection("users").Documents(ctx)
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
		log.Println(data)
		if user.Role == "User" {
			users = append(users, user)
			UId = append(UId, doc.Ref.ID)
			log.Println("UId: ",UId)
		}
	}
	return users, UId, nil
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
	data := dsnap.DataTo(&user)
	 ///convert from map[string]interface{} to struct type
	log.Println(data)
	return user, nil
}

func CreateUser(user User) (User, error) {
	var FI config.FirebaseInstance = config.FI
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	iter, temp, err := FI.Client.Collection("users").Add(ctx, user)
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

	updateData, _ := helper.StructToMapString(user)
	_, err := FI.Client.Collection("users").Doc(userId).Set(ctx, updateData, firestore.MergeAll)
	if err != nil {
		log.Fatalln(err)
	}
	return user, nil
}

func DeleteUser(userId string) error {
	var FI config.FirebaseInstance = config.FI
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := FI.Client.Collection("users").Doc(userId).Delete(ctx)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	return nil
}
