package service

import (
	"fmt"

	"github.com/CUBS-sources-code/CUBS-coin/errs"
	"github.com/CUBS-sources-code/CUBS-coin/logs"
	"github.com/CUBS-sources-code/CUBS-coin/repository"
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
			fmt.Println("err")
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