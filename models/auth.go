package models

import (
	"context"
	"fmt"
	"log"
	"time"
	"github.com/DucGiDay/go-fiber-restapi-firebase/config"
	"google.golang.org/api/iterator"
)

type Auth struct {
	Password string    `json:"Password"`
	Email    string    `json:"Email"`
	Role	 string	   `ison:"Role"`
}

func GetAllAuths() ([]Auth, error) {
	var FI config.FirebaseInstance = config.FI
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var auths []Auth

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
		var auth Auth
		data := doc.DataTo(&auth) //convert thành struct và lưu vào user
		log.Println(data)
		if auth.Role == "Admin" || auth.Role == "Editer" {
			auths = append(auths, auth)
		}
	}
	return auths, nil
}