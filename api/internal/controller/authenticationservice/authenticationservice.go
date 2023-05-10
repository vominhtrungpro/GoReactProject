package authenticationservice

import (
	"GoReactProject/internal/model/authenticationmodel"
	"GoReactProject/internal/model/usermodel"
	"GoReactProject/internal/repository/userrepository"
	"crypto/rand"
	"encoding/hex"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var sampleSecretKey = []byte("SecretYouShouldHide")

func GenerateJWT() string {
	var sampleSecretKey = []byte("SecretYouShouldHide")
	token := jwt.New(jwt.SigningMethodEdDSA)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(10 * time.Minute)
	claims["authorized"] = true
	claims["username"] = "username"
	tokenString, err := token.SignedString(sampleSecretKey)
	if err != nil {
		return ""
	}
	return tokenString
}

func GetUserByUsername(username string) (usermodel.UserDto, error) {
	result, err := userrepository.GetUserByUsername(username)
	if err != nil {
		return result, err
	}
	return result, nil
}

func GetUserByEmail(email string) (usermodel.UserDto, error) {
	var user usermodel.UserDto
	result, err := userrepository.GetUserByEmail(email)
	if err != nil {
		return user, err
	}
	return result.User, nil
}

func CreateAccessToken(user usermodel.UserDto) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.Id,
		"username": user.Username,
		"email":    user.Email,
		"exp":      time.Now().AddDate(0, 0, 7),
	})

	tokenString, err := token.SignedString(sampleSecretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func CreateAccessToken2(user usermodel.UserDto) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.MapClaims{
		"id":    user.Id,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 240).Unix(),
		"role":  "None",
		"name":  user.Username,
	})

	secrettoken := os.Getenv("SECRET_TOKEN")

	tokenString, err := token.SignedString([]byte(secrettoken))

	return tokenString, err
}

func CreateRefreshToken() (authenticationmodel.Refreshtokenresponse, error) {
	var refreshtoken authenticationmodel.Refreshtokenresponse
	b := make([]byte, 64)
	if _, err := rand.Read(b); err != nil {
		return refreshtoken, err
	}
	token := hex.EncodeToString(b)
	refreshtoken.Token = token
	refreshtoken.Tokenexpire = time.Now().AddDate(0, 0, 30)
	return refreshtoken, nil
}
