package services

import (
	"GOTEST/cmd/models"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateJWT(user models.LoginedUser) string {
	claims := jwt.MapClaims{
		"usercd": user.Usercd,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	fmt.Print(token)

	return "TEST"
}

func AnalysisToken(string) {

}
