package handlers

import (
	"context"
	"simple_api/internal/userService"
	"simple_api/internal/web/users"
)

type UsersHandler struct {
	Service *userService.UserService
}

// GetUsers implements users.StrictServerInterface.
func (u *UsersHandler) GetUsers(ctx context.Context, request users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	allUsers, err := u.Service.GetUsers()
	if err != nil {
		return nil, err
	}

	response := make(users.GetUsers200JSONResponse, 0, len(allUsers))
	for _, user := range allUsers {
		response = append(response, users.User{
			Id:       &user.ID,
			Email:    &user.Email,
			Password: &user.Password,
		})
	}

	return response, nil
}

// PostUsers implements users.StrictServerInterface.
func (u *UsersHandler) PostUsers(ctx context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	body := request.Body

	user := userService.User{
		Email:    *body.Email,
		Password: *body.Password,
	}

	if err := u.Service.PostUser(&user); err != nil {
		return nil, err
	}

	response := users.PostUsers201JSONResponse{
		Id:       &user.ID,
		Email:    &user.Email,
		Password: &user.Password,
	}

	return response, nil
}

// PatchUsersId implements users.StrictServerInterface.
func (u *UsersHandler) PatchUsersId(ctx context.Context, request users.PatchUsersIdRequestObject) (users.PatchUsersIdResponseObject, error) {
	body := request.Body
	userID := request.Id

	user := userService.User{
		Email:    *body.Email,
		Password: *body.Password,
	}

	if err := u.Service.PatchUserByID(userID, &user); err != nil {
		return nil, err
	}

	response := users.PatchUsersId200JSONResponse{
		Id:       &userID,
		Email:    &user.Email,
		Password: &user.Password,
	}

	return response, nil
}

// DeleteUsersId implements users.StrictServerInterface.
func (u *UsersHandler) DeleteUsersId(ctx context.Context, request users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {
	userID := request.Id
	if err := u.Service.DeleteUserByID(userID); err != nil {
		return nil, err
	}

	response := users.DeleteUsersId204Response{}
	return response, nil
}

func NewUsersHandler(service *userService.UserService) *UsersHandler {
	return &UsersHandler{Service: service}
}
