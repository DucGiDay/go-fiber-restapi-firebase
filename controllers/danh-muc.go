package controllers

import (
	"encoding/json"

	"github.com/DucGiDay/go-fiber-restapi-firebase/models"
	"github.com/gofiber/fiber/v2"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

func ListDanhMuc(c *fiber.Ctx) error {
	dangKienThucs, idDangKienThucs, err := models.List()
	donViKienThucs, idDonViKienThuc, err := models.ListDonViKienThucs()
	moTaChiTiets, idMoTaChiTiet, err := models.ListMoTaChiTiets()
	if err != nil {
		return err
	}
	responseDangKienThucs := []string{}
	for i, dangKienThuc := range dangKienThucs {
		dangKienThucByte, _ := json.Marshal(dangKienThuc)
		dangKienThucString := string(dangKienThucByte)
		responseDangKienThucString := dangKienThucString + " - id: " + idDangKienThucs[i]
		responseDangKienThucs = append(responseDangKienThucs, responseDangKienThucString)
	}
	responseDonViKienThucs := []string{}
	for i, donViKienThuc := range donViKienThucs {
		donViKienThucByte, _ := json.Marshal(donViKienThuc)
		donViKienThucString := string(donViKienThucByte)
		responseDonViKienThucString := donViKienThucString + " - id: " + idDonViKienThuc[i]
		responseDonViKienThucs = append(responseDonViKienThucs, responseDonViKienThucString)
	}
	responseMoTaChiTiets := []string{}
	for i, moTaChiTiet := range moTaChiTiets {
		moTaChiTietByte, _ := json.Marshal(moTaChiTiet)
		moTaChiTietString := string(moTaChiTietByte)
		responseMoTaChiTietString := moTaChiTietString + " - id: " + idMoTaChiTiet[i]
		responseMoTaChiTiets = append(responseMoTaChiTiets, responseMoTaChiTietString)
	}
	responseDatas := map[string]interface{}{
		"dangKienThuc":  "responseDangKienThucs",
		"donViKienThuc": "responseDonViKienThucs",
		"moTaChiTiet":   "responseMoTaChiTiets",
	}

	// fmt.Println("responseDatas", responseDatas)
	return c.JSON(responseDatas)
	// return c.JSON(&fiber.Map{
	// 	"dangKienThuc":  responseDangKienThucs,
	// 	"donViKienThuc": responseDonViKienThucs,
	// 	"moTaChiTiet":   responseMoTaChiTiets,
	// })
}
