package store

import "github.com/ilya-rusyanov/gophkeeper/internal/client/entity"

// AuthStorager is storage of user's own credentials
type AuthStorager interface {
	Load() (entity.MyAuthentication, error)
}
