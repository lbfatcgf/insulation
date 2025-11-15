package redisutil

import (
	"context"
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"insulation/server/base/pkg/config"

	"github.com/redis/go-redis/v9"
	"golang.org/x/exp/constraints"
)

var clinet *redis.Client

const (
	RedisKeyPrefix = "insulation"
)

// 生成redis key
// return RedisKeyPrefix:str1:str2:str3...
func GenKey(str ...string) string {
	var b strings.Builder
	b.WriteString(RedisKeyPrefix)
	for _, v := range str {
		b.WriteString(":")
		b.WriteString(v)
	}
	return b.String()
}

func InitRedis() error {
	link := config.Global().DataSource.Redis.DSN
	opt, err := redis.ParseURL(link)
	if err != nil {
		return err
	}
	clinet = redis.NewClient(opt)
	return nil
}

func GetRedis() *redis.Client {
	return clinet
}

func SetString(ctx context.Context, key, value string) error {
	return SetStringWithExpire(ctx, key, value, 0)
}

func SetStringWithExpire(ctx context.Context, key, value string, expire time.Duration) error {
	cmd := clinet.Set(ctx, key, value, expire)
	return cmd.Err()
}

func GetString(ctx context.Context, key string) (string, error) {
	cmd := clinet.Get(ctx, key)
	if cmd.Err() != nil {
		return "", cmd.Err()
	}
	return cmd.Result()
}

func SetInter[T constraints.Integer](ctx context.Context, key string, value T) error {
	return SetInterWithExprie(ctx, key, value, 0)
}

func SetInterWithExprie[T constraints.Integer](ctx context.Context, key string, value T, expire time.Duration) error {
	cmd := clinet.Set(ctx, key, value, expire)
	return cmd.Err()
}

func GetInter[T constraints.Integer](ctx context.Context, key string) (T, error) {
	cmd := clinet.Get(ctx, key)
	if cmd.Err() != nil {
		return 0, cmd.Err()
	}
	rv, err := strconv.ParseInt(cmd.Val(), 10, 64)
	if err != nil {
		return 0, err
	}
	return T(rv), nil
}

func SetFloat[T constraints.Float](ctx context.Context, key string, value T) error {
	return SetFloatWithExprie(ctx, key, value, 0)
}

func SetFloatWithExprie[T constraints.Float](ctx context.Context, key string, value T, exprie time.Duration) error {
	cmd := clinet.Set(ctx, key, value, 0)
	return cmd.Err()
}

func GetFloat[T constraints.Float](ctx context.Context, key string) (T, error) {
	cmd := clinet.Get(ctx, key)
	if cmd.Err() != nil {
		return 0, cmd.Err()
	}
	value := cmd.Val()
	rv, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, err
	}
	return T(rv), nil
}

func SetBool(ctx context.Context, key string, value bool) error {
	return SetBoolWithExprie(ctx, key, value, 0)
}

func SetBoolWithExprie(ctx context.Context, key string, value bool, exprie time.Duration) error {
	cmd := clinet.Set(ctx, key, value, 0)
	return cmd.Err()
}

func GetBool(ctx context.Context, key string) (bool, error) {
	cmd := clinet.Get(ctx, key)
	if cmd.Err() != nil {
		return false, cmd.Err()
	}
	value := cmd.Val()
	rv, err := strconv.ParseBool(value)
	if err != nil {
		return false, err
	}
	return rv, nil
}

func SetJsonString(ctx context.Context, key string, value any) error {
	return SetJsonStringWithExpire(ctx, key, value, 0)
}

func SetJsonStringWithExpire(ctx context.Context, key string, value any, expire time.Duration) error {
	js, err := json.Marshal(value)
	if err != nil {
		return err
	}
	cmd := clinet.Set(ctx, key, js, expire)
	return cmd.Err()
}

func GetJsonString(ctx context.Context, key string, value any) error {
	cmd := clinet.Get(ctx, key)
	if cmd.Err() != nil {
		return cmd.Err()
	}
	return json.Unmarshal([]byte(cmd.Val()), value)
}

func SetExprie(ctx context.Context, key string, expire time.Duration) error {
	cmd := clinet.Expire(ctx, key, expire)
	return cmd.Err()
}

func Incr(ctx context.Context, key string) error {
	cmd := clinet.Incr(ctx, key)
	return cmd.Err()
}

func Decr(ctx context.Context, key string) error {
	cmd := clinet.Decr(ctx, key)
	return cmd.Err()
}

func HasKey(ctx context.Context, key string) (bool, error) {
	cmd := clinet.Exists(ctx, key)
	if cmd.Err() != nil {
		return false, cmd.Err()
	}
	return cmd.Val() > 0, nil
}
