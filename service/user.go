package service

type UserResponse struct {
	StudentId string `json:"student_id"`
	Name      string `json:"name"`
	Balance   int    `json:"balance"`
	CreatedAt  string `json:"created_at"`
}

type UserService interface {
	GetUsers() ([]UserResponse, error)
	GetUser(string) (*UserResponse, error)
}
