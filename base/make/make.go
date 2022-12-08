package make

import "reflect"

var (
	Int    = reflect.TypeOf(0)
	String = reflect.TypeOf("")
)

func IMake(T reflect.Type, fptr any) {
	swap := func(in []reflect.Value) []reflect.Value {
		return []reflect.Value{
			reflect.MakeSlice(reflect.SliceOf(T),
				int(in[0].Int()),
				int(in[1].Int())),
		}
	}

	fn := reflect.ValueOf(fptr).Elem()

	v := reflect.MakeFunc(fn.Type(), swap)

	fn.Set(v)
}
