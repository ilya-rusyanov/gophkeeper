package shutdown

import "context"

type shutdowner interface {
	Stop(context.Context) error
}
