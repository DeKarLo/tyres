package services

import (
	"fmt"
	"log"
	"time"

	"tyres.kz/internal/models"
	"tyres.kz/internal/repositories"
)

type PostServiceInterface interface {
	CreatePost(post *models.Post) error
	GetAllPosts() ([]*models.Post, error)
	GetPostByID(id int) (*models.Post, error)
	GetPostByName(name string) (*models.Post, error)
	UpdatePost(post *models.Post) error
	DeletePost(id int) error
}

type PostService struct {
	postRepository repositories.PostRepositoryInterface
	logger         *log.Logger
}

func NewPostService(postRepository repositories.PostRepositoryInterface, logger *log.Logger) (PostServiceInterface, error) {
	if postRepository == nil {
		return nil, fmt.Errorf("userRepository is nil")
	}

	if logger == nil {
		return nil, fmt.Errorf("logger is nil")
	}

	return &PostService{postRepository: postRepository, logger: logger}, nil
}

func (service *PostService) CreatePost(post *models.Post) error {
	if post == nil {
		service.logger.Println("Attempted to create nil value post")
		return nil
	}

	post.CreatedAt = time.Now()
	post.UpdatedAt = time.Now()

	err := service.postRepository.Create(post)
	if err != nil {
		service.logger.Printf("Error creating post: %v", err)
		return err
	}

	return nil
}

func (service *PostService) GetAllPosts() ([]*models.Post, error) {
	posts, err := service.postRepository.GetAll()
	if err != nil {
		service.logger.Printf("Error getting all posts: %v", err)
		return nil, err
	}

	return posts, nil
}

func (service *PostService) GetPostByID(id int) (*models.Post, error) {
	post, err := service.postRepository.GetByID(id)
	if err != nil {
		service.logger.Printf("Error getting post by ID: %v", err)
		return nil, err
	}

	return post, nil
}

func (service *PostService) GetPostByName(name string) (*models.Post, error) {
	post, err := service.postRepository.GetByName(name)
	if err != nil {
		service.logger.Printf("Error getting post by name: %v", err)
		return nil, err
	}

	return post, nil
}

func (service *PostService) UpdatePost(post *models.Post) error {
	post.UpdatedAt = time.Now()
	err := service.postRepository.Update(post)
	if err != nil {
		service.logger.Printf("Error updating post: %v", err)
		return err
	}

	return nil
}

func (service *PostService) DeletePost(id int) error {
	err := service.postRepository.Delete(id)
	if err != nil {
		service.logger.Printf("Error deleting post: %v", err)
		return err
	}

	return nil
}
