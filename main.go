package library

import (
	"crypto/ed25519"
	library "git.ailur.dev/ailur/fg-library/v3"
	"github.com/google/uuid"
	"io"
	"time"
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
	Name   string           `validate:"required"`
	User   uuid.UUID        `validate:"required"`
	Reader io.LimitedReader // Only used for write operations
}

type Quota struct {
	User  uuid.UUID `validate:"required"`
	Bytes int64     `validate:"required"`
}

var (
	StorageService = uuid.MustParse("00000000-0000-0000-0000-000000000003")
	OAuthService   = uuid.MustParse("00000000-0000-0000-0000-000000000004")
)

func OAuthSignup(oauth OAuthInformation, information *library.ServiceInitializationInformation) (OAuthResponse, error) {
	message, err := information.SendAndAwaitISMessage(OAuthService, 1, oauth, time.Second*3)
	if err != nil {
		return OAuthResponse{}, err
	}

	return message.Message.(OAuthResponse), nil

}

func GetPublicKey(information *library.ServiceInitializationInformation) (ed25519.PublicKey, error) {
	message, err := information.SendAndAwaitISMessage(OAuthService, 2, nil, time.Second*3)
	if err != nil {
		return nil, err
	}

	return message.Message.(ed25519.PublicKey), nil
}

func GetOAuthHostname(information *library.ServiceInitializationInformation) (string, error) {
	message, err := information.SendAndAwaitISMessage(OAuthService, 0, nil, time.Second*3)
	if err != nil {
		return "", err
	}

	return message.Message.(string), nil
}

func InitializeOAuth(oauth OAuthInformation, information *library.ServiceInitializationInformation) (oauthResponse OAuthResponse, pk ed25519.PublicKey, hostname string, err error) {
	pk, err = GetPublicKey(information)
	if err != nil {
		return
	}

	hostname, err = GetOAuthHostname(information)
	if err != nil {
		return
	}

	oauthResponse, err = OAuthSignup(oauth, information)
	if err != nil {
		return
	}

	return
}
