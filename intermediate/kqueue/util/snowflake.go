package util

import (
	"github.com/sony/sonyflake"
)

var sf *sonyflake.Sonyflake

func InitIDGenerator(machineID uint16) {
	st := sonyflake.Settings{
		MachineID: func() (uint16, error) {
			return machineID, nil
		},
	}
	sf = sonyflake.NewSonyflake(st)
	if sf == nil {
		panic("sonyflake not created")
	}
}

func NextID() (uint64, error) {
	return sf.NextID()
}
