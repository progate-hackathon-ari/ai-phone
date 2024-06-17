package oauth

import (
	"context"

	"golang.org/x/oauth2"
)

type OAuth2 interface {
	FetchToken(ctx context.Context, authorizationCode string) (*oauth2.Token, error)
	GetUserInfo(ctx context.Context, token *oauth2.Token) (*UserInfo, error)
}

type UserInfo struct {
	UserID string
	Email  string
	Icon   string
}
