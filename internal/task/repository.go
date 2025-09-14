package task

type Repo struct {
	data map[int]Task
}

func NewRepo() *Repo {
	return &Repo{}
}

func (r *Repo) Save(task Task) error {
	r.data[task.ID] = task
	return nil
}

func (r *Repo) GetList() ([]string, error) {

	tasks := make([]string, 0)
	for _, value := range r.data {
		tasks = append(tasks, value.Title)
	}
	return tasks, nil
}

func (r *Repo) FindById(id int) (Task, error) {
	return r.data[id], nil
}

func (r *Repo) Update(task Task) error {
	r.data[task.ID] = task
	return nil
}

func (r *Repo) Delete(id int) error {
	delete(r.data, id)
	return nil
}
