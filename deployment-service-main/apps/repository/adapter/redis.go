package adapter

import (
	"context"
	"deployment-service/logger"
	"time"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

var ctx = context.Background()

// Redis Operations

func (db *RedDB) GetKey(key string) ([]byte, error) {
	value, err := db.connection.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (db *RedDB) SetKey(key string, value []byte, expiry time.Duration) error {
	err := db.connection.Set(ctx, key, value, expiry).Err()
	if err != nil {
		return err
	}
	return nil
}

func (db *RedDB) HGet(key, field string) (string, error) {
	value, err := db.connection.HGet(ctx, key, field).Result()
	if err != nil {
		return "", err
	}
	return value, nil
}

func (db *RedDB) HLen(key string) (int64, error) {
	value, err := db.connection.HLen(ctx, key).Result()
	if err != nil {
		return 0, err
	}
	return value, nil
}

func (db *RedDB) HSet(key, field, value string, expiry time.Duration) error {
	err := db.connection.HSet(ctx, key, field, value).Err()
	if err != nil {
		return err
	}
	db.connection.Expire(ctx, key, expiry)
	return nil
}

func (db *RedDB) DelKey(key string) error {
	return db.connection.Del(ctx, key).Err()
}

func (db *RedDB) ZAdd(key string, member string, score float64) error {
	item := redis.Z{
		Score:  score,
		Member: member,
	}
	err := db.connection.ZAdd(ctx, key, &item).Err()
	return err
}

func (db *RedDB) ZAddBulk(key string, items []*redis.Z) error {
	err := db.connection.ZAdd(ctx, key, items...).Err()
	return err
}

func (db *RedDB) ZIncrBy(key string, member string, score float64) error {
	err := db.connection.ZIncrBy(ctx, key, score, member).Err()
	return err
}

func (db *RedDB) ZCard(key string) (int64, error) {
	out, err := db.connection.ZCard(ctx, key).Result()
	return out, err
}

func (db *RedDB) ZScore(key string, member string) (float64, error) {
	out, err := db.connection.ZScore(ctx, key, member).Result()
	return out, err
}

func (db *RedDB) ZMScore(key string, member []string) ([]float64, error) {
	out, err := db.connection.ZMScore(ctx, key, member...).Result()
	return out, err
}

func (db *RedDB) ZRangeByScore(key string, min, max string) ([]string, error) {
	opts := redis.ZRangeBy{
		Min: min,
		Max: max,
	}
	items, err := db.connection.ZRangeByScore(ctx, key, &opts).Result()
	if err != nil && err != redis.Nil {
		return nil, err
	}
	return items, nil
}

func (db *RedDB) ZRangeByScoreWithScores(key string, min, max string) ([]redis.Z, error) {
	opts := redis.ZRangeBy{
		Min: min,
		Max: max,
	}
	items, err := db.connection.ZRangeByScoreWithScores(ctx, key, &opts).Result()
	if err != nil && err != redis.Nil {
		return nil, err
	}
	return items, nil
}

func (db *RedDB) ZRangeWithScores(key string, min, max int64) ([]redis.Z, error) {
	items, err := db.connection.ZRangeWithScores(ctx, key, min, max).Result()
	if err != nil && err != redis.Nil {
		return nil, err
	}
	return items, nil
}

func (db *RedDB) ZRevRangeWithScores(key string, min, max int64) ([]redis.Z, error) {
	items, err := db.connection.ZRevRangeWithScores(ctx, key, min, max).Result()
	if err != nil && err != redis.Nil {
		return nil, err
	}
	return items, nil
}

func (db *RedDB) ZRevRangeByScoreWithScores(key string, min, max string) ([]redis.Z, error) {
	opts := redis.ZRangeBy{
		Min: min,
		Max: max,
	}
	items, err := db.connection.ZRevRangeByScoreWithScores(ctx, key, &opts).Result()
	if err != nil && err != redis.Nil {
		return nil, err
	}
	return items, nil
}

func (db *RedDB) ZRem(key string, member []string) error {
	err := db.connection.ZRem(ctx, key, member).Err()
	return err
}

func (db *RedDB) ZRemRangeByScore(key, min, max string) error {
	err := db.connection.ZRemRangeByScore(ctx, key, min, max).Err()
	return err
}

func (db *RedDB) ZCount(key, min, max string) (int64, error) {
	value, err := db.connection.ZCount(ctx, key, min, max).Result()
	if err != nil {
		return -1, err
	}
	return value, nil
}

func (db *RedDB) Pipeliner() redis.Pipeliner {
	p := db.connection.Pipeline()
	return p
}

func (db *RedDB) ExecPipeliner(pipe redis.Pipeliner) error {
	_, err := pipe.Exec(ctx)
	return err
}

func (db *RedDB) CheckIfLocked(key string) bool {
	out, err := db.GetKey(key)
	if err == redis.Nil || string(out) == "0" {
		return false
	}
	return true
}

func (db *RedDB) LockRedis(key string) {
	err := db.SetKey(key, []byte("1"), time.Second*15)
	if err != nil {
		logger.Logger.Error(logger.RedisError, zap.Any(logger.KEY_ERROR, err.Error()))
		return
	}
	return
}

func (db *RedDB) UnLockRedis(key string) {
	err := db.SetKey(key, []byte("0"), time.Second*15)
	if err != nil {
		logger.Logger.Error(logger.RedisError, zap.Any(logger.KEY_ERROR, err.Error()))
		return
	}
	return
}
