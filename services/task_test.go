package services_test

import (
	"go-todo/database"
	"go-todo/models"
	"go-todo/services"
	"testing"
	"time"

	"gorm.io/gorm"
)

func Setup() *gorm.DB {
	db, err := database.GetDatabase()
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.Task{})
	return db
}

func TestSelectAll(t *testing.T) {
	db := Setup()
	tasks := services.NewService(db).SelectAll()
	t.Logf("PASSED Got %d tasks", len(tasks))
}

func TestSelectById(t *testing.T) {
	db := Setup()
	task := services.NewService(db).SelectByID(1)
	t.Logf("PASSED Got task with id 1 and title: %d", &task.Title)
}

func TestSelectForToday(t *testing.T) {
	db := Setup()
	tasks := services.NewService(db).SelectForToday()
	t.Logf("PASSED Got %d tasks for today", len(tasks))
}

func TestSelectForTomorrow(t *testing.T) {
	db := Setup()
	tasks := services.NewService(db).SelectForTomorrow()
	t.Logf("PASSED Got %d tasks for tomorrow", len(tasks))
}

func TestSelectForThisWeek(t *testing.T) {
	db := Setup()
	tasks := services.NewService(db).SelectForThisWeek()
	t.Logf("PASSED Got %d tasks for this week", len(tasks))
}

func TestInsert(t *testing.T) {
	db := Setup()
	task := models.Task{Title: "Test task", Description: "Test Task", ExpiryDate: time.Now().AddDate(0, 0, 1), Complete: 5}
	services.NewService(db).Insert(task)
	t.Logf("PASSED Inserted task with title: %d", &task.Title)
}

func TestUpdate(t *testing.T) {
	db := Setup()
	task := models.Task{Title: "Test task", Description: "Test Task", ExpiryDate: time.Now().AddDate(0, 0, 1), Complete: 5}
	services.NewService(db).Insert(task)
	task.Title = "Updated task"
	services.NewService(db).Update(task)
	t.Logf("PASSED Updated task with title: %d", &task.Title)
}

func TestDelete(t *testing.T) {
	db := Setup()
	task := models.Task{Title: "Test task", Description: "Test Task", ExpiryDate: time.Now().AddDate(0, 0, 1), Complete: 5}
	services.NewService(db).Insert(task)
	id := services.NewService(db).SelectByTitle(task.Title)[0].ID
	services.NewService(db).Delete(id)
	t.Logf("PASSED Deleted task with id: %d", id)
}

func TestMarkAsComplete(t *testing.T) {
	db := Setup()
	task := models.Task{Title: "Test task", Description: "Test Task", ExpiryDate: time.Now().AddDate(0, 0, 1), Complete: 5}
	services.NewService(db).Insert(task)
	id := services.NewService(db).SelectByTitle(task.Title)[0].ID
	services.NewService(db).MarkAsComplete(id)
	t.Logf("PASSED Marked task with id: %d marked as complete", id)
}
