package repositories

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"tyres.kz/internal/models"
)

type PostRepositoryInterface interface {
	Create(post *models.Post) error
	GetAll() ([]*models.Post, error)
	GetByID(id int) (*models.Post, error)
	GetByName(name string) (*models.Post, error)
	Update(post *models.Post) error
	Delete(id int) error
}

type PostRepository struct {
	db     *sql.DB
	logger *log.Logger
}

func NewPostRepository(db *sql.DB, logger *log.Logger) (*PostRepository, error) {
	if db == nil {
		return nil, fmt.Errorf("db is nil")
	}

	if logger == nil {
		return nil, fmt.Errorf("logger is nil")
	}

	return &PostRepository{db: db, logger: logger}, nil
}

func (repository *PostRepository) Create(post *models.Post) error {
	if post == nil {
		repository.logger.Println("Attempted to create nil value post")
		return nil
	}

	query := "insert into posts (name, description, price, created_at, updated_at) values (?, ?, ?, ?, ?)"
	_, err := repository.db.Exec(query, post.Title, post.Content, post.Price, post.CreatedAt, post.UpdatedAt)

	if err != nil {
		repository.logger.Printf("Error inserting post: %v", err)
		return err
	}

	return nil
}

func (repository *PostRepository) GetAll() ([]*models.Post, error) {
	query := "select * from posts"
	rows, err := repository.db.Query(query)

	if err != nil {
		repository.logger.Printf("Error getting all posts: %v", err)
		return nil, err
	}

	defer rows.Close()

	posts := make([]*models.Post, 0)
	for rows.Next() {
		post := models.Post{}
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.Price, &post.CreatedAt, &post.UpdatedAt)
		if err != nil {
			repository.logger.Printf("Error scanning post: %v", err)
			return nil, err
		}
		posts = append(posts, &post)
	}

	return posts, nil
}

func (repository *PostRepository) GetByID(id int) (*models.Post, error) {
	query := "select * from posts where id = ?"
	row := repository.db.QueryRow(query, id)

	post := models.Post{}
	err := row.Scan(&post.ID, &post.Title, &post.Content, &post.Price, &post.CreatedAt, &post.UpdatedAt)
	if err != nil {
		repository.logger.Printf("Error scanning post: %v", err)
		return nil, err
	}

	return &post, nil
}

func (repository *PostRepository) GetByName(name string) (*models.Post, error) {
	query := "select * from posts where name = ?"
	row := repository.db.QueryRow(query, name)

	post := models.Post{}
	err := row.Scan(&post.ID, &post.Title, &post.Content, &post.Price, &post.CreatedAt, &post.UpdatedAt)
	if err != nil {
		repository.logger.Printf("Error scanning post: %v", err)
		return nil, err
	}

	return &post, nil
}

func (repository *PostRepository) Update(post *models.Post) error {
	query := "update posts set name = ?, description = ?, price = ?, updated_at = ? where id = ?"
	_, err := repository.db.Exec(query, post.Title, post.Content, post.Price, post.UpdatedAt, post.ID)

	if err != nil {
		repository.logger.Printf("Error updating post: %v", err)
		return err
	}

	return nil
}

func (repository *PostRepository) Delete(id int) error {
	query := "delete from posts where id = ?"
	_, err := repository.db.Exec(query, id)

	if err != nil {
		repository.logger.Printf("Error deleting post: %v", err)
		return err
	}

	return nil
}
