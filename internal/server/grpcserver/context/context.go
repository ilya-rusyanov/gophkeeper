package context

// ContextKey is type for passing data through context
type ContextKey int

const (
	// ContextKeyLogin for passing user logins
	ContextKeyLogin ContextKey = iota
)
