package test

import (
	"fmt"
	"testing"
)

func TestFlag(t *testing.T) {
	HH()
	fmt.Println("Host:", GlobalConfig.Host)
	fmt.Println("Host:", GlobalConfig.ListenPort)
	fmt.Println("Host:", GlobalConfig.User)
	fmt.Println("Host:", GlobalConfig.Password)
	fmt.Println("Host:", GlobalConfig.Debug)
}
