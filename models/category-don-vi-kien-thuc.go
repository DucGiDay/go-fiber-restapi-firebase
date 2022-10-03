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
		data := doc.DataTo(&donViKienThuc)
		log.Println(data, donViKienThuc)
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

	updateData, _ := helper.StructToMapString(donViKienThuc)

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

	_, err := FI.Client.Collection("Category_Don_vi_kien_thuc").Doc(id).Delete(ctx)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	DeleteMoTaChiTietByIdDVKT(id)
	return nil
}

func DeleteDonViKienThucByIdDKT(id string) error {
	var FI config.FirebaseInstance = config.FI
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	iter := FI.Client.Collection("Category_Don_vi_kien_thuc").Where("Id_category_dkt", "==", id).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
		_, err = FI.Client.Collection("Category_Don_vi_kien_thuc").Doc(doc.Ref.ID).Delete(ctx)
		if err != nil {
			log.Fatalln(err)
			return err
		}
		DeleteMoTaChiTietByIdDVKT(doc.Ref.ID)
	}

	return nil
}
