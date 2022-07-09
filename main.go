package main

import (
	"go-todo/controllers"
	"go-todo/database"
	"go-todo/models"
	"go-todo/services"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func main() {
	// Init
	app := iris.New()
	app.Logger().SetLevel("debug")

	// Connect to database
	db, err := database.GetDatabase()
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
