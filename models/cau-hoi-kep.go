package models

import (
	"context"
	"fmt"
	"log"
	"time"

	config "github.com/DucGiDay/go-fiber-restapi-firebase/config"
	"github.com/DucGiDay/go-fiber-restapi-firebase/helper"
	"google.golang.org/api/iterator"
)

type CauHoiKep struct {
	Don_Kep              int                    `json:"Don_Kep"`
	Date                 time.Time              `json:"Date"`
	Question_Main        string                 `json:"Question_Main"`
	Id_cate_dkt          string                 `json:"Id_cate_dkt"`
	Id_cate_dvkt         string                 `json:"Id_cate_dvkt"`
	Id_cate_mtct         string                 `json:"Id_cate_mtct"`
	Question_Requirement string                 `json:"Question_Requirement"`
	Level_Main           string                 `json:"Level_Main"`
	Sub_Question         map[string]interface{} `json:"Sub_Question"`
}

func ListCauHoiKep() ([]string, error) {
	var FI config.FirebaseInstance = config.FI
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var cauHois []string
	var Ids []string
	iter := FI.Client.Collection("cau_Hoi").Where("Don_Kep", "==", 1).Documents(ctx)
	// iter := FI.Client.Collection("cau_Hoi").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		// var cauHoi Cauhoi
		// data := doc.DataTo(&cauHoi)
		data := doc.Data()
		data["id"] = doc.Ref.ID
		cauhoi, _ := helper.MapToJson(data)
		log.Println(data, cauhoi)
		cauHois = append(cauHois, string(cauhoi))
		Ids = append(Ids, doc.Ref.ID)
	}

	return cauHois, nil
}
