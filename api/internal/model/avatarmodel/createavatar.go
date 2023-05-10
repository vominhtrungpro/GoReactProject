package avatarmodel

type CreateAvatarRequest struct {
	Image  []byte
	Userid int
}

type CreateAvatarAPIRequest struct {
	Userid int
}

type CreateUserResponse struct {
	Resulttype bool
	Message    string
}
