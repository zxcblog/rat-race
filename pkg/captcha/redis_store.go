package captcha

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisStore struct {
	redis      *redis.Client
	expireTime time.Duration
	prefix     string
}

func NewRedisStore(redis *redis.Client, expireTime time.Duration, prefix string) *RedisStore {
	return &RedisStore{redis: redis, expireTime: expireTime, prefix: prefix}
}

func (r RedisStore) GetId(id string) string {
	if len(r.prefix) < 1 {
		return id
	}
	return fmt.Sprintf("%s:%s", r.prefix, id)
}

// Set 将验证码存储到redis中, 并设置验证码过期时间为半个小时
func (r RedisStore) Set(id string, value string) error {
	return r.redis.SetNX(context.Background(), r.GetId(id), value, r.expireTime).Err()
}

// Get 通过id从redis中获取验证码
// clear 为 true 时会删除验证码
func (r RedisStore) Get(id string, clear bool) string {
	ctx := context.Background()
	id = r.GetId(id)

	val := r.redis.Get(ctx, id).Val()
	if clear {
		r.redis.Del(ctx, id)
	}

	return val
}

func (r RedisStore) Verify(id, answer string, clear bool) bool {
	if id == "" || answer == "" {
		return false
	}

	v := r.Get(id, clear)
	return v != "" && v == answer
}
