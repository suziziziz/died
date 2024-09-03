package died

import (
	"reflect"
)

type Died interface {
	Inject(function any)
	Invoke()

	// # Require(requirerFn func(dependency Type)) any
	//
	// Require a instantied dependency, invoke all dependencies before call this.
	Require(requirerFn any) any
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
