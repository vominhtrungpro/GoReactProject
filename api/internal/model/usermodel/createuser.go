package usermodel

type CreateUserRequest struct {
	Firstname string
	Lastname  string
	Username  string
	Email     string
	Password  string
}

type CreateUserResponse struct {
	Resulttype bool
	Message    string
}
