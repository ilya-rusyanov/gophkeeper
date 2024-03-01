package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/ilya-rusyanov/gophkeeper/internal/server/entity"
)

// Builder builds authentication tokens
type Builder struct {
	key string
}

// NewBuilder constructs authentication token builder
func NewBuilder(signingKey string) *Builder {
	return &Builder{
		key: signingKey,
	}
}

// Build builds new authentication token
func (b *Builder) Build(
	lifespan time.Duration, login string,
) (entity.AuthToken, error) {
	var result entity.AuthToken

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, entity.TokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(
				time.Now().Add(lifespan)),
		},
		Login: login,
	})

	tokenString, err := token.SignedString([]byte(b.key))
	if err != nil {
		return result, fmt.Errorf("failed to sign token: %w", err)
	}

	result = entity.AuthToken(tokenString)

	return result, nil
}
