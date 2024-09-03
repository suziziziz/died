package died

import (
	"reflect"
	"slices"
)

func (d *died) sortDependencies() {
	for idx, rets := range d.returns {
		idxArg := -1

		for _, ret := range rets {
			hasArg := slices.IndexFunc(d.args, func(a []reflect.Type) bool {
				return slices.Contains(a, ret)
			})

			if hasArg > idxArg {
				idxArg = hasArg
			}
		}

		if idxArg != -1 {
			saveFunction := d.functions[idx]
			saveArgs := d.args[idx]
			saveReturns := d.returns[idx]

			d.functions = slices.Delete(d.functions, idx, idx+1)
			d.args = slices.Delete(d.args, idx, idx+1)
			d.returns = slices.Delete(d.returns, idx, idx+1)

			d.functions = slices.Insert(d.functions, idxArg, saveFunction)
			d.args = slices.Insert(d.args, idxArg, saveArgs)
			d.returns = slices.Insert(d.returns, idxArg, saveReturns)
		}
	}
}
