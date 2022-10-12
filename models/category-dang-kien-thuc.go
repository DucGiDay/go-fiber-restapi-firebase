package models

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/DucGiDay/go-fiber-restapi-firebase/helper"

	"cloud.google.com/go/firestore"
	"github.com/DucGiDay/go-fiber-restapi-firebase/config"
	"google.golang.org/api/iterator"
)

type DangKienThuc struct {
	Name string `json:"Name"`
	Slug string `json:"Slug"`
}

func List() ([]DangKienThuc, []string, error) {
	var FI config.FirebaseInstance = config.FI
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var dangKienThucs []DangKienThuc
	var IDs []string

	iter := FI.Client.Collection("Category_Dang_Kien_Thuc").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}
		var dangKienThuc DangKienThuc
		data := doc.DataTo(&dangKienThuc)
		log.Println(data, dangKienThuc)

		dangKienThucs = append(dangKienThucs, dangKienThuc)
		IDs = append(IDs, doc.Ref.ID)
	}

	return dangKienThucs, IDs, nil
}

func Read(id string) (DangKienThuc, error, string) {
	var FI config.FirebaseInstance = config.FI
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var dangKienThuc DangKienThuc
	dsnap, err := FI.Client.Collection("Category_Dang_Kien_Thuc").Doc(id).Get(ctx)
	// FI.Client.Close()
	if err != nil {
		log.Fatalln(err)
		return dangKienThuc, err, "" ///đang ko trả về đc lỗi nếu ko tìm thấy dữ liệu
	}
	data := dsnap.DataTo(&dangKienThuc) ///convert from map[string]interface{} to struct type
	data2 := dsnap.Data()
	data2["id"] = dsnap.Ref.ID
	log.Println(data)
	log.Println(data2)
	dangKienThucJson, _ := helper.MapToJson(data2)

	return dangKienThuc, nil, string(dangKienThucJson)
}

func Create(dangKienThuc DangKienThuc) (DangKienThuc, error) {
	var FI config.FirebaseInstance = config.FI
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	iter, temp, err := FI.Client.Collection("Category_Dang_Kien_Thuc").Add(ctx, dangKienThuc)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(iter, temp)

	return dangKienThuc, nil
}

func Update(id string, dangKienThuc DangKienThuc) (DangKienThuc, error) {
	var FI config.FirebaseInstance = config.FI
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	updateData, _ := helper.StructToMapString(dangKienThuc)
	_, err := FI.Client.Collection("Category_Dang_Kien_Thuc").Doc(id).Set(ctx, updateData, firestore.MergeAll)
	if err != nil {
		log.Fatalln(err)
	}
	return dangKienThuc, nil
}

func Delete(id string) error {
	var FI config.FirebaseInstance = config.FI
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := FI.Client.Collection("Category_Dang_Kien_Thuc").Doc(id).Delete(ctx)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	DeleteDonViKienThucByIdDKT(id)
	return nil
}
