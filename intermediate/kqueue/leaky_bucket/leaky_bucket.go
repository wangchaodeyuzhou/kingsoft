package leaky_bucket

import (
	"sync"
	"sync/atomic"
	"time"
)

/*
LeakyBucket 漏桶
桶具有一定的容量，即最大容量的请求数，
当排队请求的数量超过桶的容量时，再进来的请求就直接过滤掉，不再被处理。
换句话说就是请求先在桶中排队，系统或服务只以一个恒定的速度从桶中将请求取出进行处理。
如果排队的请求数量超过桶能够容纳的最大数量，即桶装满了，则直接丢弃。
*/
type LeakyBucket struct {
	capacity     int32         // 桶的容量
	rate         time.Duration // 漏水的速率
	curBucketNum int32         // 当前桶内的令牌数量
	lastLeakTime time.Time     // 上一次漏水的时间
	timer        *time.Timer   // 定时器, 定时漏水
	mu           sync.Mutex    // 互斥锁, 用于对桶内令牌数量的并发访问控制
}

func NewLeakBucket(capacity int32, rate time.Duration) *LeakyBucket {
	return &LeakyBucket{
		capacity:     capacity,
		rate:         rate,
		lastLeakTime: time.Now(),
		timer:        time.NewTimer(rate),
	}
}

func (lb *LeakyBucket) startLeakyTimer() {
	go func() {
		for range lb.timer.C {
			lb.leak()
		}
	}()
}

func (lb *LeakyBucket) leak() {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(lb.lastLeakTime)

	// 漏水操作，移除需要漏掉的令牌
	tokensToLeak := int(elapsed / lb.rate)
	if tokensToLeak > 0 {
		if int32(tokensToLeak) >= lb.curBucketNum {
			atomic.StoreInt32(&lb.curBucketNum, 0)
		} else {
			atomic.AddInt32(&lb.curBucketNum, int32(-tokensToLeak))
		}
	}

	lb.lastLeakTime = now

	lb.timer.Reset(lb.rate)
}

func (lb *LeakyBucket) AddToken() bool {
	if lb.curBucketNum < lb.capacity {
		atomic.AddInt32(&lb.curBucketNum, 1)
		return true
	}
	return false
}
