package usermodel

type DeleteUserRequest struct {
	Id []int
}

type DeleteUserResponse struct {
	Resulttype bool
	Message    string
}
