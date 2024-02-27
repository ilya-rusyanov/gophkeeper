package usercred

import (
	"fmt"
	"os"

	keyring "github.com/zalando/go-keyring"

	"github.com/ilya-rusyanov/gophkeeper/internal/client/entity"
)

type Logger interface {
	Debugf(string, ...any)
	Debug(...any)
}

// UserCred is storage for app user's credentials
type UserCred struct {
	appName          string
	userNameFileName string
	log              Logger
}

// New creates storage
func New(log Logger, usernameFilename, appName string) *UserCred {
	return &UserCred{
		appName:          appName,
		userNameFileName: usernameFilename,
		log:              log,
	}
}

// Store saves login and password
func (c *UserCred) Store(cred entity.MyCredentials) error {
	if err := os.WriteFile(c.userNameFileName, []byte(cred.Login), 0o600); err != nil {
		return fmt.Errorf("failed to write login file: %w", err)
	}
	c.log.Debugf("updated user login file %q with %q", c.userNameFileName, cred.Login)

	if err := keyring.Set(c.appName, cred.Login, cred.Password); err != nil {
		return fmt.Errorf("keyring failed to store the password: %w", err)
	}
	c.log.Debug("successfully stored user credentials in keyring")

	return nil
}
