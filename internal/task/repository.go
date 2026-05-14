package task

import "gorm.io/gorm"

type Repository interface {
	GetTasks() ([]GetTasksResponse, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetTasks() ([]GetTasksResponse, error) {
	var responses []GetTasksResponse

	err := r.db.Raw(`SELECT * FROM get_tasks()`).Scan(&responses).Error
	if err != nil {
		return responses, err
	}
	return responses, nil
}
