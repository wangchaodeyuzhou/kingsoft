package leaky_bucket

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

func TestLeakyBucket_AddToken(t *testing.T) {
	// 创建一个容量为 5，速率为 1 秒的 LeakyBucket
	bucket := NewLeakBucket(5, time.Second)
	bucket.startLeakyTimer()

	// 添加 5 个令牌
	for i := 0; i < 5; i++ {
		assert.True(t, bucket.AddToken())
	}

	// 添加第 6 个令牌，超过了容量，返回 false
	assert.False(t, bucket.AddToken())

	// 等待 2 秒，让桶中的令牌漏掉 2 个
	time.Sleep(2 * time.Second)

	// 添加令牌，返回 true
	assert.True(t, bucket.AddToken())
}

func BenchmarkLeakyBucket_AddToken(b *testing.B) {
	bucket := NewLeakBucket(1000, time.Second) // 创建一个容量为1000，速率为1秒的LeakyBucket
	bucket.startLeakyTimer()

	var wg sync.WaitGroup
	wg.Add(b.N)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		go func() {
			defer wg.Done()
			bucket.AddToken()
		}()
	}

	wg.Wait()
}
