package sliding_time_window

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSlidingWindow(t *testing.T) {
	limiter := NewSliding(100*time.Millisecond, time.Second, 10)
	for i := 0; i < 5; i++ {
		assert.False(t, limiter.LimitTest())
	}
	time.Sleep(100 * time.Millisecond)
	for i := 0; i < 5; i++ {
		assert.False(t, limiter.LimitTest())
	}
	assert.True(t, limiter.LimitTest())

	for _, v := range limiter.windows {
		fmt.Println(v.timestamp, v.count)
	}

	fmt.Println("moments later...")
	time.Sleep(time.Second)
	for i := 0; i < 7; i++ {
		assert.False(t, limiter.LimitTest())
	}
	for _, v := range limiter.windows {
		fmt.Println(v.timestamp, v.count)
	}
}
