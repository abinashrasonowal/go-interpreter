package safety

import (
	"context"
	"time"
)

const DefaultTimeout = 30 * time.Second

func NewTimeoutContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), DefaultTimeout)
}
