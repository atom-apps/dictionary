package boot

import (
	"strings"

	"github.com/atom-apps/common/consts"
	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/gofiber/fiber/v2"
)

func httpMiddlewareJWT(door *casdoorsdk.Client) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		if strings.HasPrefix(ctx.Path(), "/docs/") {
			return ctx.Next()
		}

		token, ok := ctx.GetReqHeaders()[consts.JwtHttpHeader.String()]
		if !ok {
			return ctx.SendStatus(fiber.StatusUnauthorized)
		}

		if !strings.HasPrefix(token, consts.JwtTokenPrefix.String()) {
			return ctx.SendStatus(fiber.StatusUnauthorized)
		}

		token = strings.TrimSpace(token[len(consts.JwtTokenPrefix.String()):])

		claims, err := door.ParseJwtToken(token)
		if err != nil {
			return ctx.SendStatus(fiber.StatusUnauthorized)
		}
		ctx.Locals(consts.JwtCtx, claims)

		return ctx.Next()
	}
}
