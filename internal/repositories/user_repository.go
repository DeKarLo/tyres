package repositories

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"tyres.kz/internal/models"
)

type UserRepositoryInterface interface {
	Create(user *models.User) (int, error)
	GetByID(id int) (*models.User, error)
	GetByUsername(username string) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	GetByPhone(phone string) (*models.User, error)
	Update(user *models.User) error
	Delete(id int) error
}

type UserRepository struct {
	db *sql.DB
	logger *log.Logger
}

func NewUserRepository(db *sql.DB, logger *log.Logger) (*UserRepository, error) {
	if db == nil {
		return nil, fmt.Errorf("db is nil")
	}

	if logger == nil {
		return nil, fmt.Errorf("logger is nil")
	}

	return &UserRepository{db: db, logger: logger}, nil
}

func (repository *UserRepository) Create(user *models.User) (int, error) {
	if user == nil {
		repository.logger.Println("Attempted to create nil value user")
		return 0, fmt.Errorf("user is nil")
	}

	query := "insert into users (username, email, phone, hashed_password, is_admin, created_at, updated_at) values (?, ?, ?, ?, ?, ?, ?)"
	result, err := repository.db.Exec(query, user.Username, user.Email, user.Phone, user.HashedPassword, user.IsAdmin, user.CreatedAt.String(), user.UpdatedAt.String())

	if err != nil {
		repository.logger.Printf("Error inserting user: %v", err)
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		repository.logger.Printf("Error getting last insert id: %v", err)
		return 0, err
	}

	return int(id), nil
}

func (repository *UserRepository) GetByID(id int) (*models.User, error) {
	query := "select id, username, email, phone, hashed_password, is_admin, created_at, updated_at from users where id = ?"
	row := repository.db.QueryRow(query, id)

	user := &models.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Phone, &user.HashedPassword, &user.IsAdmin, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		repository.logger.Printf("Error scanning user: %v", err)
		return nil, err
	}

	return user, nil
}

func (repository *UserRepository) GetByUsername(username string) (*models.User, error) {
	query := "select id, username, email, phone, hashed_password, is_admin, created_at, updated_at from users where username = ?"
	row := repository.db.QueryRow(query, username)

	user := &models.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Phone, &user.HashedPassword, &user.IsAdmin, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		repository.logger.Printf("Error scanning user: %v", err)
		return nil, err
	}

	return user, nil
}

func (repository *UserRepository) GetByEmail(email string) (*models.User, error) {
	query := "select id, username, email, phone, hashed_password, is_admin, created_at, updated_at from users where email = ?"
	row := repository.db.QueryRow(query, email)

	user := &models.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Phone, &user.HashedPassword, &user.IsAdmin, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		repository.logger.Printf("Error scanning user: %v", err)
		return nil, err
	}

	return user, nil
}

func (repository *UserRepository) GetByPhone(phone string) (*models.User, error) {
	query := "select id, username, email, phone, hashed_password, is_admin, created_at, updated_at from users where phone = ?"
	row := repository.db.QueryRow(query, phone)

	user := &models.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Phone, &user.HashedPassword, &user.IsAdmin, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		repository.logger.Printf("Error scanning user: %v", err)
		return nil, err
	}

	return user, nil
}

func (repository *UserRepository) Update(user *models.User) error {
	if user == nil {
		repository.logger.Println("Attempted to update nil value user")
		return fmt.Errorf("user is nil")
	}

	query := "update users set username = ?, email = ?, phone = ?, hashed_password = ?, is_admin = ?, updated_at = ? where id = ?"
	_, err := repository.db.Exec(query, user.Username, user.Email, user.Phone, user.HashedPassword, user.IsAdmin, user.UpdatedAt.String(), user.ID)

	if err != nil {
		repository.logger.Printf("Error updating user: %v", err)
		return err
	}

	return nil
}

func (repository *UserRepository) Delete(id int) error {
	query := "delete from users where id = ?"
	_, err := repository.db.Exec(query, id)

	if err != nil {
		repository.logger.Printf("Error deleting user: %v", err)
		return err
	}

	return nil
}
