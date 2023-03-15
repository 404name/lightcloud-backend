// Package rate provides a rate limiter hooker,
// this package is based on golang.org/x/time/rate
package middleware

import (
	"sync"
	"time"

	"github.com/FloatTech/ttl"
)

// 限流器管理器
type LimiterManager[K comparable] struct {
	// 使用带ttl的缓存管理限流器方便清楚不用的
	limiters *ttl.Cache[K, *Limiter]
	interval time.Duration
	burst    int
}

// NewManager ..
func NewManager[K comparable](interval time.Duration, burst int) *LimiterManager[K] {
	return &LimiterManager[K]{
		// 这里设置默认过期时间就是[令牌数*生成令牌间隔]
		// 比如10秒只能请求2次，那么过了10秒没请求那这个限流器就可以清除了
		// ttl.NewCache每分钟会自动清理过期的防止内存占用
		limiters: ttl.NewCache[K, *Limiter](interval * time.Duration(burst)),
		interval: interval,
		burst:    burst,
	}
}

// Load ...
func (l *LimiterManager[K]) Load(key K) *Limiter {
	// 这里的Get会给限流器刷新生命周期为now
	// 如果已经过期或者没有 就会新建
	if val := l.limiters.Get(key); val != nil {
		return val
	}
	val := NewLimiter(l.interval, l.burst)
	l.limiters.Set(key, val)
	return val
}

// Limiter controls the frequency of handling events.
type Limiter struct {
	sync.Mutex
	limit    float64
	tokens   float64
	burst    int
	lastTime time.Time
}

// Tokens returns the left token of limiter.
func (lim *Limiter) Tokens() float64 {
	return lim.tokens
}

// 创建一个新的Limiter
func NewLimiter(interval time.Duration, burst int) *Limiter {
	return &Limiter{
		limit:    every(interval),
		burst:    burst,
		tokens:   float64(burst),
		lastTime: time.Now(),
	}
}

// 尝试获取一个令牌
func (lim *Limiter) Acquire() bool {
	return lim.AcquireN(1)
}

// 获取具体逻辑
func (lim *Limiter) AcquireN(n int) bool {
	lim.Lock()
	defer lim.Unlock()
	// 更新上一次和这次间隔生成的token
	lim.advance(time.Now())
	// 尝试获取令牌够不够
	nf := float64(n)

	// 虽然令牌是float类型，但其实还是看它在什么时候能达到整数nf
	if lim.tokens >= nf {
		lim.tokens -= nf
		return true
	}
	return false
}

// 更新上一次和这次间隔生成的token
func (lim *Limiter) advance(now time.Time) {
	last := lim.lastTime
	// 获取间隔时间
	elapsed := now.Sub(last)

	// 间隔时间最多可以生成令牌
	if maxElapsed := lim.durationFromTokens(float64(lim.burst) - lim.tokens); elapsed > maxElapsed {
		elapsed = maxElapsed
	}

	// 时间转换成token个数(float)
	// 为什么令牌不是整数个呢？因为时间不是整秒计算的
	delta := lim.tokensFromDuration(elapsed)
	tokens := lim.tokens + delta

	// 更新内容
	lim.tokens = tokens
	lim.lastTime = now
}

func every(interval time.Duration) float64 {
	return 1 / interval.Seconds()
}

func (lim *Limiter) durationFromTokens(tokens float64) time.Duration {
	seconds := tokens / lim.limit
	return time.Nanosecond * time.Duration(1e9*seconds)
}

func (lim *Limiter) tokensFromDuration(d time.Duration) float64 {
	sec := float64(d/time.Second) * lim.limit
	nSec := float64(d%time.Second) * lim.limit
	return sec + nSec/1e9
}
