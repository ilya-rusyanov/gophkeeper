package gophkeepergw

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/ilya-rusyanov/gophkeeper/internal/client/entity"
	"github.com/ilya-rusyanov/gophkeeper/proto"
)

// GophKeeperGW is a gateway to the actual service
type GophKeeperGW struct {
	serverAddr string
}

// New creates an instance of the gateway
func New(serverAddr string) *GophKeeperGW {
	return &GophKeeperGW{
		serverAddr: serverAddr,
	}
}

// Register registers new user
func (gk *GophKeeperGW) Register(ctx context.Context, cred entity.MyCredentials) error {
	return gk.withConn(func(conn *grpc.ClientConn) error {
		c := proto.NewGophkeeperClient(conn)

		arg := proto.RegisterRequest{
			Credentials: &proto.UserCredentials{
				Login:    cred.Login,
				Password: cred.Password,
			},
		}

		_, err := c.Register(ctx, &arg)
		if err != nil {
			return fmt.Errorf("server error: %w", err)
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
		return fmt.Errorf("failed to dial to server: %w", err)
	}
	defer conn.Close()

	err = f(conn)
	if err != nil {
		return fmt.Errorf("failed to execute closure: %w", err)
	}

	return nil
}
