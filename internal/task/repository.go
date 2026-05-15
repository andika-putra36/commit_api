package task

import "gorm.io/gorm"

type Repository interface {
	GetTasks() ([]GetTasksResponse, error)
	GetTask(id int) (GetTaskResponse, error)
	CreateTask(input CreateTaskRequest) error
	UpdateTask(id int, input UpdateTaskRequest) error
	DeleteTask(id int) error
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

func (r *repository) GetTask(id int) (GetTaskResponse, error) {
	var response GetTaskResponse

	err := r.db.Raw(`SELECT * FROM get_task(?)`, id).Scan(&response).Error
	if err != nil {
		return response, err
	}
	return response, nil
}

func (r *repository) CreateTask(input CreateTaskRequest) error {
	err := r.db.Exec(
		`CALL create_task(?, ?, ?, ?)`,
		input.Title,
		input.Subtitle,
		input.StartAt,
		input.FinishAt,
	).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) UpdateTask(id int, input UpdateTaskRequest) error {
	err := r.db.Exec(
		`CALL update_task(?, ?, ?, ?, ?, ?)`,
		id,
		input.Title,
		input.Subtitle,
		input.StartAt,
		input.FinishAt,
		input.IsDone,
	).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) DeleteTask(id int) error {
	err := r.db.Exec(
		`CALL delete_task(?)`,
		id,
	).Error
	if err != nil {
		return err
	}
	return nil
}
