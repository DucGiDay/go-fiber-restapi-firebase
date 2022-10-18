package controllers

import(
	"time"
	"fmt"
	models "github.com/DucGiDay/go-fiber-restapi-firebase/models"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)
const jwtSecret = "asecret"

type request struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(ctx *fiber.Ctx) error {
	var body request
	err := ctx.BodyParser(&body)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse json",
		})
		
	}
	users, errs := models.GetAllAuths()
	if errs != nil {
		return errs
	}
	responseDatas := []string{}
	fmt.Println("auths",users)
	for i, user := range users{
		fmt.Println(user.Email)
		if body.Email != user.Email || body.Password != user.Password {
			i = i +1
		}else{
			UserByte, _ := json.Marshal(user)
			UserString := string(UserByte)
			responseDatas = append(responseDatas, UserString)
		}
	}
	fmt.Println("responseDatas",responseDatas)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = "1"
	claims["exp"] = time.Now().Add(time.Hour * 24 * 7) // a week

	s, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
		
	}

	fmt.Println("token", s)
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": s,
		"user": responseDatas,
	})
}
