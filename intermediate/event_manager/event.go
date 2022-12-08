package event_manager

import (
	"fmt"
	"reflect"

	"git.shiyou.kingsoft.com/go/log"
)

func Trigger(ctx any, event any) error {
	return eventMgr.dispatchEvent(ctx, event)
}

func Register(fn any) bool {
	return eventMgr.register(fn)
}

var eventMgr = NewManager()

type Manager struct {
	listener map[string][]any
}

func NewManager() *Manager {
	return &Manager{
		listener: make(map[string][]any),
	}
}

func (m *Manager) register(fn any) bool {
	fnValue := reflect.ValueOf(fn)
	fmt.Println(fmt.Sprintf("fnValue : %v", fnValue))
	fnType := fnValue.Type()
	fmt.Println(fmt.Sprintf("fnType : %v", fnType))
	fnName := fnType.String()
	fmt.Println(fmt.Sprintf("fnName : %v", fnName))

	// 判断是不是 函数类型
	if fnType.Kind() != reflect.Func {
		panic(fmt.Sprintf("fn  = %v is not a func", fnName))
	}

	// 判断参数个数
	if fnType.NumIn() != 2 {
		panic(fmt.Sprintf("fn = %v param num not 2", fnName))
	}

	// 第二个参数为 eventType
	eventType := fnType.In(1)
	if eventType.Kind() != reflect.Pointer {
		panic(fmt.Sprintf("fn = %v, second param must be pointer", fnName))
	}

	outType := fnType.Out(0)
	if !outType.Implements(reflect.TypeOf((*error)(nil)).Elem()) {
		panic(fmt.Sprintf("fn = %v, return value type not error", fnName))
	}

	if !outType.AssignableTo(reflect.TypeOf((*error)(nil)).Elem()) {
		panic(fmt.Sprintf("fn = %v, return value type not error", fnName))
	}

	eventName := eventType.Elem().Name()
	fns := m.listener[eventName]

	// 防止重复注册
	for _, p := range fns {
		if &p == &fn {
			log.Error("duplicate register event handler", "event", eventType.Name(), "handler", fn)
			return false
		}
	}

	m.listener[eventName] = append(fns, fn)
	return true
}

// ctx, 和 event 是 Trigger 传过来的
func (m *Manager) dispatchEvent(ctx any, e any) error {
	eventName := reflect.TypeOf(e).Elem().Name()
	fmt.Println("eventName : ", eventName)
	fns := m.listener[eventName]
	for _, fn := range fns {
		method := reflect.ValueOf(fn)
		in := []reflect.Value{
			reflect.ValueOf(ctx),
			reflect.ValueOf(e),
		}
		result := method.Call(in)
		if result[0].IsNil() {
			continue
		}

		switch v := result[0].Interface().(type) {
		case error:
			return v
		default:
			log.Error("unKnown event result", "event", eventName, "result", v)
			return nil
		}
	}

	return nil
}
