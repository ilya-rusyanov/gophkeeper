package interceptor

import (
	"context"
	"fmt"
	"slices"

	"github.com/golang-jwt/jwt/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/ilya-rusyanov/gophkeeper/internal/server/entity"
	grpcctx "github.com/ilya-rusyanov/gophkeeper/internal/server/grpcserver/context"
)

// Auth returns authentication interceptor
func Auth(
	key string,
	skipMethods ...string,
) grpc.ServerOption {
	interceptor := authInterceptor{
		key:         key,
		skipMethods: skipMethods,
	}

	return grpc.UnaryInterceptor(
		func(
			ctx context.Context,
			req interface{},
			info *grpc.UnaryServerInfo,
			handler grpc.UnaryHandler,
		) (interface{}, error) {
			return interceptor.Intercept(
				ctx,
				req,
				info,
				handler,
			)
		},
	)
}

type authInterceptor struct {
	key         string
	skipMethods []string
}

func (i *authInterceptor) Intercept(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (any, error) {
	loginedCtx := ctx

	if !slices.Contains(i.skipMethods, info.FullMethod) {
		var token string
		if md, ok := metadata.FromIncomingContext(ctx); ok {
			values := md.Get("auth")
			if len(values) > 0 {
				token = values[0]
			}
		}
		if len(token) == 0 {
			return nil, status.Error(codes.Unauthenticated, "missing token")
		}

		login, err := i.extractLogin(token)
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, "invalid token")
		}

		loginedCtx = context.WithValue(ctx, grpcctx.ContextKeyLogin, login)
	}

	return handler(loginedCtx, req)
}

func (i *authInterceptor) extractLogin(token string) (string, error) {
	var (
		claims entity.TokenClaims
		res    string
	)

	_, err := jwt.ParseWithClaims(token, &claims,
		func(t *jwt.Token) (interface{}, error) {
			return []byte(i.key), nil
		})
	if err != nil {
		return "", fmt.Errorf("failed to parse token: %w", err)
	}

	res = claims.Login

	return res, nil
}
