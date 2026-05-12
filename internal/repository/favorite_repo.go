package repository

import (
	"godp/internal/domain"

	"gorm.io/gorm"
)

type FavoriteRepository interface {
	Create(favorite *domain.Favorite) error
	FindByUserID(userID uint) ([]domain.Favorite, error)
	Delete(userID, outfitID uint) error
}

type favoriteRepository struct{ db *gorm.DB }

func NewFavoriteRepository(db *gorm.DB) FavoriteRepository { return &favoriteRepository{db: db} }

func (r *favoriteRepository) Create(favorite *domain.Favorite) error {
	return r.db.Where(domain.Favorite{UserID: favorite.UserID, OutfitID: favorite.OutfitID}).FirstOrCreate(favorite).Error
}

func (r *favoriteRepository) FindByUserID(userID uint) ([]domain.Favorite, error) {
	var favorites []domain.Favorite
	err := r.db.Preload("Outfit").Where("user_id = ?", userID).Order("created_at DESC").Find(&favorites).Error
	return favorites, err
}

func (r *favoriteRepository) Delete(userID, outfitID uint) error {
	return r.db.Where("user_id = ? AND outfit_id = ?", userID, outfitID).Delete(&domain.Favorite{}).Error
}
