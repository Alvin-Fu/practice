package main

import (
    "github.com/8treenet/freedom"
    "practice/testfreedom/user/server/conf"
    "github.com/jinzhu/gorm"
    "time"
    "github.com/go-redis/redis"
)

func main(){
    app := freedom.NewApplication()
    installDB(app)
    installRedis(app)
}

func installDB(app freedom.Application){
    app.InstallDB(func()  interface{}{
        conf := conf.Get().DB
        db, err := gorm.Open("mysql", conf.Addr)
        if err != nil {
            freedom.Logger().Fatalf(err.Error())
        }
        db = db.Debug()
        db.DB().SetMaxIdleConns(conf.MaxIdleConns)
        db.DB().SetMaxOpenConns(conf.MaxOpenConns)
        db.DB().SetConnMaxLifetime(time.Duration(conf.ConnMaxLifeTime) * time.Second)
        return db
    })
}

func installRedis(app freedom.Application){
    app.InstallRedis(func()redis.Cmdable{
        conf := conf.Get().Redis
        opts := &redis.Options{
            Addr: conf.Addr,
            Password: conf.Password,
            DB: conf.DB,
        }
        client := redis.NewClient(opts)
        if err := client.Ping().Err(); err != nil {
            freedom.Logger().Fatalf(err.Error())
        }
        return client
    })
}

func installMiddleware(app freedom.Application){}
