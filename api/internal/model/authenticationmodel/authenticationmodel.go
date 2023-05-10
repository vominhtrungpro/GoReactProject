package authenticationmodel

import "time"

type LoginRequest struct {
	Username string
	Password string
}

type LoginGoogleRequest struct {
	Email string
}

type LoginResponse struct {
	Responsetype bool
	Message      string
	Accesstoken  string
	Refreshtoken string
}

type Refreshtokenresponse struct {
	Token       string
	Tokenexpire time.Time
}
