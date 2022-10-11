package models

import (
	"context"
	"fmt"
	"log"
	"time"

	helper "github.com/DucGiDay/go-fiber-restapi-firebase/helper"

	// "cloud.google.com/go/firestore"
	"cloud.google.com/go/firestore"
	config "github.com/DucGiDay/go-fiber-restapi-firebase/config"
	"google.golang.org/api/iterator"
)

type Cauhoi struct {
	Don_Kep          int       `json:"Don_Kep"`
	Date             time.Time `json:"Date"`
	Content_Question string    `json:"Content_Question"`
	Option_ans       []string  `json:"Option_ans"`
	True_ans         string    `json:"True_ans"`
	Id_cate_dkt      string    `json:"Id_cate_dkt"`
	Id_cate_dvkt     string    `json:"Id_cate_dvkt"`
	Id_cate_mtct     string    `json:"Id_cate_mtct"`
	Requirement      string    `json:"Requirement"`
	Slug             string    `json:"Slug"`
	Level            string    `json:"level"`
}

func ListCauHoi() ([]Cauhoi, []string, error) {
	var FI config.FirebaseInstance = config.FI
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var cauHois []Cauhoi
	var Ids []string
	iter := FI.Client.Collection("cau_Hoi").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		var cauHoi Cauhoi
		data := doc.DataTo(&cauHoi)
		log.Println(data, cauHoi)
		cauHois = append(cauHois, cauHoi)
		Ids = append(Ids, doc.Ref.ID)
	}

	return cauHois, Ids, nil
}

func ReadCauHoi(id string) (Cauhoi, error) {
	var FI config.FirebaseInstance = config.FI
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var cauHoi Cauhoi
	dsnap, err := FI.Client.Collection("cau_Hoi").Doc(id).Get(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	data := dsnap.DataTo(&cauHoi)
	log.Println(data)
	return cauHoi, nil

}

func CreateCauHoi(cauHoi Cauhoi) (Cauhoi, error) {
	var FI config.FirebaseInstance = config.FI
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	iter, temp, err := FI.Client.Collection("cau_Hoi").Add(ctx, cauHoi)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(iter, temp)
	return cauHoi, nil
}

func UpdateCauHoi(id string, cauHoi Cauhoi) (Cauhoi, error) {
	var FI config.FirebaseInstance = config.FI
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	updateData, _ := helper.StructToMapString(cauHoi)
	_, err := FI.Client.Collection("cau_Hoi").Doc(id).Set(ctx, updateData, firestore.MergeAll)
	if err != nil {
		log.Fatalln(err)
	}
	return cauHoi, nil
}

func DeleteCauHoi(id string) error {
	var FI config.FirebaseInstance = config.FI
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := FI.Client.Collection("cau_Hoi").Doc(id).Update(ctx, []firestore.Update{
		{
			Path:  "capital",
			Value: firestore.Delete,
		},
	})
	if err != nil {
		log.Println("An error has occurred: %s", err)
	}
	_, err = FI.Client.Collection("cau_Hoi").Doc(id).Delete(ctx)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	return nil
}
