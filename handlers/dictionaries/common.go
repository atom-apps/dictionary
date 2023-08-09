package dictionaries

import (
	"context"

	"github.com/atom-apps/common/consts"
	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Common struct{}

func (*Common) Claim(ctx context.Context, door *casdoorsdk.Client) (*casdoorsdk.Claims, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "metadata is not provided")
	}

	tokens := md.Get(consts.JwtHttpHeader.String())
	if len(tokens) == 0 {
		return nil, status.Error(codes.Unauthenticated, "token is empty")
	}

	return door.ParseJwtToken(tokens[0])
}
