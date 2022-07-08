package main

import (
	"os"

	"github.com/kataras/iris/v12"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")

	// data needed for connection string
	dbhost := "database"
	dbname := "gotodo"
	dbuser := os.Getenv("DB_USER")
	dbpassword := os.Getenv("EB_PASSWORD")
	dbport := "3306"

	db, err := gorm.Open("mysql", dbuser+":"+dbpassword+"@tcp("+dbhost+":"+dbport+")/"+dbname+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		app.Logger().Fatalf("error while loading the tables: %v", err)
		return
	}

	app.Run(
		iris.Addr(":8080"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)
}
