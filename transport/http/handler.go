package http

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

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

	tasks, err := h.repo.GetList()
	if err != nil {
		http.Error(w, "Не возможно прочитать данные", http.StatusMethodNotAllowed)
	}

	combined := strings.Join(tasks, "\n")

	resp := json.RawMessage(combined)

	lenResponse := len(resp)

	if lenResponse == 0 {
		_, err = w.Write([]byte("Нет активных задач"))
		if err != nil {
			return
		}
	}

	_, err = w.Write(resp)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
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
		http.Error(w, "ID не найден", http.StatusMethodNotAllowed)
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
	if r.Method != http.MethodPut {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
	}

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Параметр id обязателен", http.StatusBadRequest)
		return
	}

	id, _ := strconv.Atoi(idStr)

	fTask, err := h.repo.FindByID(id)

	if err != nil {
		http.Error(w, "Не существующий ID", http.StatusBadRequest)
		return
	}

	title := r.URL.Query().Get("title")
	fTask.Title = title
	err = h.repo.Update(fTask)
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

	err := h.repo.Delete(id)
	if err != nil {
		return

	}
}
