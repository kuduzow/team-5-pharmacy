package services

import (
	"errors"

	"github.com/kuduzow/team-5-pharmacy/internal/models"
	"github.com/kuduzow/team-5-pharmacy/internal/repository"
	"gorm.io/gorm"
)

var( 
	ErrUserNotFound = errors.New("Пользователь не найден")
)

type UserService interface {
	CreateUser(req models.CreateUserRequest) (*models.User, error)

	GetUserByID(id uint) (*models.User, error)

	UpdateUser(id uint, req models.UpdateUserRequest) (*models.User, error)

	//DeleteUser(id uint) error
}

type userService struct {
	users repository.UserRepository
}



func NewUserService(users repository.UserRepository) UserService {
	return &userService{users : users}
}

func(s *userService) validateUserCreate(req *models.CreateUserRequest) error {
	if req.FullName == ""{
		return errors.New("имя не может быть пустым")
	}

	if req.Phone == ""{
		return errors.New("телефон не может быть пустым")
	}

	if req.Email == ""{
		return errors.New("емаил не может пустым")
	}
	if req.DefaultAddress == ""{
		return errors.New("адрес доставки не может быть пустым")
	}
	return nil
}
func (s *userService) CreateUser(req models.CreateUserRequest) (*models.User, error) {
	if err := s.validateUserCreate(&req); err != nil {
		return nil, err
	}

	user := models.User{
		FullName:       req.FullName,
		Email:          req.Email,
		Phone:          req.Phone,
		DefaultAddress: req.DefaultAddress,
	}

	if err := s.users.Create(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *userService) GetUserByID(id uint) (*models.User, error) {
	user, err := s.users.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return user, nil
}

func (s *userService) UpdateUser(id uint, req models.UpdateUserRequest) (*models.User, error) {

	user,err := s.users.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	s.applyUserUpdate(req)

	if err := s.users.Update(user); err != nil {
		return nil, err
	}
	return user, nil
}

func(s *userService) applyUserUpdate(req models.UpdateUserRequest) error{
	if *req.FullName == ""{
		return errors.New("имя не может быть пустым")
	}

	if *req.Phone == ""{
		return errors.New("телефон не может быть пустым")
	}

	if *req.Email == ""{
		return errors.New("емаил не может пустым")
	}
	if *req.DefaultAddress == ""{
		return errors.New("адрес доставки не может быть пустым")
	}
	return nil
}


func (s *userService) DeleteUser(id uint) error {

	if _,err := s.users.GetByID(id); err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrUserNotFound
		} 
		return err
	}
	return s.users.Delete(id)
}