package db

import (
	"fmt"
	"log"

	_ "github.com/godror/godror"
	"xorm.io/xorm"
)

const (
	driverName = "godror"
	host       = "1.1.1.1"
	port       = 1521
	user       = "admin"
	password   = "password"
	dbname     = "ORCL"
)

func Conn() *xorm.Engine {

	oralInfo := fmt.Sprintf("%s/%s@%s:%d/%s", user, password, host, port, dbname)
	engine, err := xorm.NewEngine(driverName, oralInfo)
	// engine.ShowSQL(true)

	if err != nil {
		log.Panic(err)
	}

	err = engine.Ping()

	if err != nil {
		panic(err)
	}
	// defer db.Close()

	fmt.Println("Connected to Oracle!")
	return engine
}
