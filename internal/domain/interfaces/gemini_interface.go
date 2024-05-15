package interfaces

import (
	"context"
)

type Gemini interface {
	Ask(ctx context.Context, input string) (result *string, err error)
}
