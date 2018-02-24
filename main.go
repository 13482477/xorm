package main

import (
	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"
	"time"
	//"github.com/shopspring/decimal"
	_ "fmt"
	"fmt"
)

type User struct {
	Id       int64
	Name     string
	Salt     string
	Age      int
	Password string    `xorm:"varchar(200)"`
	Created  time.Time `xorm:"created"`
	Updated  time.Time `xorm:"updated"`
	//Amount   decimal.Decimal `xorm:"Decimal(20,2)"`
}

func main() {

	engine, error := xorm.NewEngine("mysql", "root:password@tcp(localhost:3306)/xorm")
	if error != nil {
		panic(error)
	}

	//engine.Sync2(new(User))

	//engine.Insert(User{3, "lizhiqiang", "lizhiqiang", 30,"password", time.Now(), time.Now()})

	user := &User{Id:1}
	engine.Get(user)

	fmt.Printf("id=%s, name=%s", user.Id, user.Name)

}
