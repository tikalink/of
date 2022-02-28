package of

import (
	"context"
	"encoding/json"

	"github.com/tikalink/of/content"
)

type TypeHandleFunc = func(id string, data json.RawMessage) error

type Core interface {
	Connection
	Context() context.Context
	Ask(ctx context.Context, data []byte, last int64) error
	RegisterDataHandler(ct content.Type, fn TypeHandleFunc) error
}
