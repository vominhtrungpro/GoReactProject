package usermodel

import "time"

type UserDto struct {
	Id               int
	Firstname        string
	Lastname         string
	Username         string
	Email            string
	Password         string
	Refreshtokens    *string
	RefreshTokensexp *time.Time
}

type GetUserResponse struct {
	User       []UserDto
	Resulttype bool
	Message    string
}

type GetOnlyUserResponse struct {
	User       UserDto
	Resulttype bool
	Message    string
}
