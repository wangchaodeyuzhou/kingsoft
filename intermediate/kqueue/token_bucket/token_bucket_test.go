package token_bucket

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestTokenBucket(t *testing.T) {
	// 初始化令牌桶
	capacity := 10
	fillInterval := time.Second
	tokensPerFill := 2
	tb := NewTokenBucket(int32(capacity), int32(tokensPerFill), fillInterval)

	go tb.startFillTimer()

	// 测试并发请求
	wg := sync.WaitGroup{}
	concurrentRequests := 20
	var taken int32
	for i := 0; i < concurrentRequests; i++ {
		wg.Add(1)
		go func() {
			if tb.GetTakeWithoutTimer(1) {
				atomic.AddInt32(&taken, 1)
			}
			wg.Done()
		}()
	}
	wg.Wait()

	// 测试限流情况
	if tb.GetTake(1) {
		t.Errorf("Unexpectedly took token_bucket when bucket is empty")
	}
}

func TestTokenBucket_Concurrency(t *testing.T) {
	// 初始化令牌桶
	capacity := 10
	fillInterval := time.Millisecond * 500
	tokensPerFill := 2
	tb := NewTokenBucket(int32(capacity), int32(tokensPerFill), fillInterval)

	go tb.startFillTimer()

	// 并发请求
	wg := sync.WaitGroup{}
	concurrentRequests := 20
	var taken int32
	for i := 0; i < concurrentRequests; i++ {
		wg.Add(1)
		go func() {
			if tb.GetTakeWithoutTimer(1) {
				atomic.AddInt32(&taken, 1)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	// 检查是否限流
	if int(taken) > capacity {
		t.Errorf("Too many tokens taken, expected at most %d, got %d", capacity, taken)
	}
}

// Benchmark基准测试
func BenchmarkTokenBucket(b *testing.B) {
	capacity := 1000
	fillInterval := time.Millisecond
	tokensPerFill := 100
	tb := NewTokenBucket(int32(capacity), int32(tokensPerFill), fillInterval)

	// 启动定时器
	tb.startFillTimer()

	b.ResetTimer()            // 重置计时器
	concurrentRequests := 100 // 并发请求数量
	var wg sync.WaitGroup
	wg.Add(concurrentRequests)
	for i := 0; i < concurrentRequests; i++ {
		go func() {
			for j := 0; j < b.N; j++ {
				tb.GetTake(1)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func BenchmarkGetTakeWithoutTimer(b *testing.B) {
	capacity := 1000
	fillInterval := time.Millisecond
	tokensPerFill := 100
	tb := NewTokenBucket(int32(capacity), int32(tokensPerFill), fillInterval)

	b.ResetTimer()
	concurrentRequests := 100 // 并发请求数量
	var wg sync.WaitGroup
	wg.Add(concurrentRequests)
	for i := 0; i < concurrentRequests; i++ {
		go func() {
			for j := 0; j < b.N; j++ {
				tb.GetTakeWithoutTimer(1)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
