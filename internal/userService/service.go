package userService

import "simple_api/internal/taskService"

type UserService struct {
	repo UserServiceRepository
}

func NewUserService(repo UserServiceRepository) *UserService {
	return &UserService{repo: repo}
}

func (u *UserService) GetUsers() ([]User, error) {
	return u.repo.GetUsers()
}

func (u *UserService) GetTasksForUser(id uint) ([]taskService.Task, error) {
	return u.repo.GetTasksForUser(id)
}

// func (u *UserService) PostUser(user *User) (*User, error) {
func (u *UserService) PostUser(user *User) error {
	return u.repo.PostUser(user)
}

func (u *UserService) PatchUserByID(id uint, user *User) error {
	return u.repo.PatchUserByID(id, user)
}

func (u *UserService) DeleteUserByID(id uint) error {
	return u.repo.DeleteUserByID(id)
}
