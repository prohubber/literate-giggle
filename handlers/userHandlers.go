package handlers

import (
	"context"

	"main/project/internal/userService"
	"main/project/internal/web/users"
)

type UserHandler struct {
	Service *userService.UserService
}

func NewUserHandler(service *userService.UserService) *UserHandler {
	return &UserHandler{Service: service}
}

func (h *UserHandler) GetUsers(_ context.Context, _ users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	allUsers, err := h.Service.GetAllUsers()
	if err != nil {
		return nil, err
	}

	response := users.GetUsers200JSONResponse{}
	for _, usr := range allUsers {
		user := users.User{
			Id:        &usr.ID,
			Email:     &usr.Email,
			CreatedAt: &usr.CreatedAt,
			UpdatedAt: &usr.UpdatedAt,
		}
		response = append(response, user)
	}
	return response, nil
}

func (h *UserHandler) PostUsers(_ context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	userRequest := request.Body
	userToCreate := userService.User{
		Email:    *userRequest.Email,
		Password: *userRequest.Password,
	}
	// Метод CreateUser возвращает только ошибку, поэтому убираем второе значение
	err := h.Service.CreateUser(&userToCreate)
	if err != nil {
		return nil, err
	}

	// В ответе используем поля из уже созданного объекта
	response := users.PostUsers201JSONResponse{
		Id:        &userToCreate.ID,
		Email:     &userToCreate.Email,
		CreatedAt: &userToCreate.CreatedAt,
		UpdatedAt: &userToCreate.UpdatedAt,
	}
	return response, nil
}

func (h *UserHandler) PatchUsersId(ctx context.Context, request users.PatchUsersIdRequestObject) (users.PatchUsersIdResponseObject, error) {
	userRequest := request.Body
	userToUpdate := userService.User{
		ID:       uint(request.Id), // Добавляем ID для обновления
		Email:    *userRequest.Email,
		Password: *userRequest.Password,
	}
	err := h.Service.UpdateUser(&userToUpdate)
	if err != nil {
		return nil, err
	}

	response := users.PatchUsersId200JSONResponse{
		Id:        &userToUpdate.ID,
		Email:     &userToUpdate.Email,
		CreatedAt: &userToUpdate.CreatedAt,
		UpdatedAt: &userToUpdate.UpdatedAt,
	}
	return response, nil
}

func (h *UserHandler) DeleteUser(_ context.Context, request users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {
	id := request.Id
	err := h.Service.DeleteUserByID(uint(id))
	if err != nil {
		return nil, err
	}
	return users.DeleteUsersId204Response{}, nil
}

func (h *UserHandler) DeleteUsersId(ctx context.Context, request users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {
	id := request.Id
	err := h.Service.DeleteUserByID(uint(id))
	if err != nil {
		return nil, err
	}
	return users.DeleteUsersId204Response{}, nil
}
