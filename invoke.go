package died

import "reflect"

// Initialize injection of dependencies
func (d *died) Invoke() {
	d.instances = make(map[reflect.Type]reflect.Value)

	for fnIdx, fn := range d.functions {
		args := []reflect.Value{}

		for _, arg := range d.args[fnIdx] {
			args = append(args, d.instances[arg])
		}

		fnInst := reflect.ValueOf(fn).Call(args)

		for retIdx, ret := range d.returns[fnIdx] {
			d.instances[ret] = fnInst[retIdx]
		}
	}
}
