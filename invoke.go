package died

import (
	"fmt"
	"reflect"
)

func (d *died) Invoke() {
	d.instances = make(map[reflect.Type]reflect.Value)

	for fnIdx, fn := range d.functions {
		args := []reflect.Value{}

		for _, arg := range d.args[fnIdx] {
			if _, ok := d.instances[arg]; !ok {
				panic(fmt.Sprint("instance ", arg, " required by ", reflect.TypeOf(fn), " is missing"))
			}

			args = append(args, d.instances[arg])
		}

		fnInst := reflect.ValueOf(fn).Call(args)

		for retIdx, ret := range d.returns[fnIdx] {
			d.instances[ret] = fnInst[retIdx]
		}
	}
}
