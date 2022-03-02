package module

import (
	"context"
	"encoding/json"

	"github.com/tikalink/of"
	"github.com/tikalink/of/option"
)

type emptyModule struct {
	name of.Name
}

func (m emptyModule) WithInit(o option.InitializeOption) of.Module {
	return nil
}

func (m emptyModule) WithOption(o option.Option) of.Module {
	return nil
}

func (m emptyModule) RegisterAPI(api of.API) error {
	return nil
}

func (m emptyModule) Init() error {
	return nil
}

func (m emptyModule) Destroy() error {
	return nil
}

func (m emptyModule) PreloadCore(core of.Core) error {
	return nil
}

func (m emptyModule) SetCore(core of.Core) error {
	return nil
}

func (m emptyModule) Valid() bool {
	return false
}

func (m emptyModule) Run(ctx context.Context) error {
	return nil
}

func (m emptyModule) Name() of.Name {
	return m.name
}

func (m emptyModule) HandleData(id string, data json.RawMessage) error {
	return nil
}

func NewEmptyModule(name of.Name) of.Module {
	return newEmptyModule(name)
}

func newEmptyModule(name of.Name) Loader {
	return &emptyModule{
		name: name,
	}
}
