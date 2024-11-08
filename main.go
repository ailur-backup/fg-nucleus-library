package library

import (
	"github.com/google/uuid"
)

type OAuthInformation struct {
	Token       string   `json:"token"`
	Name        string   `json:"name"`
	RedirectUri string   `json:"redirectUri"`
	KeyShareUri string   `json:"keyShareUri"`
	Scopes      []string `json:"scopes"`
}

type OAuthResponse struct {
	AppID     string `json:"appId"`
	SecretKey string `json:"secretKey"`
}

type File struct {
	Name  string    `validate:"required"`
	User  uuid.UUID `validate:"required"`
	Bytes []byte    // Only used in write operations
}

type Quota struct {
	User  uuid.UUID `validate:"required"`
	Bytes int64     `validate:"required"`
}
