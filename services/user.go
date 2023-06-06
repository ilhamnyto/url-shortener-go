package services

import (
	"database/sql"
	"strings"
	"time"

	"github.com/ilhamnyto/url-shortener-go/entity"
	"github.com/ilhamnyto/url-shortener-go/pkg/encryption"
	"github.com/ilhamnyto/url-shortener-go/pkg/token"
	"github.com/ilhamnyto/url-shortener-go/repositories"
)

type InterfaceUserService interface {
	CreateUser(req *entity.CreateUserRequest) *entity.CustomError
	Login(req *entity.UserLoginRequest) (*entity.UserLoginResponse, *entity.CustomError)
	UserProfile(userId int) (*entity.UserProfileResponse, *entity.CustomError)
	UserProfileByUsername(username string) (*entity.UserProfileResponse, *entity.CustomError)
	UpdateUserPassword(userId int, req *entity.UpdatePasswordRequest) *entity.CustomError
}

type UserService struct {
	repo repositories.InterfaceUserRepository
}

func NewUserService(repo repositories.InterfaceUserRepository) InterfaceUserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(req *entity.CreateUserRequest) *entity.CustomError {
	if !strings.Contains(req.Email, "@") {
		return entity.BadRequestError("Wrog email format.")
	}
	if len(req.Username) < 6 {
		return entity.BadRequestError("Username should have at least 6 characters")
	}
	if len(req.Password) < 8 {
		return entity.BadRequestError("Password should have at least 8 characters")
	}

	usernameIsExist, emailIsExsist, err := s.repo.CheckUsernameAndEmail(req.Username, req.Email)

	if err != nil {
		return entity.RepositoryError(err.Error())
	}

	if usernameIsExist {
		return entity.BadRequestError("Username already been used.")
	}

	if emailIsExsist {
		return entity.BadRequestError("Email already been used")
	}

	salt, err := encryption.GenerateSalt()

	if err != nil {
		return entity.GeneralError(err.Error())
	}

	hashedPassword, err := encryption.HashPassword(req.Password, salt)

	if err != nil {
		return entity.GeneralError(err.Error())
	}

	u := entity.User{
		Username: req.Username,
		Email: req.Email,
		Password: hashedPassword,
		Salt: salt,
		CreatedAt: time.Now(),
	}

	if err := s.repo.Create(u); err != nil {
		return entity.RepositoryError(err.Error())
	}

	return nil
}

func (s *UserService) Login(req *entity.UserLoginRequest) (*entity.UserLoginResponse, *entity.CustomError) {

	user, err := s.repo.GetUserCredential(req.Username)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, entity.NotFoundError("Username not found.")
		}

		return nil, entity.RepositoryError(err.Error())
	}
	
	if err = encryption.ValidatePassword(user.Password, req.Password, user.Salt); err != nil {
		return nil, entity.BadRequestError("Wrong password.")
	}
	
	token, err := token.GenerateToken(user.ID)
	
	if err != nil {
		return nil, entity.RepositoryError(err.Error())
	}
	
	return &entity.UserLoginResponse{Token: token}, nil
}

func (s *UserService) UserProfile(userId int) (*entity.UserProfileResponse, *entity.CustomError) {
	user, err := s.repo.GetUserById(userId)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, entity.NotFoundError("Username not found.")
		}

		return nil, entity.RepositoryError(err.Error())
	}

	return &entity.UserProfileResponse{Username: user.Username, Email: user.Email, CreatedAt: user.CreatedAt}, nil
}

func (s *UserService) UserProfileByUsername(username string) (*entity.UserProfileResponse, *entity.CustomError) {
	user, err := s.repo.GetUserByUsername(username)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, entity.NotFoundError("Username not found.")
		}

		return nil, entity.RepositoryError(err.Error())
	}

	return &entity.UserProfileResponse{Username: user.Username, Email: user.Email, CreatedAt: user.CreatedAt}, nil
}

func (s *UserService) UpdateUserPassword(userId int, req *entity.UpdatePasswordRequest) *entity.CustomError {
	if req.Password != req.ConfirmPassword {
		return entity.BadRequestError("Password didn't match.")
	}	

	if len(req.Password) < 8 || len(req.ConfirmPassword) < 8 {
		return entity.BadRequestError("Password should have at least 8 characters")
	}

	user, err := s.repo.GetUserCredentialById(userId)

	if err != nil {
		if err == sql.ErrNoRows {
			return entity.NotFoundError("User not found.")
		}

		return entity.RepositoryError(err.Error())
	}

	if err := encryption.ValidatePassword(user.Password, req.Password, user.Salt); err == nil {
		return entity.BadRequestError("Password cant be the same as before.")
	}


	hashedPassword, err := encryption.HashPassword(req.Password, user.Salt)

	if err != nil {
		return entity.GeneralError(err.Error())
	}

	if err := s.repo.UpdateUserPassword(hashedPassword, userId); err != nil {
		return entity.RepositoryError(err.Error())
	}

	return nil
}

