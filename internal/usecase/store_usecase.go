package usecase

import (
	"godp/internal/domain"
	"godp/internal/repository"
)

type StoreUsecase interface {
	Create(store *domain.Store) error
	AttachOutfit(storeID, outfitID uint) error
	FindNearby(longitude, latitude float64, radiusMeter int, outfitID uint, isMember bool) ([]domain.Store, error)
}

type storeUsecase struct{ storeRepo repository.StoreRepository }

func NewStoreUsecase(storeRepo repository.StoreRepository) StoreUsecase {
	return &storeUsecase{storeRepo: storeRepo}
}
func (u *storeUsecase) Create(store *domain.Store) error { return u.storeRepo.Create(store) }
func (u *storeUsecase) AttachOutfit(storeID, outfitID uint) error {
	return u.storeRepo.AttachOutfit(&domain.StoreOutfit{StoreID: storeID, OutfitID: outfitID})
}
func (u *storeUsecase) FindNearby(longitude, latitude float64, radiusMeter int, outfitID uint, isMember bool) ([]domain.Store, error) {
	limit := 0
	if !isMember {
		limit = 5
		if radiusMeter > 1000 || radiusMeter <= 0 {
			radiusMeter = 1000
		}
	}
	if isMember && radiusMeter <= 0 {
		radiusMeter = 5000
	}
	return u.storeRepo.FindNearby(longitude, latitude, radiusMeter, limit, outfitID)
}
