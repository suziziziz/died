package died

import (
	"reflect"
	"slices"
)

func (d *Died) sortDependencies() {
	for idx, rets := range d.returns {
		apply := false

		for _, ret := range rets {
			argI := slices.IndexFunc(d.args, func(a []reflect.Type) bool {
				return slices.Index(a, ret) != -1
			})

			if argI != -1 {
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
