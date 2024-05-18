package interfaces

import (
	"context"
)

type Queue interface {
	Publish(ctx context.Context, name string, b []byte) (err error)
}
