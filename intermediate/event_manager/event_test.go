package event_manager

import (
	"context"
	"fmt"
	"testing"
)

type LoginEvent struct {
	OldLvl int32
	NewLvl int32
}

//func ToLogin(ctx context.Context, event *LoginEvent) error {
//	fmt.Printf("role login %v", event)
//	fmt.Println("to Login")
//	return nil
//}
//
//func ToCreate(ctx context.Context, event *LoginEvent) error {
//	fmt.Printf("role create %v", event)
//	fmt.Println("to create")
//	return errors.New("test")
//}

func TestEvent(t *testing.T) {
	err := Trigger(context.TODO(), &LoginEvent{
		OldLvl: 1,
		NewLvl: 2,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	Trigger(context.Background(), &LoginEvent{
		OldLvl: 2,
		NewLvl: 3,
	})
}
