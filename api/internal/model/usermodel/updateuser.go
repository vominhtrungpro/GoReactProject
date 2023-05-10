package usermodel

type UpdateUserRequest struct {
	Id        int
	Firstname string
	Lastname  string
	Username  string
	Email     string
	Password  string
}

type UpdateUserResponse struct {
	Resulttype bool
	Message    string
}
