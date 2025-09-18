package task

import (
	"errors"
	"strings"
)

type Repo struct {
	data    map[int]Task
	Counter int
}

func NewRepo() *Repo {
	return &Repo{data: make(map[int]Task), Counter: 0}
}

func (r *Repo) Save(title string) error {
	id := r.GetNextCount()
	r.data[id] = Task{ID: id, Title: strings.Trim(title, "\""), Done: true}
	return nil
}

func (r *Repo) GetNextCount() int {
	r.Counter++
	return r.Counter
}

func (r *Repo) GetList() ([]Task, error) {
	tasks := make([]Task, 0)
	for _, v := range r.data {
		tasks = append(tasks, v)
	}
	return tasks, nil
}

func (r *Repo) FindByID(id int) (Task, error) {
	if value, ok := r.data[id]; ok == true {
		return value, nil
	}
	return Task{}, errors.New("Not Found")
}

func (r *Repo) Update(task Task) error {
	r.data[task.ID] = task
	return nil
}

func (r *Repo) Delete(id int) error {
	delete(r.data, id)
	return nil
}
