package main

import (
	"github.com/judwhite/go-svc/svc"
	"os"
	"path/filepath"
	"syscall"

	"log"
	"practice/redis/redisgo"
	"time"
	"timerevent"
)

type config map[string]interface{}
type program struct {
	redisPool *redisgo.RedisPool
}

func main() {
	prg := &program{}
	if err := svc.Run(prg, syscall.SIGINT, syscall.SIGTERM); err != nil {
		log.Fatal(err)
	}
}

func (p *program) Init(env svc.Environment) error {
	if env.IsWindowsService() {
		dir := filepath.Dir(os.Args[0])
		return os.Chdir(dir)
	}
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	return nil
}

func (p *program) Start() error {
	var err error
	p.redisPool, err = redisgo.NewRedisPool("127.0.0.1", "6379", "")
	if err != nil {
		log.Fatalf("redis pool err: %v", err)
		os.Exit(1)
	}
	p.redisPool.Breath("", "127.0.0.1", "6379", "", p.CallBack)
	go func() {
		for {
			log.Println(p.redisPool.Do("set", time.Now().Unix(), "hello"))
			time.Sleep(500 * time.Millisecond)
		}
	}()
	return nil
}

func (p *program) Stop() error {
	return nil
}

func (pg *program) CallBack(eventId timerevent.EventID, private interface{}) {
	switch eventId {
	case redisgo.TimerTypeRedisBreath:
		p := pg.redisPool
		if p.Pool() == nil {
			return
		}
		c := p.Pool().Get()
		defer c.Close()
		if c.Err() != nil {
			log.Printf("conn err: %v", c.Err())
			err := p.Pool().Close()
			if err != nil {
				log.Printf("close redis pool err: %v", err)
			}
			pg.redisPool, _ = redisgo.NewRedisPool("127.0.0.1", "6379", "")
		}
	}
}
