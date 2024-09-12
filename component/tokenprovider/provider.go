package tokenprovider

import (
	"errors"
	"greport/common"
	"time"
)

const (
	DefaultExpiryTime = 10 * 60
	RefreshExpiryTime = 30 * 24 * 60 * 60
	ServiceName       = "tokenprovider"
	RefreshToken      = "refreshtoken"
	MaximumToken      = 4 // maximum number of refresh token
)

type Provider interface {
	Generate(data TokenPayload, expiry int) (*Token, error)
	Validate(token string) (*TokenPayload, error)
	RefreshToken(token string) (*Token, error)
}

type Token struct {
	Token         string    `json:"token"`
	RefreshToken  string    `json:"refresh_token"`
	Created       time.Time `json:"created"`
	Expiry        int       `json:"expiry"`
	RefreshExpiry int       `json:"refresh_expiry"`
}

type TokenPayload struct {
	UserId     int       `json:"user_id"`
	Username   string    `json:"username"`
	Role       string    `json:"role"`
	IssuedAt   time.Time `json:"issued_at"`
	Expiration time.Time `json:"expiration"`
}

type TokenManager struct {
	Id         string `json:"id"`
	DeviceType string `json:"device_type"`
	Token      string `json:"token"`
}

var (
	ErrTokenExpired = common.NewCustomError(
		errors.New("token expired"),
		"token expired",
		"ErrTokenExpired",
	)
	ErrEncodingToken = common.NewCustomError(
		errors.New("error encoding the token"),
		"Error encoding the token",
		"ErrEncodingToken",
	)
	ErrInvalidToken = common.NewCustomError(
		errors.New("invalid Token"),
		"Err Invalid Token",
		"ErrInvalidToken",
	)
	ErrMaximumToken = common.NewCustomError(
		errors.New("maximum number of token"),
		"Err maximum number of Token",
		"ErrMaximumToken",
	)
)
