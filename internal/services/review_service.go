package services

import (
	"errors"

	"github.com/kuduzow/team-5-pharmacy/internal/models"
	"github.com/kuduzow/team-5-pharmacy/internal/repository"
)

type ReviewService interface {
	Create(input models.ReviewCreateInput) (*models.Review, error)
	GetByID(id uint) (*models.Review, error)
	Update(id uint, input models.ReviewUpdateInput) (*models.Review, error)
	Delete(id uint) error
}

type reviewService struct {
	repo repository.ReviewRepository
}

func NewReviewService(repo repository.ReviewRepository) ReviewService {
	return &reviewService{repo: repo}
}

func (s *reviewService) Create(input models.ReviewCreateInput) (*models.Review, error) {
	if input.Rating < 1 || input.Rating > 5 {
		return nil, errors.New("rating must be 1-5")
	}
	if input.Text == "" {
		return nil, errors.New("text is required")
	}

	review := &models.Review{
		UserID:     input.UserID,
		MedicineID: input.MedicineID,
		Rating:     input.Rating,
		Text:       input.Text,
	}

	if err := s.repo.Create(review); err != nil {
		return nil, err
	}

	return review, nil
}

func (s *reviewService) GetByID(id uint) (*models.Review, error) {
	review, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if review == nil {
		return nil, errors.New("review not found")
	}
	return review, nil
}

func (s *reviewService) Update(id uint, input models.ReviewUpdateInput) (*models.Review, error) {
	review, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if review == nil {
		return nil, errors.New("review not found")
	}

	if input.Rating != nil {
		if *input.Rating < 1 || *input.Rating > 5 {
			return nil, errors.New("rating must be 1-5")
		}
		review.Rating = *input.Rating
	}

	if input.Text != nil {
		if *input.Text == "" {
			return nil, errors.New("text cannot be empty")
		}
		review.Text = *input.Text
	}

	if err := s.repo.Update(review); err != nil {
		return nil, err
	}

	return review, nil
}

func (s *reviewService) Delete(id uint) error {
	return s.repo.Delete(id)
}
