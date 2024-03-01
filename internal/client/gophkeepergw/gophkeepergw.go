package gophkeepergw

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"

	"github.com/ilya-rusyanov/gophkeeper/internal/client/entity"
	"github.com/ilya-rusyanov/gophkeeper/proto"
)

type Logger interface {
	Debugf(string, ...any)
}

// GophKeeperGW is a gateway to the actual service
type GophKeeperGW struct {
	serverAddr string
	log        Logger
}

// New creates an instance of the gateway
func New(serverAddr string, log Logger) *GophKeeperGW {
	return &GophKeeperGW{
		serverAddr: serverAddr,
		log:        log,
	}
}

// Register registers new user
func (gk *GophKeeperGW) Register(
	ctx context.Context, cred entity.MyCredentials,
) (entity.MyAuthentication, error) {
	var res entity.MyAuthentication

	err := gk.withConn(func(conn *grpc.ClientConn) error {
		var header metadata.MD

		c := proto.NewGophkeeperClient(conn)

		arg := proto.RegisterRequest{
			Credentials: toProtoCredentials(cred),
		}

		_, err := c.Register(ctx, &arg, grpc.Header(&header))
		if err != nil {
			return fmt.Errorf("server error: %w", err)
		}
		gk.log.Debugf("got server header: %q", header)

		tokens := header.Get("token")
		if len(tokens) != 1 {
			return fmt.Errorf("received %d tokens from server", len(tokens))
		}

		res = entity.NewMyAuthentication(tokens[0])

		return nil
	})
	if err != nil {
		return res, fmt.Errorf("connection failed: %w", err)
	}

	return res, nil
}

// LogIn logs user in
func (gk *GophKeeperGW) LogIn(
	ctx context.Context, cred entity.MyCredentials,
) (entity.MyAuthentication, error) {
	var res entity.MyAuthentication

	err := gk.withConn(func(conn *grpc.ClientConn) error {
		var header metadata.MD

		c := proto.NewGophkeeperClient(conn)

		arg := proto.LogInRequest{
			Credentials: &proto.UserCredentials{
				Login:    cred.Login,
				Password: cred.Password,
			},
		}

		_, err := c.LogIn(ctx, &arg, grpc.Header(&header))
		if err != nil {
			return fmt.Errorf("server error: %w", err)
		}
		gk.log.Debugf("got server header: %q", header)

		tokens := header.Get("token")
		if len(tokens) != 1 {
			return fmt.Errorf("received %d tokens from server", len(tokens))
		}

		res = entity.NewMyAuthentication(tokens[0])

		return nil
	})
	if err != nil {
		return res, fmt.Errorf("connection failed: %w", err)
	}

	return res, nil
}

// Store stores given data
func (gk *GophKeeperGW) Store(
	ctx context.Context, in entity.ServiceStoreRequest,
) error {
	return gk.withConn(func(conn *grpc.ClientConn) error {
		c := proto.NewGophkeeperClient(conn)

		md := metadata.Pairs("auth", string(in.AuthData))
		authCtx := metadata.NewOutgoingContext(ctx, md)

		protoMeta, err := toProtoMeta(in.Record.Meta)
		if err != nil {
			return fmt.Errorf("failed to convert meta info to proto: %w", err)
		}

		protoPayload, err := toProtoPayload(in.Record.Payload)
		if err != nil {
			return fmt.Errorf("failed to convert payload to proto: %w", err)
		}

		arg := proto.StoreRequest{
			Type: string(in.Record.Type),
			Name:  in.Record.Name,
			Meta:protoMeta,
			Payload:protoPayload,
		}

		_, err = c.Store(authCtx, &arg)
		if err != nil {
			return fmt.Errorf("server failed to store data: %w", err)
		}

		return nil
	})
}

func (gk *GophKeeperGW) withConn(f func(conn *grpc.ClientConn) error) error {
	conn, err := grpc.Dial(
		gk.serverAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return fmt.Errorf("failed to dial server: %w", err)
	}
	defer conn.Close()

	err = f(conn)
	if err != nil {
		return fmt.Errorf("failed to execute closure: %w", err)
	}

	return nil
}

func toProtoCredentials(in entity.MyCredentials) *proto.UserCredentials {
	return &proto.UserCredentials{
		Login:    in.Login,
		Password: in.Password,
	}
}

func toProtoMeta(in entity.Meta) (string, error) {
	sb := strings.Builder{}

	err := json.NewEncoder(&sb).Encode(&in)
	if err != nil {
		return "", fmt.Errorf("json encoder failed to encode meta: %w", err)
	}

	return sb.String(), nil
}

func toProtoPayload(in any) ([]byte, error) {
	res, err := json.Marshal(in)
	if err != nil {
		return res, fmt.Errorf("failed to encode payload to json: %w", err)
	}

	return res, nil
}
