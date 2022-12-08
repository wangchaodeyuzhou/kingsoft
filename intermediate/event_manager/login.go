package event_manager

import (
	"context"
	"errors"
	"fmt"
)

type LoginTEvent struct {
	OldLvl int32
	NewLvl int32
}

// register:event
func ToLogin(ctx context.Context, event *LoginTEvent) error {
	fmt.Printf("role login %v", event)
	fmt.Println("to Login")
	return nil
}

// register:event
func ToCreate(ctx context.Context, event *LoginTEvent) error {
	fmt.Printf("role create %v", event)
	fmt.Println("to create")
	return errors.New("test")
}
