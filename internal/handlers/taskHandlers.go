package handlers

import (
	"context"
	"errors"
	"simple_api/internal/taskService"
	"simple_api/internal/web/tasks"
)

type Handler struct {
	Service *taskService.TaskService
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
			IsDone: tsk.IsDone,
			UserId: &tsk.UserID,
		}
		response = append(response, task)
	}

	// САМОЕ ПРЕКРАСНОЕ. Возвращаем просто респонс и nil!
	return response, nil
}

// func (h *Handler) GetTasksId(_ context.Context, request tasks.GetTasksIdRequestObject) (tasks.GetTasksIdResponseObject, error) {
// 	// Получение всех задач из сервиса
// 	userID := request.Id

// 	allTasks, err := h.Service.GetTasksByUserID(userID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	response := tasks.GetTasksId200JSONResponse{}

// 	for _, tsk := range allTasks {
// 		task := tasks.Task{
// 			Id:     &tsk.ID,
// 			Task:   &tsk.Task,
// 			IsDone: &tsk.IsDone,
// 			UserId: &tsk.UserID,
// 		}
// 		response = append(response, task)
// 	}

// 	return response, nil
// }

func (h *Handler) PostTasks(_ context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	// Распаковываем тело запроса напрямую, без декодера!
	taskRequest := request.Body

	if taskRequest.UserId == nil {
		return nil, errors.New("trying to create task without user id")
	}

	if taskRequest.IsDone == nil {
		defaultIsDone := false
		taskRequest.IsDone = &defaultIsDone
	}

	// Обращаемся к сервису и создаем задачу
	taskToCreate := taskService.Task{
		Task:   *taskRequest.Task,
		IsDone: taskRequest.IsDone,
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
		IsDone: createdTask.IsDone,
		UserId: &createdTask.UserID,
	}
	// Просто возвращаем респонс!
	return response, nil
}

// PatchTasksId implements tasks.StrictServerInterface.
func (h *Handler) PatchTasksId(ctx context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	taskID := request.Id
	taskRequest := request.Body

	var updatedTask taskService.Task
	if taskRequest.Task != nil {
		updatedTask.Task = *taskRequest.Task
	}
	if taskRequest.IsDone != nil {
		updatedTask.IsDone = taskRequest.IsDone
	}
	if taskRequest.UserId != nil {
		updatedTask.UserID = *taskRequest.UserId
	}

	updatedTask, err := h.Service.UpdateTaskByID(taskID, updatedTask)
	if err != nil {
		return nil, err
	}

	// создаем структуру респонс
	response := tasks.PatchTasksId200JSONResponse{
		Id:     &updatedTask.ID,
		Task:   &updatedTask.Task,
		IsDone: updatedTask.IsDone,
		UserId: &updatedTask.UserID,
	}
	// Просто возвращаем респонс!
	return response, nil
}

// DeleteTasksId implements tasks.StrictServerInterface.
func (h *Handler) DeleteTasksId(ctx context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	taskID := request.Id

	err := h.Service.DeleteTaskByID(taskID)
	if err != nil {
		return nil, err
	}

	return tasks.DeleteTasksId204Response{}, nil
}

func NewTasksHandler(service *taskService.TaskService) *Handler {
	return &Handler{
		Service: service,
	}
}
