package dbc

import (
	"fmt"

	"github.com/tikafog/of/config"

	"github.com/tikafog/of/dbc/bootnode"
	"github.com/tikafog/of/dbc/kernel"
)

type client interface {
	*bootnode.Client | *kernel.Client
}

type OpenFunc[T client] func(name, path string, dbconfig config.Database) (T, error)

type module[T client] struct {
	name  string
	funcs map[string]OpenFunc[T]
}

func (m module[T]) Name() string {
	return m.name
}

func open[T client](name, path string, dbconfig config.Database) (T, error) {
	m := module[T]{
		name: name,
		funcs: map[string]OpenFunc[T]{
			"bootnode": openBootNode[T],
			"kernel":   openKernel[T],
		},
	}
	return m.open(path, dbconfig)
}

func (m module[T]) open(path string, dbconfig config.Database) (T, error) {
	v, exist := m.funcs[m.name]
	if exist {
		return v(m.name, path, dbconfig)
	}
	return nil, fmt.Errorf("module[%s] not found", m.name)
}
