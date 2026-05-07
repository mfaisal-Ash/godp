package usecase

import (
	"godp/internal/entity"
	"godp/internal/repository"
)

type LocationUsecase interface {
	Create(location *entity.Location) error
	FindAll(isMember bool) ([]entity.Location, error)
	FindNearby(longitude, latitude float64, radiusMeter int, isMember bool) ([]entity.Location, error)
}

type locationUsecase struct{ locationRepo repository.LocationRepository }

func NewLocationUsecase(locationRepo repository.LocationRepository) LocationUsecase {
	return &locationUsecase{locationRepo: locationRepo}
}

func (u *locationUsecase) Create(location *entity.Location) error {
	return u.locationRepo.Create(location)
}

func (u *locationUsecase) FindAll(isMember bool) ([]entity.Location, error) {
	limit := 0
	if !isMember {
		limit = 10
	}
	return u.locationRepo.FindAll(limit)
}

func (u *locationUsecase) FindNearby(longitude, latitude float64, radiusMeter int, isMember bool) ([]entity.Location, error) {
	limit := 0
	if !isMember {
		limit = 10
		if radiusMeter > 1000 || radiusMeter <= 0 {
			radiusMeter = 1000
		}
	}
	if isMember && radiusMeter <= 0 {
		radiusMeter = 5000
	}
	return u.locationRepo.FindNearby(longitude, latitude, radiusMeter, limit)
}
