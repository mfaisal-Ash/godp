package domain

import "time"

type Favorite struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	OutfitID  uint      `json:"outfit_id"`
	CreatedAt time.Time `json:"created_at"`
}

type FavoriteRepository interface {
	Create(favorite *Favorite) error
	FindByUserID(userID uint) ([]Favorite, error)
	Delete(userID uint, outfitID uint) error
}

type FavoriteUsecase interface {
	AddFavorite(userID uint, outfitID uint) error
	GetUserFavorites(userID uint) ([]Favorite, error)
	RemoveFavorite(userID uint, outfitID uint) error
}
