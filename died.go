package died

import (
	"reflect"
)

type Died struct {
	functions []any
	args      [][]reflect.Type
	returns   [][]reflect.Type

	instances map[reflect.Type]reflect.Value
}

func New() *Died {
	return &Died{}
}
