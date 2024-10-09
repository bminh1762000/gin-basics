package repository

import (
	"github.com/bminh1762000/jwt-auth-go/models"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(username, password string) (models.User, error)
}

type Task interface {
	CreateTask(task models.Task) (int, error)
	UpdateTask(task models.Task) error
	DeleteTask(taskId int) error
	GetAllTasks() ([]models.Task, error)
	GetTaskById(taskId int) (models.Task, error)
}

type Category interface {
	CreateCategory(category models.Category) (int, error)
	UpdateCategory(category models.Category) error
	DeleteCategory(categoryId int) error
	GetAllCategories() ([]models.Category, error)
	GetCategoryById(categoryId int) (models.Category, error)
}

type TaskCategory interface {
	CreateTaskCategory(taskId, categoryId int) error
	DeleteTaskCategory(taskId, categoryId int) error
	GetCategoriesByTaskId(taskId int) ([]models.Category, error)
	GetTasksByCategoryId(categoryId int) ([]models.Task, error)
}

type Repository struct {
	Authorization
	Task
	Category
	TaskCategory
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
