package userrepository

import (
	"GoReactProject/data/migrations"
	"GoReactProject/internal/model/usermodel"
	"time"
)

func GetUsers() usermodel.GetUserResponse {
	db := migrations.ConnectPostgresServer()
	var result usermodel.GetUserResponse
	var usersdb []migrations.Users
	var users []usermodel.UserDto
	if err := db.Order("id").Find(&usersdb).Error; err != nil {
		result.Resulttype = false
		result.Message = "Fail"
		result.User = users
		return result
	}
	for _, element := range usersdb {
		user := usermodel.UserDto(element)
		users = append(users, user)
	}
	result.Resulttype = true
	result.Message = "Success"
	result.User = users
	return result
}

func GetUserById(id int) usermodel.GetOnlyUserResponse {
	db := migrations.ConnectPostgresServer()
	var result usermodel.GetOnlyUserResponse
	var userdb migrations.Users
	var users usermodel.UserDto
	if err := db.Where("id = ?", id).First(&userdb).Error; err != nil {
		result.Resulttype = false
		result.Message = "Fail"
		result.User = users
		return result
	}
	user := usermodel.UserDto(userdb)
	result.Resulttype = true
	result.Message = "Success"
	result.User = user
	return result
}

func GetUserByEmail(email string) (usermodel.GetOnlyUserResponse, error) {
	db := migrations.ConnectPostgresServer()
	var result usermodel.GetOnlyUserResponse
	var userdb migrations.Users
	var users usermodel.UserDto
	if err := db.Where("email = ?", email).First(&userdb).Error; err != nil {
		result.Resulttype = false
		result.Message = "Fail"
		result.User = users
		return result, err
	}
	user := usermodel.UserDto(userdb)
	result.Resulttype = true
	result.Message = "Success"
	result.User = user
	return result, nil
}

func GetUserByUsername(username string) (usermodel.UserDto, error) {
	db := migrations.ConnectPostgresServer()
	var result usermodel.UserDto
	var userdb migrations.Users
	if err := db.Where("username = ?", username).First(&userdb).Error; err != nil {
		return result, err
	}
	result1 := usermodel.UserDto(userdb)
	return result1, nil
}

func CreateUser(request usermodel.CreateUserRequest) usermodel.CreateUserResponse {
	db := migrations.ConnectPostgresServer()
	var result usermodel.CreateUserResponse
	hero := migrations.Users{
		Firstname: request.Firstname,
		Lastname:  request.Lastname,
		Username:  request.Username,
		Email:     request.Email,
		Password:  request.Password,
	}
	if err := db.Create(&hero).Error; err != nil {
		result.Resulttype = false
		result.Message = "Fail!"
		return result
	}
	result.Resulttype = true
	result.Message = "Success!"
	return result
}

func UpdateUser(request usermodel.UpdateUserRequest) usermodel.UpdateUserResponse {
	db := migrations.ConnectPostgresServer()
	var result usermodel.UpdateUserResponse
	var user migrations.Users
	if err := db.Where("id = ?", request.Id).First(&user).Error; err != nil {
		result.Resulttype = false
		result.Message = "Fail"
		return result
	}
	user.Firstname = request.Firstname
	user.Lastname = request.Lastname
	user.Email = request.Email
	if err := db.Save(&user).Error; err != nil {
		result.Resulttype = false
		result.Message = "Fail!"
		return result
	}
	result.Resulttype = true
	result.Message = "Success!"
	return result
}

func UpdateRefreshTokenUser(id int, token string, expiredate time.Time) (usermodel.UpdateUserResponse, error) {
	db := migrations.ConnectPostgresServer()
	var result usermodel.UpdateUserResponse
	var user migrations.Users
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		result.Resulttype = false
		result.Message = "Fail"
		return result, err
	}
	user.Refreshtokens = &token
	user.RefreshTokensexp = &expiredate
	if err := db.Save(&user).Error; err != nil {
		result.Resulttype = false
		result.Message = "Fail!"
		return result, err
	}
	result.Resulttype = true
	result.Message = "Success!"
	return result, nil
}

func DeleteUsers(request usermodel.DeleteUserRequest) usermodel.DeleteUserResponse {
	db := migrations.ConnectPostgresServer()
	var result usermodel.DeleteUserResponse
	var users []migrations.Users
	if err := db.Delete(&users, request.Id).Error; err != nil {
		result.Resulttype = false
		result.Message = "Fail!"
		return result
	}
	result.Resulttype = true
	result.Message = "Success!"
	return result
}
