package dictionaries

import (
	"context"

	"github.com/atom-providers/jwt"
	"google.golang.org/grpc/metadata"
)

type Common struct{}

func (*Common) Claim(ctx context.Context, j *jwt.JWT) (*jwt.Claims, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, jwt.TokenInvalid
	}

	tokens := md.Get(jwt.HttpHeader)
	if len(tokens) == 0 {
		return nil, jwt.TokenInvalid
	}

	return j.ParseToken(tokens[0])
}
