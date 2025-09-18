package task

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"status"`
}

type Repository interface {
	Save(title string) error
	GetList() ([]Task, error)
	FindByID(id int) (Task, error)
	Update(task Task) error
	Delete(id int) error
}
