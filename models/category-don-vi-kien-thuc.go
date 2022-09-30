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

type DonViKienThuc struct {
	Name            string `json:"Name"`
	Slug            string `json:"Slug"`
	Id_category_dkt string `json:"Id_category_dkt"`
}

func ListDonViKienThucs() ([]DonViKienThuc, []string, error) {
	var FI config.FirebaseInstance = config.FI
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var donViKienThucs []DonViKienThuc
	var IDs []string

	iter := FI.Client.Collection("Category_Don_vi_kien_thuc").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}
		var donViKienThuc DonViKienThuc
		data := doc.DataTo(&donViKienThuc) //convert thành struct và lưu vào user
		log.Println(data, donViKienThuc)

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
		donViKienThucs = append(donViKienThucs, donViKienThuc)
		IDs = append(IDs, doc.Ref.ID)
	}

	return donViKienThucs, IDs, nil
}

func ReadDonViKienThuc(id string) (DonViKienThuc, error) {
	var FI config.FirebaseInstance = config.FI
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var donViKienThuc DonViKienThuc
	dsnap, err := FI.Client.Collection("Category_Don_vi_kien_thuc").Doc(id).Get(ctx)
	if err != nil {
		log.Fatalln(err)
		return donViKienThuc, err ///đang ko trả về đc lỗi nếu ko tìm thấy dữ liệu
	}
	data := dsnap.DataTo(&donViKienThuc) ///convert from map[string]interface{} to struct type
	log.Println(data)

	return donViKienThuc, nil
}

func CreateDonViKienThuc(donViKienThuc DonViKienThuc) (DonViKienThuc, error) {
	var FI config.FirebaseInstance = config.FI
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	data, _ := helper.StructToMapString(donViKienThuc)
	iter, temp, err := FI.Client.Collection("Category_Don_vi_kien_thuc").Add(ctx, data)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(iter, temp)

	return donViKienThuc, nil
}

func UpdateDonViKienThuc(id string, donViKienThuc DonViKienThuc) (DonViKienThuc, error) {
	var FI config.FirebaseInstance = config.FI
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Tạm thời update từng field vì chưa tìm ra nguyên nhân các key bị viết hoa lung tung
	updateData, _ := helper.StructToMapString(donViKienThuc)
	// temp, _ := helper.StructToMapString(donViKienThuc)
	// updateData := map[string]interface{}{
	// 	"Name":            temp["name"],
	// 	"Id_category_dkt": temp["id_category_dkt"],
	// 	"Slug":            temp["slug"],
	// }
	_, err := FI.Client.Collection("Category_Don_vi_kien_thuc").Doc(id).Set(ctx, updateData, firestore.MergeAll)
	if err != nil {
		log.Fatalln(err)
	}

	return donViKienThuc, nil
}

func DeleteDonViKienThuc(id string) error {
	var FI config.FirebaseInstance = config.FI
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := FI.Client.Collection("Category_Don_vi_kien_thuc").Doc(id).Update(ctx, []firestore.Update{
		{
			Path:  "capital",
			Value: firestore.Delete,
		},
	})
	if err != nil {
		log.Printf("An error has occurred: %s", err)
	}
	_, err = FI.Client.Collection("Category_Don_vi_kien_thuc").Doc(id).Delete(ctx)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	return nil
}
