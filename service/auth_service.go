package service

import (
	"time"

	"github.com/CUBS-sources-code/CUBS-coin/errs"
	"github.com/CUBS-sources-code/CUBS-coin/logs"
	"github.com/CUBS-sources-code/CUBS-coin/repository"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type authService struct {
	userRepository repository.UserRepository
}

func NewAuthService(userRepository repository.UserRepository) authService {
	return authService{userRepository: userRepository}
}

func (s authService) SignUp(signUpRequest SignUpRequest) (*TokenResponse, error) {
	id := signUpRequest.StudentId
	name := signUpRequest.Name
	password := signUpRequest.Password

	if id == "" || name == "" || password == "" {
		return nil, errs.NewBadRequestError("invalid signup credentials")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(signUpRequest.Password), bcrypt.DefaultCost)
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

	role := user.Role

	token, exp, err := createJWTToken(id, role)
	if err != nil {
		return nil, errs.NewUnexpectedError()
	}

	tokenResponse := TokenResponse{
		Token: token,
		Exp:   exp,
		User: id,
	}

	return &tokenResponse, nil
}

func (s authService) SignIn(signInRequest SignInRequest) (*TokenResponse, error) {
	id := signInRequest.StudentId
	password := signInRequest.Password

	if id == "" || password == "" {
		return nil, errs.NewBadRequestError("invalid signip credentials")
	}

	user, err := s.userRepository.GetById(id)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewNotFoundError("user doesn't existed")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errs.NewBadRequestError("invalid password")
	}

	role := user.Role

	token, exp, err := createJWTToken(id, role)
	if err != nil {
		return nil, errs.NewUnexpectedError()
	}

	tokenResponse := TokenResponse{
		Token: token,
		Exp:   exp,
		User: id,
	}

	return &tokenResponse, nil
}

func createJWTToken(id string, role string) (string, int64, error) {
	exp := time.Now().Add(time.Minute * 30).Unix()
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user"] = id
	claims["exp"] = exp
	claims["role"] = role
	t, err := token.SignedString([]byte(viper.GetString("app.jwt-secret")))

	if err != nil {
		return "", 0, err
	}
	return t, exp, nil
}