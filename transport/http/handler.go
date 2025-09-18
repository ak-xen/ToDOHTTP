package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ak-xen/ToDOHTTP/internal/task"
)

type Handler struct {
	repo task.Repository
}

func NewHandler(repo task.Repository) *Handler {
	return &Handler{repo}
}
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Неизвестный метод CRUD!"))
	w.Header().Set("Content-Type", "application/json")
}
func (h *Handler) SaveTask(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}
	title := r.URL.Query().Get("title")
	err := h.repo.Save(title)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
}

func (h *Handler) GetTasks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	tasks, err := h.repo.GetList()
	if err != nil {
		http.Error(w, "Не возможно прочитать данные", http.StatusMethodNotAllowed)
	}

	lenResponse := len(tasks)

	if lenResponse == 0 {
		_, err = w.Write([]byte("Нет активных задач"))
		if err != nil {
			return
		}
	}
	resp := make(map[int]string)

	for _, v := range tasks {
		resp[v.ID] = v.Title
	}
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		return
	}

}

func (h *Handler) FindId(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")

	idInt, _ := strconv.Atoi(id)

	taskF, err := h.repo.FindByID(idInt)
	if err != nil {
		http.Error(w, "ID не найден", http.StatusBadRequest)
		return
	}

	resp, err := json.Marshal(taskF)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")

	_, err = w.Write(resp)
	if err != nil {
		return
	}

}

func (h *Handler) UpdateTask(w http.ResponseWriter, r *http.Request) {

	var t task.Task

	err := json.NewDecoder(r.Body).Decode(&t)

	_, err = h.repo.FindByID(t.ID)

	if err != nil {
		http.Error(w, "Не существующий ID", http.StatusBadRequest)
		return
	}
	err = h.repo.Update(t)
	if err != nil {
		return
	}
}

func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Метод не поддерживается", http.StatusBadRequest)
		return
	}

	idStr := r.URL.Query().Get("id")

	id, _ := strconv.Atoi(idStr)

	_, err := h.repo.FindByID(id)
	if err != nil {
		http.Error(w, "ID не найден", http.StatusBadRequest)
		return
	}

	err = h.repo.Delete(id)
	if err != nil {
		return

	}
}
