package repository_subscribers

import "patreon/internal/app/models"

type Repository interface {
	// Create Errors:
	//		app.GeneralError with Errors
	//			repository.DefaultErrDB
	Create(subscriber *models.Subscriber) error
	// Delete Errors:
	//		app.GeneralError with Errors
	//			repository.DefaultErrDB
	Delete(subscriber *models.Subscriber) error
	// GetCreators Errors:
	//		app.GeneralError with Errors
	//			repository.DefaultErrDB
	GetCreators(userID int64) ([]int64, error)
	// GetSubscribers Errors:
	//		app.GeneralError with Errors
	//			repository.DefaultErrDB
	GetSubscribers(creatorID int64) ([]int64, error)
	// Get Errors:
	//		app.GeneralError with Errors
	//			repository.DefaultErrDB
	Get(userID int64, creatorID int64) (bool, error)
}