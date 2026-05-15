package task

type Service interface {
	GetTasks() ([]GetTasksResponse, error)
	GetTask(id int) (GetTaskResponse, error)
	CreateTask(input CreateTaskRequest) error
	UpdateTask(id int, input UpdateTaskRequest) error
	DeleteTask(id int) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetTasks() ([]GetTasksResponse, error) {
	return s.repository.GetTasks()
}

func (s *service) GetTask(id int) (GetTaskResponse, error) {
	return s.repository.GetTask(id)
}

func (s *service) CreateTask(input CreateTaskRequest) error {
	return s.repository.CreateTask(input)
}

func (s *service) UpdateTask(id int, input UpdateTaskRequest) error {
	return s.repository.UpdateTask(id, input)
}

func (s *service) DeleteTask(id int) error {
	return s.repository.DeleteTask(id)
}
