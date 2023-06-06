package repositories

import (
	"database/sql"

	"github.com/ilhamnyto/url-shortener-go/entity"
)

var (
	queryCreateUser = `INSERT INTO users (username, email, password, salt, created_at) VALUES ($1, $2, $3, $4, $5)`

	queryCheckUsername = `SELECT count(username) from users WHERE username = $1`

	queryCheckEmail = `SELECT count(email) from users WHERE email = $1`

	queryGetUserCredential = `SELECT id, password, salt from users WHERE username = $1`

	queryGetUserById = `SELECT username, email, date_trunc('second', created_at) as created_at from users WHERE id = $1`

	queryGetUserByUsername = `SELECT username, email, date_trunc('second', created_at) as created_at from users WHERE username = $1`

	queryUpdateUserPassword = `UPDATE users SET password = $1 where id = $2`

	queryGetUserCredentialById = `SELECT id, password, salt from users WHERE id = $1`
)

type InterfaceUserRepository interface {
	Create(user entity.User) (error)
	CheckUsernameAndEmail(username string, email string) (bool, bool, error)
	GetUserCredential(username string) (*entity.User, error)
	GetUserCredentialById(userId int) (*entity.User, error)
	GetUserById(userId int) (*entity.User, error)
	GetUserByUsername(username string) (*entity.User, error)
	UpdateUserPassword(password string, userId int) error
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) InterfaceUserRepository {
	return &UserRepository{db: db}
} 

func (r *UserRepository) Create(user entity.User) (error) {
	stmt, err := r.db.Prepare(queryCreateUser)

	if err != nil {
		return err
	}

	if _, err := stmt.Exec(user.Username, user.Email, user.Password, user.Salt, user.CreatedAt); err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) CheckUsernameAndEmail(username string, email string) (bool, bool, error) {
	stmt, err := r.db.Prepare(queryCheckUsername)

	if err != nil {
		return false, false, err
	}

	row := stmt.QueryRow(username)

	var usernameCount int

	if err = row.Scan(&usernameCount); err != nil {
		return false, false, err
	}

	stmt, err = r.db.Prepare(queryCheckEmail)

	if err != nil {
		return false, false, err
	}

	row = stmt.QueryRow(email)

	var emailCount int

	if err = row.Scan(&emailCount); err != nil {
		return false, false, err
	}
	
	return usernameCount > 0, emailCount > 0, nil
}

func (r *UserRepository) GetUserCredential(username string) (*entity.User, error) {
	stmt, err := r.db.Prepare(queryGetUserCredential)

	if err != nil {
		return nil, err
	}

	row := stmt.QueryRow(username)

	var user entity.User

	if err = row.Scan(&user.ID, &user.Password, &user.Salt); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetUserCredentialById(userId int) (*entity.User, error) {
	stmt, err := r.db.Prepare(queryGetUserCredentialById)

	if err != nil {
		return nil, err
	}

	row := stmt.QueryRow(userId)

	var user entity.User

	if err = row.Scan(&user.ID, &user.Password, &user.Salt); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetUserById(userId int) (*entity.User, error) {
	stmt, err := r.db.Prepare(queryGetUserById)

	if err != nil {
		return nil, err
	}

	row := stmt.QueryRow(userId)

	var user entity.User

	if err = row.Scan(&user.Username, &user.Email, &user.CreatedAt); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetUserByUsername(username string) (*entity.User, error) {
	stmt, err := r.db.Prepare(queryGetUserByUsername)

	if err != nil {
		return nil, err
	}

	row := stmt.QueryRow(username)

	var user entity.User

	if err = row.Scan(&user.Username, &user.Email, &user.CreatedAt); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) UpdateUserPassword(password string, userId int) error {
	stmt, err := r.db.Prepare(queryUpdateUserPassword)

	if err != nil {
		return err
	}

	if _, err := stmt.Exec(password, userId); err != nil {
		return err
	}

	return nil
}