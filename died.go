package died

import (
	"reflect"
)

type Died interface {
	Inject(function any)
	Invoke()
}

type died struct {
	functions []any
	args      [][]reflect.Type
	returns   [][]reflect.Type

	instances map[reflect.Type]reflect.Value
}

// Instances a new DIED
func New() Died {
	return &died{}
}
