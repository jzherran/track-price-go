package main

import (
	"time"

	"github.com/go-chi/chi/middleware"
)

// DefaultHTTPTimeout ...
const DefaultHTTPTimeout = 30 * time.Second

// NewTimeoutMiddleware ...
func NewTimeoutMiddleware() Middleware {
	return middleware.Timeout(DefaultHTTPTimeout)
}

// NewHTTPRecovererMiddleware ...
func NewHTTPRecovererMiddleware() Middleware {
	return middleware.Recoverer
}

// NewRequestIDMiddleware ...
func NewRequestIDMiddleware() Middleware {
	return middleware.RequestID
}
