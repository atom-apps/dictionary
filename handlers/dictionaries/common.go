package dictionaries

import (
	"context"

	"github.com/atom-providers/jwt"
	"go-micro.dev/v4/metadata"
)

type Common struct{}

func (*Common) Claim(ctx context.Context, j *jwt.JWT) (*jwt.Claims, error) {
	token, ok := metadata.Get(ctx, jwt.HttpHeader)
	if ok {
		return nil, jwt.TokenInvalid
	}

	return j.ParseToken(token)
}
