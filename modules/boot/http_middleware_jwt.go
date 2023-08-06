package boot

import (
	"strings"
	"time"

	"github.com/atom-providers/jwt"
	"github.com/gofiber/fiber/v2"
)

func httpMiddlewareJWT(j *jwt.JWT) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		token, ok := ctx.GetReqHeaders()[jwt.HttpHeader]
		if !ok {
			return ctx.SendStatus(fiber.StatusUnauthorized)
		}

		if !strings.HasPrefix(token, jwt.TokenPrefix) {
			return ctx.SendStatus(fiber.StatusUnauthorized)
		}

		token = token[len(jwt.TokenPrefix):]

		claims, err := j.ParseToken(token)
		if err != nil {
			return ctx.SendStatus(fiber.StatusUnauthorized)
		}
		ctx.Locals(jwt.CtxKey, claims)

		// 10 分钟过期前需要交换新token回前端
		if time.Now().Add(time.Minute * 10).After(claims.ExpiresAt.Time) {
			newToken, err := j.CreateTokenByOldToken(token, claims)
			if err == nil {
				ctx.Response().Header.Set("Authorization", newToken)
			}
		}
		return ctx.Next()
	}
}
