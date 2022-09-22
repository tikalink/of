package merr

import (
	"errors"
	"fmt"
)

type ModuleError interface {
	Name() string
	Index() uint32
	Errors() int
	NewIndex(str string) Index
	IndexString(index Index) string
	IndexError(index Index) error

	New(str string) error
	Errorf(format string, args ...interface{}) error
	Is(err, target error) bool
	Unwrap(err error) error
	As(err error, target any) bool
	Wrap(err error, s string) error
}

type moduleError struct {
	idx    uint32
	name   string
	count  uint32
	errors []error
}

func (m *moduleError) IndexString(index Index) string {
	//fmt.Sprintf("Module[%v]:%v", index.Name(), err.Error())
	return m.IndexError(index).Error()
}

func (m *moduleError) Name() string {
	return m.name
}

func (m *moduleError) Index() uint32 {
	return m.idx
}

func (m *moduleError) Errors() int {
	return len(m.errors)
}

func (m *moduleError) New(str string) error {
	_, err := m.NewIndexError(str)
	return err
}

func (m *moduleError) Wrap(err error, s string) error {
	_, err = m.WrapIndexError(err, s)
	return err
}

func (m *moduleError) WrapIndex(err error, s string) Index {
	idx, _ := m.WrapIndexError(err, s)
	return idx
}

func (m *moduleError) WrapIndexError(err error, s string) (Index, error) {
	m.count += 1
	idx := makeErrIndex(m.Index(), m.count)
	werr := WrapString(err, s)
	m.errors = append(m.errors, werr)
	return idx, IndexError(idx)
}

func (m *moduleError) NewIndex(str string) Index {
	idx, _ := m.NewIndexError(str)
	return idx
}

func (m *moduleError) NewIndexError(str string) (Index, error) {
	m.count += 1
	idx := makeErrIndex(m.Index(), m.count)
	err := New(str)
	m.errors = append(m.errors, err)
	return idx, IndexError(idx)
}

func (m *moduleError) Errorf(format string, args ...interface{}) error {
	_, err := m.NewIndexErrorf(format, args...)
	return err
}

func (m *moduleError) NewIndexErrorf(format string, args ...interface{}) (Index, error) {
	m.count += 1
	idx := makeErrIndex(m.Index(), m.count)
	err := Errorf(format, args...)
	m.errors = append(m.errors, err)
	return idx, IndexError(idx)
}

func (m *moduleError) getErrorIndex(index Index) uint32 {
	idx := (uint32(index) - 1) ^ IndexModule(m.Index())
	fmt.Println("errindex: ", idx)
	if idx >= m.count {
		return 0
	}
	return idx
}

func (m *moduleError) IndexError(index Index) error {
	return m.errors[m.getErrorIndex(index)]
}

func (m *moduleError) Is(err, target error) bool {
	return errors.Is(err, target)
}

func (m *moduleError) Unwrap(err error) error {
	return errors.Unwrap(err)
}

func (m *moduleError) As(err error, target any) bool {
	return errors.As(err, target)
}

func newModuleWithIndex(str string, idx uint32) ModuleError {
	return &moduleError{
		idx:    idx,
		name:   str,
		count:  0,
		errors: nil,
	}
}

func NewModule(str string) ModuleError {
	return newModuleWithIndex(str, getModuleIndex())
}
