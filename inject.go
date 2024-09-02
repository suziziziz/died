package died

import (
	"reflect"

	"github.com/suziziziz/died/internal/inout"
)

// Provide dependencies with the passed builder.
//
// You can provide more than one dependency per builder if necessary, but
// dependencies must have a unique return type.
//
// Arguments must be the same as the return types of the provided dependencies.
//
// No need to sort, DIED will do it for you.
func (d *died) Inject(function any) {
	fn := reflect.TypeOf(function)

	if fn.Kind() != reflect.Func {
		panic("argument named 'function' must be of type 'func'")
	}

	args := inout.ExtractIn(fn)
	returns := inout.ExtractOut(fn)

	d.functions = append(d.functions, function)
	d.args = append(d.args, args)
	d.returns = append(d.returns, returns)

	d.sortDependencies()
}
