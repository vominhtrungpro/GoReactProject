package userservice

import (
	"GoReactProject/internal/model/authenticationmodel"
	"GoReactProject/internal/model/usermodel"
	"GoReactProject/internal/repository/userrepository"
)

func GetUsers() usermodel.GetUserResponse {
	result := userrepository.GetUsers()
	return result
}

func GetUserById(id int) usermodel.GetOnlyUserResponse {
	result := userrepository.GetUserById(id)
	return result
}

func GetUserByEmail(email string) usermodel.GetOnlyUserResponse {
	result, err := userrepository.GetUserByEmail(email)
	if err != nil {
		return result
	}
	return result
}

func CreateUser(request usermodel.CreateUserRequest) usermodel.CreateUserResponse {
	result := userrepository.CreateUser(request)
	return result
}

func UpdateUser(request usermodel.UpdateUserRequest) usermodel.UpdateUserResponse {
	result := userrepository.UpdateUser(request)
	return result
}

func DeleteUsers(request usermodel.DeleteUserRequest) usermodel.DeleteUserResponse {
	result := userrepository.DeleteUsers(request)
	return result
}

func UpdateRefreshTokenUser(id int, refreshtoken authenticationmodel.Refreshtokenresponse) (string, error) {
	_, err := userrepository.UpdateRefreshTokenUser(id, refreshtoken.Token, refreshtoken.Tokenexpire)
	if err != nil {
		return "", err
	}
	return refreshtoken.Token, nil
}
