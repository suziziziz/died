package inout

import "reflect"

func ExtractOut(fn reflect.Type) []reflect.Type {
	returns := []reflect.Type{}

	for rIdx := range fn.NumOut() {
		returns = append(returns, fn.Out(rIdx))
	}

	return returns
}
