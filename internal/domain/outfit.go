package domain

import "time"

type Outfit struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Category    string    `json:"category"`
	Style       string    `json:"style"`
	Gender      string    `json:"gender"`
	ImageURL    string    `json:"image_url"`
	Description string    `json:"description"`
	LocationID  uint      `json:"location_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type OutfitRepository interface {
	Create(outfit *Outfit) error
	FindAll() ([]Outfit, error)
	FindByID(id uint) (*Outfit, error)
	FindByCategory(category string) ([]Outfit, error)
	FindByLocation(locationID uint) ([]Outfit, error)
}

type OutfitUsecase interface {
	CreateOutfit(outfit *Outfit) error
	GetAllOutfits() ([]Outfit, error)
	GetOutfitByID(id uint) (*Outfit, error)
	GetOutfitsByCategory(category string) ([]Outfit, error)
	GetOutfitRecommendation(locationID uint) ([]Outfit, error)
}
