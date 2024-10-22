package service

type UserResponse struct {
	StudentId string `json:"student_id"`
	Name      string `json:"name"`
	Balance   int    `json:"balance"`
	CreatedAt  string `json:"created_at"`
}

type NewUserRequest struct {
	StudentId string `json:"student_id"`
	Name      string `json:"name"`
	Password  string `json:"password"`
}

type UserService interface {
	GetUsers() ([]UserResponse, error)
	GetUser(string) (*UserResponse, error)
	CreateUser(NewUserRequest) (*UserResponse, error)
	ChangeRoleToAdmin(string) (*UserResponse, error)
	ChangeRoleToMember(string) (*UserResponse, error)
}
