package time

import (
	"testing"
	"time"
)

func TestK(t *testing.T) {
	ticker := time.NewTicker(5 * time.Second)

	for range ticker.C {

	}

	ticker.Stop()

	ticker.Reset(10 * time.Second)
}
