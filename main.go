package main

import (
	"go-todo/controllers"
	"go-todo/models"
	"go-todo/services"

	"os"

	"github.com/joho/godotenv"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Init
	app := iris.New()
	app.Logger().SetLevel("debug")

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		app.Logger().Fatalf("Error loading .env file")
	}

	// Data needed for connection string
	dbhost := os.Getenv("DB_HOST")
	dbname := "gotodo"
	dbuser := os.Getenv("DB_USER")
	dbpassword := os.Getenv("DB_PASSWORD")
	dbport := "3306"

	// Connect to the database
	connstring := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8&parseTime=True&loc=Local"
	println(connstring)
	db, err := gorm.Open(mysql.Open(connstring), &gorm.Config{})
	if err != nil {
		app.Logger().Fatalf("Error while connecting to database: %v", err)
		return
	}

	// Migration
	db.AutoMigrate(&models.Task{})

	// Setup MVC
	taskService := services.NewService(db)
	tasks := mvc.New(app.Party("/tasks"))
	tasks.Register(taskService)
	tasks.Handle(new(controllers.TaskController))

	// Run server
	app.Run(
		iris.Addr(":3000"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)
}
