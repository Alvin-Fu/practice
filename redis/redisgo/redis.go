package redisgo

import (
	"github.com/garyburd/redigo/redis"
	"github.com/pkg/errors"
	"log"
	"time"
	"timerevent"
)

const TimerTypeRedisBreath = 10001

type RedisPool struct {
	pool *redis.Pool
}

func NewRedisPool(host, port, password string) (*RedisPool, error) {
	redisPool := &redis.Pool{
		MaxIdle:     500,
		MaxActive:   0,
		IdleTimeout: 300 * time.Second,
		Dial: func() (redis.Conn, error) {
			options := make([]redis.DialOption, 0)
			if password != "" {
				options = append(options, redis.DialPassword(password))
			}
			options = append(options, redis.DialReadTimeout(5*time.Second))
			options = append(options, redis.DialConnectTimeout(5*time.Second))
			dial, err := redis.Dial("tcp", host+":"+port, options...)
			if err != nil {
				log.Printf("redis err: %v", err)
				return nil, err
			}
			return dial, nil
		},
	}
	c := redisPool.Get()
	defer c.Close()
	if c.Err() != nil {
		log.Printf("redis pool err!")
	}
	return &RedisPool{
		pool: redisPool,
	}, c.Err()
}

func (rp *RedisPool) Pool() *redis.Pool {
	return rp.pool
}

func (rp *RedisPool) SetPool(p *redis.Pool) {
	rp.pool = p
}

func (p *RedisPool) Breath(svrName, host, port, password string, backFunc timerevent.CallBackFunc) {
	timerevent.StartTimer(svrName+host+port, TimerTypeRedisBreath, 1*time.Second,
		backFunc, nil, -1)
}

func (p *RedisPool) Do(commandName string, args ...interface{}) (reply interface{}, err error) {
	if commandName == "" || len(args) < 1 {
		return nil, errors.Errorf("Invalid parameters")
	}
	if p.pool == nil {
		return nil, errors.Errorf("redis pool nil")
	}
	c := p.pool.Get()
	defer c.Close()
	log.Printf("Args number: %d", len(args))
	return c.Do(commandName, args...)
}
