package paseto

import (
	"2ndbrand-api/component/myredis"
	"2ndbrand-api/component/tokenprovider"
	"context"
	"errors"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
)

type pasetoProvider struct {
	paseto *paseto.V2
	secret string
}

func NewPasetoProvider(secret string) *pasetoProvider {
	if len(secret) < chacha20poly1305.KeySize {
		panic(errors.New("invalid key size"))
	}
	return &pasetoProvider{
		secret: secret,
		paseto: paseto.NewV2(),
	}
}

// GenerateToken
// - data: tokenprovider.TokenPayload
// - expiry: seconds
func (p *pasetoProvider) Generate(data tokenprovider.TokenPayload, expiry int) (*tokenprovider.Token, error) {
	if expiry < 1 {
		expiry = tokenprovider.DefaultExpiryTime
	}
	// generate access token
	if data.Expiration.IsZero() {
		data.Expiration = time.Now().Add(time.Duration(expiry) * time.Second)
	}
	if data.IssuedAt.IsZero() {
		data.IssuedAt = time.Now()
	}
	accessToken, err := p.paseto.Encrypt([]byte(p.secret), data, nil)
	if err != nil {
		return nil, err
	}
	// Generate refresh token
	// refreshToken := uuid.New().String()
	data.Expiration = time.Now().Add(time.Duration(tokenprovider.RefreshExpiryTime) * time.Second)
	refreshToken, err := p.paseto.Encrypt([]byte(p.secret), data, nil)
	if err != nil {
		return nil, err
	}
	key, _ := myredis.GenKey(tokenprovider.ServiceName, tokenprovider.RefreshToken, data.Username)
	mrd := myredis.NewClient(myredis.DB_USER)
	defer mrd.Close()
	// check number of refresh token
	if err := p.checkMaximumToken(key); err != nil {
		return nil, err
	}
	// add refresh token
	if err := p.addRefreshToken(key, refreshToken); err != nil {
		return nil, err
	}

	return &tokenprovider.Token{
		Token:         accessToken,
		Expiry:        expiry,
		Created:       time.Now(),
		RefreshToken:  refreshToken,
		RefreshExpiry: tokenprovider.RefreshExpiryTime,
	}, nil
}

func (p *pasetoProvider) Validate(token string) (*tokenprovider.TokenPayload, error) {
	var payload tokenprovider.TokenPayload
	err := p.paseto.Decrypt(token, []byte(p.secret), &payload, nil)
	if err != nil {
		return nil, tokenprovider.ErrInvalidToken
	}
	if payload.Expiration.Before(time.Now()) {
		return nil, tokenprovider.ErrTokenExpired
	}
	return &payload, nil
}

func (p *pasetoProvider) RefreshToken(token string) (*tokenprovider.Token, error) {
	// validate token
	payload, err := p.Validate(token)
	if err != nil {
		return nil, err
	}
	payload.Expiration = time.Now().Add(time.Duration(tokenprovider.RefreshExpiryTime) * time.Second)
	payload.IssuedAt = time.Now()
	//generate new token
	newToken, err := p.Generate(*payload, tokenprovider.RefreshExpiryTime)
	if err != nil {
		return nil, err
	}
	key, _ := myredis.GenKey(tokenprovider.ServiceName, tokenprovider.RefreshToken, payload.Username)
	// remove old key
	if err := p.removeRefreshToken(key, token); err != nil {
		return nil, err
	}
	// add new key
	if err := p.addRefreshToken(key, newToken.RefreshToken); err != nil {
		return nil, err
	}
	return newToken, nil
}

func (p *pasetoProvider) removeRefreshToken(key, token string) error {
	mrd := myredis.NewClient(myredis.DB_USER)
	defer mrd.Close()
	// remove old key
	if err := mrd.SRem(context.Background(), key, token).Err(); err != nil {
		return err
	}
	return nil
}

// add new token
func (p *pasetoProvider) addRefreshToken(key, token string) error {
	mrd := myredis.NewClient(myredis.DB_USER)
	defer mrd.Close()
	if err := mrd.SAdd(context.Background(), key, token).Err(); err != nil {
		return err
	}
	return nil
}

// check maximum number of token
func (p *pasetoProvider) checkMaximumToken(key string) error {
	mrd := myredis.NewClient(myredis.DB_USER)
	defer mrd.Close()
	n, err := mrd.SCard(context.Background(), key).Result()
	if err != nil {
		return err
	}
	// check number of refresh token
	if n >= tokenprovider.MaximumToken {
		return tokenprovider.ErrMaximumToken
	}
	return nil
}
