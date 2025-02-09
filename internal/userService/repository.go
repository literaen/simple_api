package userService

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserServiceRepository interface {
	GetUsers() ([]User, error)
	PostUser(user *User) error
	PatchUserByID(id uint, user *User) error
	DeleteUserByID(id uint) error
}

type userServiceRepository struct {
	db *gorm.DB
}

// Функция для хеширования пароля
func hashPassword(password string) (string, error) {
	// Генерируем хеш с использованием bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// // Функция для проверки пароля
// func checkPassword(hashedPassword, password string) bool {
// 	// Проверяем введенный пароль с хешированным паролем
// 	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
// 	return err == nil
// }

func NewUserRepository(db *gorm.DB) *userServiceRepository {
	return &userServiceRepository{db: db}
}

func (u *userServiceRepository) GetUsers() ([]User, error) {
	var users []User
	err := u.db.Find(&users).Error
	return users, err
}

func (u *userServiceRepository) PostUser(user *User) error {
	password := user.Password

	hashPassword, err := hashPassword(password)
	if err != nil {
		return err
	}
	user.Password = hashPassword

	result := u.db.Create(&user)
	return result.Error
}

func (u *userServiceRepository) PatchUserByID(id uint, user *User) error {
	password := user.Password

	hashPassword, err := hashPassword(password)
	if err != nil {
		return err
	}
	user.Password = hashPassword

	if err := u.db.Model(&User{}).Where("id = ?", id).Updates(&user); err != nil {
		return err.Error
	}
	return nil
}

func (u *userServiceRepository) DeleteUserByID(id uint) error {
	if err := u.db.Where("id = ?", id).Delete(&User{}); err != nil {
		return err.Error
	}
	return nil
}
