package died

import (
	"reflect"
	"slices"
)

func (d *died) sortDependencies() {
	for idx, rets := range d.returns {
		apply := false

		for _, ret := range rets {
			hasArg := slices.ContainsFunc(d.args, func(a []reflect.Type) bool {
				return slices.Contains(a, ret)
			})

			if hasArg {
				apply = true
			}
		}

		if apply {
			saveFunction := d.functions[idx]
			saveArgs := d.args[idx]
			saveReturns := d.returns[idx]

			d.functions = slices.Delete(d.functions, idx, idx+1)
			d.args = slices.Delete(d.args, idx, idx+1)
			d.returns = slices.Delete(d.returns, idx, idx+1)

			d.functions = slices.Insert(d.functions, 0, saveFunction)
			d.args = slices.Insert(d.args, 0, saveArgs)
			d.returns = slices.Insert(d.returns, 0, saveReturns)
		}
	}
}
