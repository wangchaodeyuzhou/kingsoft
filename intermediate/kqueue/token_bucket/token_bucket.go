package token_bucket

import (
	"sync"
	"time"
)

/*
TokenBucket 令牌桶算法是一种流量控制算法: 最佳实践: 通过对每个用户发放令牌来限制用户的访问频率。
当用户请求到达时，如果令牌桶中有令牌，则允许用户访问，并从令牌桶中扣除一个令牌；否则，拒绝用户的请求。
令牌桶算法的结构模型包括一个令牌桶和一个计时器。
令牌桶用于存放令牌，
计时器用于定期向令牌桶中添加令牌。
*/
type TokenBucket struct {
	capacity      int32         // bucket 容量
	curTokenNum   int32         // 当前 token_bucket 数量
	tokenPerFill  int32         // 每次添加的 token_bucket 数目
	fillInternal  time.Duration // 添加 token_bucket 的时间间隔
	laterFillTime time.Time     // 上次添加 token_bucket 的时间
	mu            sync.Mutex    // 互斥锁，用于保护令牌桶的并发访问
	timer         *time.Timer   // 定时器, 定时向 token_bucket bucket put token_bucket
}

func NewTokenBucket(capacity int32, tokenPerFill int32, fillInternal time.Duration) *TokenBucket {
	return &TokenBucket{
		capacity:      capacity,
		tokenPerFill:  tokenPerFill,
		fillInternal:  fillInternal,
		laterFillTime: time.Now(),
	}
}

func (t *TokenBucket) GetTakeWithoutTimer(count int32) bool {
	t.mu.Lock()
	defer t.mu.Unlock()

	now := time.Now()

	// compute add token_bucket num
	addTokens := int32(now.Sub(t.laterFillTime)/t.fillInternal) * t.tokenPerFill

	// had addToken
	if addTokens > 0 {
		t.curTokenNum += addTokens
		// token_bucket bucket is full
		if t.curTokenNum > t.capacity {
			t.curTokenNum = t.capacity
		}
		t.laterFillTime = now
	}

	// 判断令牌是否足够
	if count <= t.curTokenNum {
		t.curTokenNum -= count
		return true
	}

	return false
}

// GetTake
//
//	@Description: GetTake 获取令牌
//	@receiver t
//	@param count take token_bucket 的数量
//	@return bool 是否获取到了令牌
func (t *TokenBucket) GetTake(count int32) bool {
	t.mu.Lock()
	defer t.mu.Unlock()

	// 判断令牌是否足够
	if count <= t.curTokenNum {
		t.curTokenNum -= count
		return true
	}

	return false
}

func (t *TokenBucket) startFillTimer() {
	t.mu.Lock()
	defer t.mu.Unlock()

	// 计算下次添加令牌的时间
	nextFillTime := t.laterFillTime.Add(t.fillInternal)
	// 计算距离下次添加令牌的时间间隔
	duration := nextFillTime.Sub(time.Now())

	// 如果定时器已经存在，则重置定时器，避免多次触发
	if t.timer != nil {
		t.timer.Stop()
	}

	// 添加令牌
	t.curTokenNum += t.tokenPerFill
	// 满了就不加了
	if t.curTokenNum > t.capacity {
		t.curTokenNum = t.capacity
	}

	t.laterFillTime = time.Now()

	t.timer = time.AfterFunc(duration, func() {
		// 定时器触发后，重新启动定时器
		t.startFillTimer()
	})
}
