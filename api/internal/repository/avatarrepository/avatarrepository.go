package avatarrepository

import (
	"GoReactProject/data/migrations"
	"GoReactProject/internal/model/avatarmodel"
)

func CreateAvatar(request avatarmodel.CreateAvatarRequest) avatarmodel.CreateUserResponse {
	db := migrations.ConnectPostgresServer()
	var result avatarmodel.CreateUserResponse
	var data migrations.Avatar
	if err := db.Where("userid = ?", request.Userid).First(&data).Error; err != nil {
		data.File = request.Image
		data.Userid = request.Userid

		if err := db.Create(&data).Error; err != nil {
			result.Resulttype = false
			result.Message = "Fail!"
			return result
		}
		result.Resulttype = true
		result.Message = "Success!"
		return result
	} else {
		data.File = request.Image
		if err := db.Save(&data).Error; err != nil {
			result.Resulttype = false
			result.Message = "Fail!"
			return result
		}
		result.Resulttype = true
		result.Message = "Success!"
		return result
	}
}

func GetAvatar(id int) []byte {
	db := migrations.ConnectPostgresServer()
	var data migrations.Avatar
	if err := db.Where("userid = ?", id).First(&data).Error; err != nil {
		return nil
	}
	return data.File
}
