package models

type Task struct {
	TaskId      int    `json:"task_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"due_date"`
	UserId      int    `json:"user_id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type TaskInput struct {
	Title       *string `json:"title" binding:"required"`
	Description string  `json:"description"`
	DueDate     *string `json:"due_date" binding:"required"`
	UserId      int     `json:"user_id"`
}
