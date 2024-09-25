package died

import (
	"reflect"

	"github.com/suziziziz/died/internal/inout"
)

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
