package service

import (
	"github.com/CUBS-sources-code/CUBS-coin/errs"
	"github.com/CUBS-sources-code/CUBS-coin/logs"
	"github.com/CUBS-sources-code/CUBS-coin/repository"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) userService {
	return userService{userRepository: userRepository}
}

func (s userService) GetUsers() ([]UserResponse, error) {

	users, err := s.userRepository.GetAll()
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	userResponses := []UserResponse{}
	for _, user := range users {
		userResponse := UserResponse{
			StudentId: user.ID,
			Name: 	user.Name,
			Balance: user.Balance,
			CreatedAt: user.CreatedAt.String(),
		}
		userResponses = append(userResponses, userResponse)
	}

	return userResponses, nil
}

func (s userService) GetUser(id string) (*UserResponse, error) {
	user, err := s.userRepository.GetById(id)
	if err != nil {

		if err == gorm.ErrRecordNotFound {
			logs.Error(err)
			return nil, errs.NewNotFoundError("user not found")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	userResponse := UserResponse{
		StudentId: user.ID,
		Name: 	user.Name,
		Balance: user.Balance,
		CreatedAt: user.CreatedAt.String(),
	}

	return &userResponse, nil
}

func (s userService) CreateUser(userRequest NewUserRequest) (*UserResponse, error) {
	id := userRequest.StudentId
	name := userRequest.Name
	password := userRequest.Password

	if id == "" || name == "" || password == "" {
		return nil, errs.NewBadRequestError("invalid signup credentials")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errs.NewUnexpectedError()
	}
	password = string(hash)

	user, err := s.userRepository.Create(id, name, password)
	if err != nil {

		if err == gorm.ErrDuplicatedKey {
			logs.Error(err)
			return nil, errs.NewBadRequestError("user already exists")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	userResponse := UserResponse{
		StudentId: user.ID,
		Name: 	user.Name,
		Balance: user.Balance,
		CreatedAt: user.CreatedAt.String(),
	}

	return &userResponse, nil
}

func (s userService) ChangeRoleToAdmin(id string) (*UserResponse, error) {
	user, err := s.userRepository.ChangeRoleToAdmin(id)
	if err != nil {

		if err == gorm.ErrRecordNotFound {
			logs.Error(err)
			return nil, errs.NewNotFoundError("user not found")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	userResponse := UserResponse{
		StudentId: user.ID,
		Name: 	user.Name,
		Balance: user.Balance,
		CreatedAt: user.CreatedAt.String(),
	}

	return &userResponse, nil
}

func (s userService) ChangeRoleToMember(id string) (*UserResponse, error) {
	user, err := s.userRepository.ChangeRoleToMember(id)
	if err != nil {

		if err == gorm.ErrRecordNotFound {
			logs.Error(err)
			return nil, errs.NewNotFoundError("user not found")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	userResponse := UserResponse{
		StudentId: user.ID,
		Name: 	user.Name,
		Balance: user.Balance,
		CreatedAt: user.CreatedAt.String(),
	}

	return &userResponse, nil
}