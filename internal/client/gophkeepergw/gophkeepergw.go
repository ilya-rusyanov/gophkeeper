package gophkeepergw

import (
	"context"
	"fmt"

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
func (gk *GophKeeperGW) Register(ctx context.Context, cred entity.MyCredentials) error {
	return gk.withConn(func(conn *grpc.ClientConn) error {
		var header metadata.MD

		c := proto.NewGophkeeperClient(conn)

		arg := proto.RegisterRequest{
			Credentials: &proto.UserCredentials{
				Login:    cred.Login,
				Password: cred.Password,
			},
		}

		_, err := c.Register(ctx, &arg, grpc.Header(&header))
		if err != nil {
			return fmt.Errorf("server error: %w", err)
		}
		gk.log.Debugf("got server header: %q", header)

		return nil
	})
}

func (gk *GophKeeperGW) withConn(f func(conn *grpc.ClientConn) error) error {
	conn, err := grpc.Dial(
		gk.serverAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return fmt.Errorf("failed to dial to server: %w", err)
	}
	defer conn.Close()

	err = f(conn)
	if err != nil {
		return fmt.Errorf("failed to execute closure: %w", err)
	}

	return nil
}
