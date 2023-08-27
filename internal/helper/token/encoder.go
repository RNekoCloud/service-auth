package token

import (
	"time"

	"github.com/o1egl/paseto"
)

func EncodeToken(id, email string) (string, error) {
	now := time.Now()
	exp := now.Add(48 * time.Hour)

	payload := paseto.JSONToken{
		Audience:   "User",
		Issuer:     "CVZN Developer",
		Jti:        id,
		Subject:    email,
		IssuedAt:   now,
		Expiration: exp,
		NotBefore:  exp,
	}

	payload.Set("data", id)
	footer := "Copyright 2023 CV Zaman Now"

	token, err := paseto.NewV2().Sign(Priv, payload, footer)

	if err != nil {
		return "", err
	}

	return token, nil
}
