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

// Instances a new DIED
func New() *Died {
	return &Died{}
}
