package service

import (
	"github.com/CUBS-sources-code/CUBS-coin/repository"
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
		return nil, err
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
		return nil, err
	}

	userResponse := UserResponse{
		StudentId: user.ID,
		Name: 	user.Name,
		Balance: user.Balance,
		CreatedAt: user.CreatedAt.String(),
	}

	return &userResponse, nil
}