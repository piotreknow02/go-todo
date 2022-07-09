package services

import (
	"go-todo/models"

	"time"

	"gorm.io/gorm"
)

//Date format
const YYYYMMDD = "2006-01-02"

// Types
type TaskService interface {
	SelectAll() []models.Task
	SelectByID(id uint) models.Task
	SelectForToday() []models.Task
	SelectForTomorrow() []models.Task
	SelectForThisWeek() []models.Task
	Insert(task models.Task)
	Update(task models.Task)
	Delete(id uint)
	MarkAsComplete(id uint)
}

type taskService struct {
	db *gorm.DB
}

// Constructor
func NewService(db *gorm.DB) TaskService {
	return &taskService{db: db}
}

// Data manipulation functions
func (r *taskService) SelectAll() []models.Task {
	var tasks []models.Task
	r.db.Find(&tasks)
	return tasks
}

func (r *taskService) SelectByID(id uint) models.Task {
	var task models.Task
	r.db.First(&task, id)
	return task
}

func (r *taskService) SelectForToday() []models.Task {
	var tasks []models.Task
	r.db.Where("expiry_date = ?", time.Now().Format(YYYYMMDD)).Find(&tasks)
	return tasks
}

func (r *taskService) SelectForTomorrow() []models.Task {
	var tasks []models.Task
	r.db.Where("expiry_date = ?", time.Now().AddDate(0, 0, 1).Format(YYYYMMDD)).Find(&tasks)
	return tasks
}

func (r *taskService) SelectForThisWeek() []models.Task {
	var tasks []models.Task
	r.db.Where("expiry_date <= ? AND expiry_date >= ?", time.Now().AddDate(0, 0, 7).Format(YYYYMMDD), time.Now().Format(YYYYMMDD)).Find(&tasks)
	return tasks
}

func (r *taskService) Insert(task models.Task) {
	r.db.Create(&task)
}

func (r *taskService) Update(task models.Task) {
	taskToUpdate := r.db.Find(&task.ID)
	if len(task.Title) > 0 {
		taskToUpdate.UpdateColumn("title", task.Title)
	}
	if len(task.Description) > 0 {
		taskToUpdate.UpdateColumn("description", task.Description)
	}
	if task.ExpiryDate.IsZero() {
		taskToUpdate.UpdateColumn("expiry_date", task.ExpiryDate)
	}
	taskToUpdate.UpdateColumn("complete", task.Complete)
}

func (r *taskService) Delete(id uint) {
	r.db.Delete(&models.Task{}, id)
}

func (r *taskService) MarkAsComplete(id uint) {
	var task models.Task
	r.db.First(&task, id)
	task.Complete = 100
	r.db.Save(&task)
}
