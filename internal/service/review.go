package service

import (
	"errors"
	"pizzaria/internal/models"
)

func ValidatePizzaReview(review *models.Review) error {
	if review.Rating < 0 || review.Rating > 5 {
		return errors.New("rating must be between 0 and 5")
	}
	return nil
}
