package usecase

import (
	"godp/internal/domain"
)

type FavoriteUsecase interface {
	Add(userID, outfitID uint) error
	FindByUser(userID uint) ([]domain.Favorite, error)
	Remove(userID, outfitID uint) error
}

type favoriteUsecase struct{ favoriteRepo domain.FavoriteRepository }

func NewFavoriteUsecase(favoriteRepo domain.FavoriteRepository) FavoriteUsecase {
	return &favoriteUsecase{favoriteRepo: favoriteRepo}
}

func (u *favoriteUsecase) Add(userID, outfitID uint) error {
	return u.favoriteRepo.Create(&domain.Favorite{UserID: userID, OutfitID: outfitID})
}
func (u *favoriteUsecase) FindByUser(userID uint) ([]domain.Favorite, error) {
	return u.favoriteRepo.FindByUserID(userID)
}
func (u *favoriteUsecase) Remove(userID, outfitID uint) error {
	return u.favoriteRepo.Delete(userID, outfitID)
}
