package inout

import "reflect"

func ExtractIn(fn reflect.Type) []reflect.Type {
	args := []reflect.Type{}

	for aIdx := range fn.NumIn() {
		args = append(args, fn.In(aIdx))
	}

	return args
}
