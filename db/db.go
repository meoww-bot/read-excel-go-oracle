package db

import (
	"fmt"
	"log"
	"read-excel-go-oracle/lib"

	"xorm.io/xorm"
)

var client *xorm.Engine = Conn()

func SyncTable() {

	err := client.Sync2(new(lib.Inventory))

	if err != nil {
		log.Panic(err)
	}
}

func AddInventory(obj []lib.Inventory) error {

	_, err := client.Insert(&obj)
	return err
}

func AddInventoryOne(obj lib.Inventory) error {

	_, err := client.Insert(&obj)
	if err != nil {
		fmt.Println(obj)
	}
	return err
}
