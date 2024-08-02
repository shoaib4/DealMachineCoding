package services

import (
	"DealSystem/models"
	"time"
)

type IDealService interface {
	Create(cd *models.Deal) error
	End(cd *models.Deal) error
	Update(cd *models.Deal, endTime time.Time, updatedCount int) error
	Claim(currUser *models.User, name string, currentTime time.Time) error
}
