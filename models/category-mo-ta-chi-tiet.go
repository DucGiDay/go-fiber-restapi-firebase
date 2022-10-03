package models

import (
	"context"
	"fmt"
	"log"
	"time"

	// "encoding/json"

	"cloud.google.com/go/firestore"
	"github.com/DucGiDay/go-fiber-restapi-firebase/config"
	"github.com/DucGiDay/go-fiber-restapi-firebase/helper"
	"google.golang.org/api/iterator"
)

type MoTaChiTiet struct {
	Name             string `json:"Name"`
	Slug             string `json:"Slug"`
	Id_category_dvkt string `json:"Id_category_dvkt"`
	IsCheck          bool   `json:"IsCheck"`
}

func ListMoTaChiTiets() ([]MoTaChiTiet, []string, error) {
	var FI config.FirebaseInstance = config.FI
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var moTaChiTiets []MoTaChiTiet
	var IDs []string

	iter := FI.Client.Collection("Category_mo_ta_chi_tiet").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}
		var moTaChiTiet MoTaChiTiet
		data := doc.DataTo(&moTaChiTiet) //convert thành struct và lưu vào user
		log.Println(data, moTaChiTiet)

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
		moTaChiTiets = append(moTaChiTiets, moTaChiTiet)
		IDs = append(IDs, doc.Ref.ID)
	}

	return moTaChiTiets, IDs, nil
}

func ReadMoTaChiTiet(id string) (MoTaChiTiet, error) {
	var FI config.FirebaseInstance = config.FI
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var moTaChiTiet MoTaChiTiet
	dsnap, err := FI.Client.Collection("Category_mo_ta_chi_tiet").Doc(id).Get(ctx)
	if err != nil {
		log.Fatalln(err)
		return moTaChiTiet, err ///đang ko trả về đc lỗi nếu ko tìm thấy dữ liệu
	}
	data := dsnap.DataTo(&moTaChiTiet) ///convert from map[string]interface{} to struct type
	log.Println(data)

	return moTaChiTiet, nil
}

func CreateMoTaChiTiet(moTaChiTiet MoTaChiTiet) (MoTaChiTiet, error) {
	var FI config.FirebaseInstance = config.FI
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	data, _ := helper.StructToMapString(moTaChiTiet)
	iter, temp, err := FI.Client.Collection("Category_mo_ta_chi_tiet").Add(ctx, data)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(iter, temp)

	return moTaChiTiet, nil
}

func UpdateMoTaChiTiet(id string, moTaChiTiet MoTaChiTiet) (MoTaChiTiet, error) {
	var FI config.FirebaseInstance = config.FI
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Tạm thời update từng field vì chưa tìm ra nguyên nhân các key bị viết hoa lung tung
	updateData, _ := helper.StructToMapString(moTaChiTiet)
	// temp, _ := helper.StructToMapString(donViKienThuc)
	// updateData := map[string]interface{}{
	// 	"Name":            temp["name"],
	// 	"Id_category_dkt": temp["id_category_dkt"],
	// 	"Slug":            temp["slug"],
	// }
	_, err := FI.Client.Collection("Category_mo_ta_chi_tiet").Doc(id).Set(ctx, updateData, firestore.MergeAll)
	if err != nil {
		log.Fatalln(err)
	}

	return moTaChiTiet, nil
}

func DeleteMoTaChiTiet(id string) error {
	var FI config.FirebaseInstance = config.FI
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// _, err := FI.Client.Collection("Category_mo_ta_chi_tiet").Doc(id).Update(ctx, []firestore.Update{
	// 	{
	// 		Path:  "capital",
	// 		Value: firestore.Delete,
	// 	},
	// })
	// if err != nil {
	// 	log.Printf("An error has occurred: %s", err)
	// }
	_, err := FI.Client.Collection("Category_mo_ta_chi_tiet").Doc(id).Delete(ctx)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	return nil
}

func DeleteMoTaChiTietByIdDVKT(id string) error {
	var FI config.FirebaseInstance = config.FI
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	iter := FI.Client.Collection("Category_mo_ta_chi_tiet").Where("Id_category_dvkt", "==", id).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
		_, err = FI.Client.Collection("Category_mo_ta_chi_tiet").Doc(doc.Ref.ID).Delete(ctx)
		if err != nil {
			log.Fatalln(err)
			return err
		}
	}

	return nil
}
