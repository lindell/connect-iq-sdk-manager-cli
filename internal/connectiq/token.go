package connectiq

import (
	"context"
	"time"
)

// Token contains information about the token used for requests to the Garmin API
type Token struct {
	AccessToken string    `json:"accessToken"`
	ExpiresAt   time.Time `json:"expiresAt"`
}

type tokenCtxKey struct{}

func SetContextToken(ctx context.Context, token Token) context.Context {
	return context.WithValue(ctx, tokenCtxKey{}, token)
}

func GetContextToken(ctx context.Context) (Token, bool) {
	val := ctx.Value(tokenCtxKey{})
	if val == nil {
		return Token{}, false
	}
	return val.(Token), true
}
