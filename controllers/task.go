package controllers

import (
	"go-todo/models"
	"go-todo/services"

	"github.com/kataras/iris/v12"
)

type TaskController struct {
	Ctx     iris.Context
	Service services.TaskService
}

// /task/all
func (c *TaskController) GetAll() {
	tasks := c.Service.SelectAll()
	c.Ctx.JSON(tasks)
}

// /task/id/:id
func (c *TaskController) GetIdBy(id uint) {
	task := c.Service.SelectByID(id)
	c.Ctx.JSON(task)
}

// /task/today
func (c *TaskController) GetToday() {
	tasks := c.Service.SelectForToday()
	c.Ctx.JSON(tasks)
}

// /task/tomorrow
func (c *TaskController) GetTomorrow() {
	tasks := c.Service.SelectForTomorrow()
	c.Ctx.JSON(tasks)
}

// /task/thisweek
func (c *TaskController) GetThisweek() {
	tasks := c.Service.SelectForThisWeek()
	c.Ctx.JSON(tasks)
}

// /task/insert
func (c *TaskController) PostInsert() {
	var task models.Task
	err := c.Ctx.ReadJSON(&task)
	if err != nil {
		c.Ctx.StatusCode(iris.StatusBadRequest)
		c.Ctx.JSON(iris.Map{"error": err.Error()})
		return
	}
	c.Service.Insert(task)
	c.Ctx.JSON(task)
}

// /task/update
func (c *TaskController) PostUpdate() {
	var task models.Task
	err := c.Ctx.ReadJSON(&task)
	if err != nil {
		c.Ctx.StatusCode(iris.StatusBadRequest)
		c.Ctx.JSON(iris.Map{"error": err.Error()})
		return
	}
	c.Service.Update(task)
	c.Ctx.JSON(task)
}

// /task/delete/:id
func (c *TaskController) GetDeleteBy(id uint) {
	c.Service.Delete(id)
}

// /task/markascomplete/:id
func (c *TaskController) GetMarkascompleteBy(id uint) {
	c.Service.MarkAsComplete(id)
}
