package avatarservice

import (
	"GoReactProject/internal/model/avatarmodel"
	"GoReactProject/internal/repository/avatarrepository"
)

func CreateAvatar(request avatarmodel.CreateAvatarRequest) avatarmodel.CreateUserResponse {
	result := avatarrepository.CreateAvatar(request)
	return result
}

func GetAvatar(id int) []byte {
	result := avatarrepository.GetAvatar(id)
	return result
}
