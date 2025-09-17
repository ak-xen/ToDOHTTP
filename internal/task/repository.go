package task

import "fmt"

type Repo struct {
	data    map[int]Task
	Counter int
}

func NewRepo() *Repo {
	return &Repo{}
}

func (r *Repo) Save(title string) error {
	id := r.Counter
	r.data[id] = Task{ID: id, Title: title, Done: true}
	fmt.Println(r.data)
	return nil
}

func (r *Repo) GetNextCount() (int, error) {
	r.Counter++
	return r.Counter, nil
}

func (r *Repo) GetList() ([]string, error) {

	tasks := make([]string, 0)
	for _, value := range r.data {
		tasks = append(tasks, value.Title)
	}
	return tasks, nil
}

func (r *Repo) FindByID(id int) (Task, error) {
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
