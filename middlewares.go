package main

import (
	"time"

	"github.com/go-chi/chi/middleware"
)

const DefaultHTTPTimeout = 30 * time.Second

func NewTimeoutMiddleware() Middleware {
	return middleware.Timeout(DefaultHTTPTimeout)
}

func NewHTTPRecoverMiddleware() Middleware {
	return middleware.Recoverer
}

func NewRequestIDMiddleware() Middleware {
	return middleware.RequestID
}
