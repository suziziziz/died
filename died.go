package died

import (
	"reflect"
)

type Died interface {
	// Provide dependencies with the passed builder.
	//
	// You can provide more than one dependency per builder if necessary, but
	// dependencies must have a unique return type.
	//
	// Arguments must be the same as the return types of the provided dependencies.
	//
	// No need to sort, DIED will do it for you.
	Inject(function any)

	// Initialize injection of dependencies
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
