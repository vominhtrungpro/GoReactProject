package migrations

import (
	"time"
)

type Users struct {
	Id               int
	Firstname        string
	Lastname         string
	Username         string
	Email            string
	Password         string
	Refreshtokens    *string
	RefreshTokensexp *time.Time
}

type Avatar struct {
	Id     int
	File   []byte
	Userid int
}
