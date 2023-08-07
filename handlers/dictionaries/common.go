package dictionaries

import (
	"context"

	"github.com/atom-providers/jwt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Common struct{}

func (*Common) Claim(ctx context.Context, j *jwt.JWT) (*jwt.Claims, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "metadata is not provided")
	}

	tokens := md.Get(jwt.HttpHeader)
	if len(tokens) == 0 {
		return nil, status.Error(codes.Unauthenticated, "token is empty")
	}

	return j.ParseToken(tokens[0])
}
