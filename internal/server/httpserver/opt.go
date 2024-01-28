package httpserver

import "net/http"

// Opt is a funcopt for server
type Opt func(*http.Server) error
