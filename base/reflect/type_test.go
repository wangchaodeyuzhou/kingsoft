package reflect

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

func TestType(t *testing.T) {
	var u Admin
	a := reflect.TypeOf(u)

	if a.Kind() == reflect.Struct {
		fmt.Println("struct")
	}

	for i, n := 0, a.NumField(); i < n; i++ {
		f := a.Field(i)
		fmt.Println(f.Type, f.Name)
	}
}

func TestAdmin(t *testing.T) {
	u := new(Admin)

	a := reflect.TypeOf(u)

	// 如果是指针，应该先使⽤用 Elem 方法获取 ⽬目标类型，指针本⾝身是没有字段成员的
	if a.Kind() == reflect.Ptr {
		a = a.Elem()
	}

	for i, n := 0, a.NumField(); i < n; i++ {
		f := a.Field(i)
		fmt.Println(f.Type, f.Name)
	}
}

func TestStu_ToString(t *testing.T) {
	var tea Tea
	methods := func(a reflect.Type) {
		for i, n := 0, a.NumMethod(); i < n; i++ {
			m := a.Method(i)
			fmt.Println(m.Type, m.Name)
		}
	}

	fmt.Println("--- value interface ---")

	methods(reflect.TypeOf(tea))

	fmt.Println("--- pointer interface ---")

	methods(reflect.TypeOf(&tea))
}

func TestUStu(t *testing.T) {
	var u UAdmin
	m := reflect.TypeOf(u)

	f, _ := m.FieldByName("title")
	fmt.Println(f.Type, f.Name)

	f, _ = m.FieldByName("UStu")
	fmt.Println(f.Type, f.Name)

	f, _ = m.FieldByName("Username")
	fmt.Println(f.Type, f.Name)

	// UAdmin[0] -> UStu[1] -> age
	f = m.FieldByIndex([]int{0, 1})
	fmt.Println(f.Type, f.Name)
}

func TestUserModel(t *testing.T) {

	var u UserModel
	a := reflect.TypeOf(u)

	f, _ := a.FieldByName("Name")
	fmt.Println(f.Type, f.Name)

	fmt.Println(f.Tag)

	fmt.Println(f.Tag.Get("field"))

	fmt.Println(f.Tag.Get("type"))
}

func TestChan(t *testing.T) {
	c := reflect.ChanOf(reflect.SendDir, String)
	fmt.Println(c)

	m := reflect.MapOf(String, Int)
	fmt.Println(m)

	s := reflect.SliceOf(Int)
	fmt.Println(s)

	aa := struct{ Name string }{}
	p := reflect.PtrTo(reflect.TypeOf(aa))
	fmt.Println(p)

}

func TestData_String(t *testing.T) {
	var d *Data
	a := reflect.TypeOf(d)

	fmt.Println(a.Size(), a.Align())

	it := reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
	fmt.Println(a.Implements(it))
}

type KK struct {
	Name string
	age  int
}

func TestValue(t *testing.T) {
	v := reflect.ValueOf(KK{Name: "Jack", age: 1213})

	fmt.Println(v.FieldByName("Name").Interface())

	f := v.FieldByName("age")
	if f.CanInterface() {
		fmt.Println("cdscds")
		fmt.Println(f.Interface())
	} else {
		fmt.Println(f.Int())
	}

}

func TestPionter(t *testing.T) {
	var p *int
	var x any = p

	v := reflect.ValueOf(p)
	fmt.Println(x == nil)

	fmt.Println(v.Kind(), v.IsNil())

}

func TestSKK(t *testing.T) {
	u := KK{"Jack", 123}

	v := reflect.ValueOf(u)
	p := reflect.ValueOf(&u)

	fmt.Println(v.CanSet(), v.FieldByName("Name").CanSet())
	fmt.Println(p.CanSet(), p.Elem().FieldByName("Name").CanSet())
}

func TestSkkd(t *testing.T) {
	u := KK{Name: "csd", age: 21342}

	p := reflect.ValueOf(&u).Elem()

	p.FieldByName("Name").SetString("Tom")

	f := p.FieldByName("age")
	fmt.Println(f.CanSet())

	fmt.Println(u, p.Interface().(KK))
	fmt.Println(&u.age)
	if f.CanAddr() {
		age := (*int)(unsafe.Pointer(f.UnsafeAddr()))
		fmt.Println("age addr:", age)
		*age = 88
	}

	fmt.Println(u, p.Interface().(KK))

}

func TestLKKD(t *testing.T) {
	s := make([]int, 0, 10)
	v := reflect.ValueOf(&s).Elem()

	v.SetLen(2)
	v.Index(0).SetInt(1)
	v.Index(1).SetInt(2)

	fmt.Println(s, v.Interface())
	v2 := reflect.Append(v, reflect.ValueOf(300))

	v2 = reflect.AppendSlice(v2, reflect.ValueOf([]int{11, 22}))

	fmt.Println(v2.Interface(), s)

	fmt.Println("=============")

	m := map[string]int{"a": 1}
	v = reflect.ValueOf(&m).Elem()

	fmt.Println(m)
	v.SetMapIndex(reflect.ValueOf("a"), reflect.ValueOf(12))
	v.SetMapIndex(reflect.ValueOf("b"), reflect.ValueOf(32))

	fmt.Println(m, v.Interface())
}

func info(m reflect.Method) {
	t := m.Type

	fmt.Println(m.Type, m.Name)

	for i, n := 0, t.NumIn(); i < n; i++ {
		fmt.Printf(" in[%d] %v\n", i, t.In(i))
	}

	for i, n := 0, t.NumOut(); i < n; i++ {
		fmt.Printf(" out[%d] %v\n", i, t.Out(i))
	}

}

func TestInfo(t *testing.T) {
	d := new(IData)

	y := reflect.TypeOf(d)

	tt, _ := y.MethodByName("Test")
	info(tt)

	sum, _ := y.MethodByName("Sum")
	info(sum)
}

// 动态调⽤用⽅方法很简单，按 In 列表准备好所需参数即可 (不包括 receiver)。
func TestInfo1(t *testing.T) {
	d := new(IData)
	v := reflect.ValueOf(d)

	exec := func(name string, in []reflect.Value) {
		m := v.MethodByName(name)
		out := m.Call(in)

		for _, v := range out {
			fmt.Println(v.Interface())
		}
	}

	exec("Test", []reflect.Value{
		reflect.ValueOf(1),
		reflect.ValueOf(2),
	})

	fmt.Println("---------------")

	exec("Sum", []reflect.Value{
		reflect.ValueOf("result = %d"),
		reflect.ValueOf(1),
		reflect.ValueOf(2),
	})
}

func TestCallSlice(t *testing.T) {
	d := new(IData)
	v := reflect.ValueOf(d)

	m := v.MethodByName("Sum")
	in := []reflect.Value{
		reflect.ValueOf("result = %d"),
		reflect.ValueOf([]int{1, 2}),
	}

	out := m.CallSlice(in)
	fmt.Println(out)

	for _, o := range out {
		fmt.Println(o.Interface())
	}

}
