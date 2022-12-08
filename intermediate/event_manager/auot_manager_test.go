package event_manager

import (
	"os"
	"sort"
	"testing"
)

func TestEventManager(t *testing.T) {
	rootWd, _ := os.Getwd()
	println(rootWd)

	events := scanFuncName(eventRootDir, "event")
	if events != nil {
		sort.Strings(events)
		genEventRegister(events, eventRootDir+"/"+eventRegisterFile)
	}
}
