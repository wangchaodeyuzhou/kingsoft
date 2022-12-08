package reflect

import (
	"errors"
	"fmt"
	"reflect"
)

var gmMethods = reflect.ValueOf(&GmMethod{})

type GmMethod struct{}

func (g *GmMethod) Hello() {
	fmt.Println("hello")
}

func (g *GmMethod) Say() {
	fmt.Println("say")
}

func CallGmMethod(methodName string) error {
	methodInfo := gmMethods.MethodByName(methodName)
	if !methodInfo.IsValid() {
		return errors.New(fmt.Sprintf("no such gm method %s", methodName))
	}

	_ = methodInfo.Call([]reflect.Value{})
	fmt.Println("fine")
	return nil
}
