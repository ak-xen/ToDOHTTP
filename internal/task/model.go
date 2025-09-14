package task

type Task struct {
	ID    int
	Title string
	Done  bool
}

type Repository interface {
	Save(task Task) error
	GetList() ([]string, error)
	FindByID(id int) (Task, error)
	Update(task Task) error
	Delete(id int) error
}
