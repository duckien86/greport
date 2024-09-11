package jwt

import (
	"2ndbrand-api/component/tokenprovider"
	"time"

	"github.com/golang-jwt/jwt"
)

type jwtProvider struct {
	secret string
}

func NewTokenJwtProvider(secret string) *jwtProvider {
	return &jwtProvider{secret: secret}
}

type userClaims struct {
	Payload tokenprovider.TokenPayload `json:"payload"`
	jwt.StandardClaims
}

func (provider *jwtProvider) Generate(payload tokenprovider.TokenPayload, expiry int) (*tokenprovider.Token, error) {
	claims := userClaims{
		payload,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Second * time.Duration(expiry)).Unix(),
			IssuedAt:  time.Now().Local().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := token.SignedString([]byte(provider.secret))
	if err != nil {
		return nil, err
	}
	return &tokenprovider.Token{
		Token:   signedString,
		Expiry:  expiry,
		Created: time.Now(),
	}, nil
}

func (j *jwtProvider) Validate(token string) (*tokenprovider.TokenPayload, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &userClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.secret), nil
	})
	if err != nil {
		return nil, tokenprovider.ErrInvalidToken
	}

	claims, ok := jwtToken.Claims.(*userClaims)
	if !ok || !jwtToken.Valid {
		return nil, tokenprovider.ErrInvalidToken
	}

	return &claims.Payload, nil
}
