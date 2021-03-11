/*
 * Copyright (c) 2021. Lorem ipsum dolor sit amet, consectetur adipiscing elit.
 * Morbi non lorem porttitor neque feugiat blandit. Ut vitae ipsum eget quam lacinia accumsan.
 * Etiam sed turpis ac ipsum condimentum fringilla. Maecenas magna.
 * Proin dapibus sapien vel ante. Aliquam erat volutpat. Pellentesque sagittis ligula eget metus.
 * Vestibulum commodo. Ut rhoncus gravida arcu.
 */

package mysql

import (
	"database/sql"
	"fmt"

	"time"

	"byrpc/sapi/putil/log"
	"domino/lib/timerevent"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

const TimerTypeMysqlBreath = 10002

type MysqlConf struct {
	Host     string
	Port     string
	User     string
	Password string
	IdleTime int64
	BaseName string
}

type MysqlSever struct {
	db *sql.DB
}

func NewMysqlSever(conf *MysqlConf) (*MysqlSever, error) {
	addr := conf.Host + ":" + conf.Port
	par := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", conf.User,
		conf.Password, addr, conf.BaseName)
	db, err := sql.Open("mysql", par)
	if err != nil {
		log.Fatal("open mysql err: %v, sql: %s", err, par)
		return nil, err
	}
	errP := db.Ping()
	if errP != nil {
		log.Fatal("ping err: %v, sql: %s", errP, par)
		return nil, errP
	}
	db.SetConnMaxLifetime(time.Duration(conf.IdleTime) * time.Second)
	db.SetMaxIdleConns(100)
	db.SetMaxOpenConns(100)
	return &MysqlSever{
		db: db,
	}, err
}

func (s *MysqlSever) GetMysqlDb() *sql.DB {
	return s.db
}

func (s *MysqlSever) Breath(conf *MysqlConf) {
	timerevent.StartTimer(conf.Host+conf.Port, TimerTypeMysqlBreath, 1*time.Second, func(eventId timerevent.EventID, private interface{}) {
		if err := s.db.Ping(); err != nil {
			plog.Fatalf("mysql db ping err: %v", err)
			s.db.Close()
			s, _ = NewMysqlSever(conf)
		}
	}, nil, -1)
}
