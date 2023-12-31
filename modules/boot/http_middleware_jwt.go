package boot

import (
	"strings"

	"github.com/atom-providers/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/rogeecn/fen"
)

func httpMiddlewareJWT(j *jwt.JWT) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		if strings.HasPrefix(ctx.Path(), "/docs/") {
			return ctx.Next()
		}
		token := ctx.Request().Header.Peek(jwt.HttpHeader)
		if token == nil {
			return ctx.SendStatus(fiber.StatusUnauthorized)
		}

		claims, err := j.ParseToken(string(token))
		if err != nil {
			return ctx.SendStatus(fiber.StatusUnauthorized)
		}
		ctx.Locals(fen.JwtCtxKey, claims)

		return ctx.Next()
	}
}
