package middleware

import (
	"context"
	"errors"
	"net/http"
	"sync"
	"time"

	"go.uber.org/zap"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
)

var once sync.Once
var limiterManager *LimiterManager[string]

type LimitConfig struct {
	// GenerationKey 根据业务生成key 下面CheckOrMark查询生成
	GenerationKey func(c *gin.Context) string
	// 检查函数,用户可修改具体逻辑,更加灵活
	CheckOrMark func(key string, expire int, limit int) error
	// Expire key 过期时间
	Expire int
	// Limit 周期时间
	Limit int
}

func (l LimitConfig) LimitWithTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := l.CheckOrMark(l.GenerationKey(c), l.Expire, l.Limit); err != nil {
			c.JSON(http.StatusOK, gin.H{"code": response.ERROR, "msg": err})
			c.Abort()
			return
		} else {
			c.Next()
		}
	}
}

// DefaultGenerationKey 默认生成key
func DefaultGenerationKey(c *gin.Context) string {
	return "GVA_Limit" + c.ClientIP()
}

func DefaultCheckOrMark(key string, expire int, limit int) (err error) {
	// 没开启redis走内存ttl漏桶限流器
	if global.GVA_REDIS == nil {
		once.Do(func() {
			// 限流器
			limiterManager = NewManager[string](time.Duration(int64(expire)/int64(limit))*time.Second, limit)

		})
		if ok := limiterManager.Load(key).Acquire(); ok {
			return nil
		}
		return errors.New("请求太过频繁，请稍后再试")
	}
	//开启redis
	if err = SetLimitWithTime(key, limit, time.Duration(expire)*time.Second); err != nil {
		global.GVA_LOG.Error("limit", zap.Error(err))
	}
	return err
}

func DefaultLimit() gin.HandlerFunc {
	return LimitConfig{
		GenerationKey: DefaultGenerationKey,
		CheckOrMark:   DefaultCheckOrMark,
		Expire:        global.GVA_CONFIG.System.LimitTimeIP,
		Limit:         global.GVA_CONFIG.System.LimitCountIP,
	}.LimitWithTime()
}

// SetLimitWithTime 设置访问次数
func SetLimitWithTime(key string, limit int, expiration time.Duration) error {
	count, err := global.GVA_REDIS.Exists(context.Background(), key).Result()
	if err != nil {
		return err
	}
	if count == 0 {
		pipe := global.GVA_REDIS.TxPipeline()
		pipe.Incr(context.Background(), key)
		pipe.Expire(context.Background(), key, expiration)
		_, err = pipe.Exec(context.Background())
		return err
	} else {
		// 次数
		if times, err := global.GVA_REDIS.Get(context.Background(), key).Int(); err != nil {
			return err
		} else {
			if times >= limit {
				if t, err := global.GVA_REDIS.PTTL(context.Background(), key).Result(); err != nil {
					return errors.New("请求太过频繁，请稍后再试")
				} else {
					return errors.New("请求太过频繁, 请 " + t.String() + " 秒后尝试")
				}
			} else {
				return global.GVA_REDIS.Incr(context.Background(), key).Err()
			}
		}
	}
}
