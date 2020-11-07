package app

import (
	"context"
)

// API abstraction
type API interface {
	Run() error
	Shutdown(context.Context) error
	Close() error
}
