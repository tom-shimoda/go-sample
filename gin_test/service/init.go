package service

import (
    "errors"
    "fmt"
    "github.com/go-xorm/xorm"
    "gin_test/model"
    "log"
)

var DbEngine *xorm.Engine

func init()  {
    driverName := "mysql"
    DsName := "user:password@(127.0.0.1:3306)/db?charset=utf8"
    err := errors.New("")
    DbEngine, err = xorm.NewEngine(driverName,DsName)
    if err != nil && err.Error() != ""{
        log.Fatal(err.Error())
    }
    DbEngine.ShowSQL(true)
    DbEngine.SetMaxOpenConns(2)
    DbEngine.Sync2(new(model.Book))
    fmt.Println("init data base ok")
}

