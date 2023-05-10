package userapi

import (
	"GoReactProject/internal/controller/userservice"
	"GoReactProject/internal/model/usermodel"
	"encoding/json"
	"net/http"
	"strconv"

	googleAuthIDTokenVerifier "github.com/futurenda/google-auth-id-token-verifier"
	"github.com/gin-gonic/gin"
)

func GetUsers(context *gin.Context) {
	result := userservice.GetUsers()
	context.IndentedJSON(http.StatusOK, result)
}

func GetUserById(context *gin.Context) {
	userid := context.Param("id")

	id, err := strconv.Atoi(userid)
	if err != nil {
		http.Error(context.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	result := userservice.GetUserById(id)
	context.IndentedJSON(http.StatusOK, result)
}

func GetUserByEmail(context *gin.Context) {
	useremail := context.Param("email")
	claimSet, err := googleAuthIDTokenVerifier.Decode(useremail)
	if err != nil {
		http.Error(context.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	email := claimSet.Email
	result := userservice.GetUserByEmail(email)
	context.IndentedJSON(http.StatusOK, result)
}

func CreateUser(context *gin.Context) {
	var request usermodel.CreateUserRequest
	err := json.NewDecoder(context.Request.Body).Decode(&request)
	if err != nil {
		http.Error(context.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	result := userservice.CreateUser(request)
	context.IndentedJSON(http.StatusOK, result)
}

func UpdateUser(context *gin.Context) {
	var request usermodel.UpdateUserRequest
	err := json.NewDecoder(context.Request.Body).Decode(&request)
	if err != nil {
		http.Error(context.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	result := userservice.UpdateUser(request)
	context.IndentedJSON(http.StatusOK, result)
}

func DeleteUser(context *gin.Context) {
	var request usermodel.DeleteUserRequest
	err := json.NewDecoder(context.Request.Body).Decode(&request)
	if err != nil {
		http.Error(context.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	result := userservice.DeleteUsers(request)
	context.IndentedJSON(http.StatusOK, result)
}
