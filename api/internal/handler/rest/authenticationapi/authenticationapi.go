package authenticationapi

import (
	"GoReactProject/internal/controller/authenticationservice"
	"GoReactProject/internal/controller/userservice"
	"GoReactProject/internal/model/authenticationmodel"
	"GoReactProject/internal/model/usermodel"
	"encoding/json"
	"net/http"

	googleAuthIDTokenVerifier "github.com/futurenda/google-auth-id-token-verifier"
	"github.com/gin-gonic/gin"
)

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

func Login(context *gin.Context) {
	var request authenticationmodel.LoginRequest
	var response authenticationmodel.LoginResponse
	err := json.NewDecoder(context.Request.Body).Decode(&request)
	if err != nil {
		http.Error(context.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	result, err := authenticationservice.GetUserByUsername(request.Username)
	if err != nil {
		response.Responsetype = false
		response.Message = "User not found!"
		response.Accesstoken = ""
		response.Refreshtoken = ""
		context.IndentedJSON(http.StatusOK, response)
		return
	}
	if request.Password != result.Password {
		response.Responsetype = false
		response.Message = "Wrong password!"
		response.Accesstoken = ""
		response.Refreshtoken = ""
		context.IndentedJSON(http.StatusOK, response)
		return
	}

	accesstoken, err := authenticationservice.CreateAccessToken2(result)
	if err != nil {
		response.Responsetype = false
		response.Message = "Failed to create access token!"
		response.Accesstoken = ""
		response.Refreshtoken = ""
		context.IndentedJSON(http.StatusOK, response)
		return
	}

	refreshtoken, err := authenticationservice.CreateRefreshToken()
	if err != nil {
		response.Responsetype = false
		response.Message = "Failed to create access token!"
		response.Accesstoken = ""
		response.Refreshtoken = ""
		context.IndentedJSON(http.StatusOK, response)
		return
	}
	newrefreshtoken, err := userservice.UpdateRefreshTokenUser(result.Id, refreshtoken)
	if err != nil {
		response.Responsetype = false
		response.Message = "Failed to update refresh token!"
		response.Accesstoken = ""
		response.Refreshtoken = ""
		context.IndentedJSON(http.StatusOK, response)
		return
	}

	response.Responsetype = true
	response.Message = "Success!"
	response.Accesstoken = accesstoken
	response.Refreshtoken = newrefreshtoken
	context.IndentedJSON(http.StatusOK, response)
}

func LoginGoogle(context *gin.Context) {
	useremail := context.Param("token")
	var response authenticationmodel.LoginResponse
	claimSet, err := googleAuthIDTokenVerifier.Decode(useremail)
	if err != nil {
		http.Error(context.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	email := claimSet.Email
	result, err := authenticationservice.GetUserByEmail(email)
	if err != nil {
		response.Responsetype = false
		response.Message = "User not found!"
		response.Accesstoken = ""
		response.Refreshtoken = ""
		context.IndentedJSON(http.StatusOK, response)
		return
	}

	accesstoken, err := authenticationservice.CreateAccessToken2(result)
	if err != nil {
		response.Responsetype = false
		response.Message = "Failed to create access token!"
		response.Accesstoken = ""
		response.Refreshtoken = ""
		context.IndentedJSON(http.StatusOK, response)
		return
	}

	refreshtoken, err := authenticationservice.CreateRefreshToken()
	if err != nil {
		response.Responsetype = false
		response.Message = "Failed to create access token!"
		response.Accesstoken = ""
		response.Refreshtoken = ""
		context.IndentedJSON(http.StatusOK, response)
		return
	}
	newrefreshtoken, err := userservice.UpdateRefreshTokenUser(result.Id, refreshtoken)
	if err != nil {
		response.Responsetype = false
		response.Message = "Failed to update refresh token!"
		response.Accesstoken = ""
		response.Refreshtoken = ""
		context.IndentedJSON(http.StatusOK, response)
		return
	}

	response.Responsetype = true
	response.Message = "Success!"
	response.Accesstoken = accesstoken
	response.Refreshtoken = newrefreshtoken
	context.IndentedJSON(http.StatusOK, response)
}
