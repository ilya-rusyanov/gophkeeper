package usercred

import (
	"fmt"
	"os"
)

type Logger interface {
	Debugf(string, ...any)
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
func (c *UserCred) Store(login, password string) error {
	err := os.WriteFile(c.userNameFileName, []byte(login), 0666)
	if err != nil {
		return fmt.Errorf("failed to write login file: %w", err)
	}
	c.log.Debugf("updated user login file %q with %q", c.userNameFileName, login)

	return nil
}
