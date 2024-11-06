package cache

import (
	"context"
	"fmt"
	"github.com/redis/rueidis"
	"github.com/zakariawahyu/go-fiberavel/config"
	"time"
)

type Storage struct {
	Client rueidis.Client
}

type Rueidis interface {
	Set(key string, val []byte, exp time.Duration) error
	Get(key string) ([]byte, error)
	HGet(key, field string) (string, error)
	HGetAll(ctx context.Context, key string) (map[string]string, error)
	Delete(key string) error
	Reset() error
	Close() error
}

func NewRedis(cfg *config.Config) (*Storage, error) {
	client, err := rueidis.NewClient(rueidis.ClientOption{
		InitAddress: []string{fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port)},
		Username:    cfg.Redis.Username,
		Password:    cfg.Redis.Password,
		SelectDB:    cfg.Redis.SelectDB,
	})

	if err != nil {
		return nil, err
	}

	if err := client.Do(context.Background(), client.B().Ping().Build()).Error(); err != nil {
		return nil, err
	}

	return &Storage{
		Client: client,
	}, nil
}

func (s *Storage) Set(key string, val []byte, exp time.Duration) error {
	return s.Client.Do(context.Background(), s.Client.B().Set().Key(key).Value(string(val)).Build()).Error()
}

func (s *Storage) Get(key string) ([]byte, error) {
	val, err := s.Client.DoCache(context.Background(), s.Client.B().Get().Key(key).Cache(), 1*time.Hour).AsBytes()
	if err != nil && rueidis.IsRedisNil(err) {
		return nil, nil
	}
	return val, err
}

func (s *Storage) HGet(key, field string) (string, error) {
	val, err := s.Client.DoCache(context.Background(), s.Client.B().Hget().Key(key).Field(field).Cache(), 1*time.Hour).ToString()
	if err != nil && rueidis.IsRedisNil(err) {
		return "", nil
	}
	return val, err
}

func (s *Storage) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	val, err := s.Client.DoCache(ctx, s.Client.B().Hgetall().Key(key).Cache(), 1*time.Hour).AsStrMap()
	if err != nil && rueidis.IsRedisNil(err) {
		return nil, nil
	}
	return val, err

}

func (s *Storage) Delete(key string) error {
	return s.Client.Do(context.Background(), s.Client.B().Del().Key(key).Build()).Error()
}

func (s *Storage) Reset() error {
	return s.Client.Do(context.Background(), s.Client.B().Flushdb().Build()).Error()

}
func (s *Storage) Close() error {
	s.Client.Close()
	return nil
}
