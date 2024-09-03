package died

import "reflect"

func (d *died) Require(requirerFn any) any {
	requirerType := reflect.TypeOf(requirerFn)

	if requirerType.Kind() != reflect.Func {
		panic("argument 'requirerFn' must be of type func(dep Type)")
	}

	depType := requirerType.In(0)

	return d.instances[depType].Interface()
}
