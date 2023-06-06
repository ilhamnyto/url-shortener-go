package entity

import "time"

type User struct {
	ID        int    	`json:"id"`
	Username  string 	`json:"username"`
	Email     string 	`json:"email"`
	Password  string 	`json:"password"`
	Salt      string 	`json:"salt"`
	CreatedAt time.Time	`json:"created_at"`
	UpdatedAt time.Time	`json:"updated_at"`
}

type CreateUserRequest struct {
	Username	string 	`json:"username"`
	Email		string	`json:"email"`
	Password	string	`json:"password"`
}

type UserLoginRequest struct {
	Username 	string 	`json:"username"`
	Password 	string 	`json:"password"`
}

type UpdatePasswordRequest struct {
	Password			string		`json:"password"`
	ConfirmPassword 	string 		`json:"confirm_password"`
	UpdatedAt 			time.Time	`json:"updated_at"`
}

type UserProfileResponse struct {
	Username 	string 		`json:"username"`
	Email		string 		`json:"email"`
	CreatedAt	time.Time	`json:"created_at"`
}

type UserLoginResponse struct {
	Token	string  `json:"token"`
}
