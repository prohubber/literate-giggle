package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"main/project/internal/taskService"
	"main/project/internal/web/tasks"

	"github.com/gorilla/mux"
)

type Handler struct {
	Service *taskService.TaskService
}

func (h *Handler) GetTasksUserUserId(ctx context.Context, request tasks.GetTasksUserUserIdRequestObject) (tasks.GetTasksUserUserIdResponseObject, error) {
	// Получаем user_id из запроса
	userID := request.UserId

	// Обращаемся к сервису для получения задач пользователя
	userTasks, err := h.Service.GetTasksByUserID(uint(userID))
	if err != nil {
		return nil, err
	}

	// Формируем ответ
	response := tasks.GetTasksUserUserId200JSONResponse{}
	for _, tsk := range userTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
		}
		response = append(response, task)
	}

	return response, nil
}

func (h *Handler) GetTasks(_ context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	// Получение всех задач из сервиса
	allTasks, err := h.Service.GetAllTasks()
	if err != nil {
		return nil, err
	}

	// Создаем переменную респон типа 200джейсонРеспонс
	// Которую мы потом передадим в качестве ответа
	response := tasks.GetTasks200JSONResponse{}

	// Заполняем слайс response всеми задачами из БД
	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
		}
		response = append(response, task)
	}

	// САМОЕ ПРЕКРАСНОЕ. Возвращаем просто респонс и nil!
	return response, nil
}

func (h *Handler) PostTasks(_ context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	// Распаковываем тело запроса напрямую, без декодера!
	taskRequest := request.Body
	// Обращаемся к сервису и создаем задачу
	taskToCreate := taskService.Task{
		Task:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
		UserID: *taskRequest.UserId,
	}
	createdTask, err := h.Service.CreateTask(taskToCreate)

	if err != nil {
		return nil, err
	}
	// создаем структуру респонс
	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Task,
		IsDone: &createdTask.IsDone,
	}
	// Просто возвращаем респонс!
	return response, nil
}

func NewHandler(service *taskService.TaskService) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) PatchTasksId(ctx context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	// Распаковываем тело запроса
	taskRequest := request.Body
	// Обращаемся к сервису и частично обновляем задачу
	taskToUpdate := taskService.Task{
		Task:   *taskRequest.Task,   // Если поле не пустое, обновим
		IsDone: *taskRequest.IsDone, // Если поле не пустое, обновим
		UserID: *taskRequest.UserId,
	}
	updatedTask, err := h.Service.UpdateTaskByID(request.Id, taskToUpdate)

	if err != nil {
		return nil, err
	}

	// создаем структуру респонс
	response := tasks.PatchTasksId200JSONResponse{
		Id:     &updatedTask.ID,
		Task:   &updatedTask.Task,
		IsDone: &updatedTask.IsDone,
	}

	// Возвращаем обновленную задачу
	return response, nil
}

func (h *Handler) DeleteTask(_ context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	// Распаковываем ID задачи
	id := request.Id

	// Обращаемся к сервису и удаляем задачу по ID
	err := h.Service.DeleteTaskByID(id)
	if err != nil {
		return nil, err
	}

	// Возвращаем успешный ответ без содержимого
	return tasks.DeleteTasksId204Response{}, nil
}

// Добавление метода для strict интерфейса с правильным первым параметром (context.Context)
func (h *Handler) DeleteTasksId(ctx context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	// Используем request.Id напрямую
	id := request.Id

	// Обращаемся к сервису и удаляем задачу по ID
	err := h.Service.DeleteTaskByID(id)
	if err != nil {
		return nil, err
	}

	// Возвращаем успешный ответ без содержимого
	return tasks.DeleteTasksId204Response{}, nil
}

func (h *Handler) UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		http.Error(w, "missing task ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid task ID", http.StatusBadRequest)
		return
	}

	var task taskService.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedTask, err := h.Service.UpdateTaskByID(uint(id), task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(updatedTask); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (h *Handler) DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		http.Error(w, "missing task ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid task ID", http.StatusBadRequest)
		return
	}

	if err := h.Service.DeleteTaskByID(uint(id)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
