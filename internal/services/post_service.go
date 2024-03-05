package services

import (
	"log"
	"tyres.kz/internal/models"
	"tyres.kz/internal/repositories"
)

type PostService struct {
	postRepository repositories.PostRepositoryInterface
	logger         *log.Logger
}

func NewPostService(postRepository repositories.PostRepositoryInterface, logger *log.Logger) *PostService {
	return &PostService{
		postRepository: postRepository,
		logger:         logger,
	}
}

func (service *PostService) CreatePost(post *models.Post) error {
	if post == nil {
		service.logger.Println("Attempted to create nil value post")
		return nil
	}

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
