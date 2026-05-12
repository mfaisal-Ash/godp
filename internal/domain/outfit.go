package domain

import "time"

type Outfit struct {
	ID            uint      `json:"id"`
	Name          string    `json:"name"`
	Category      string    `json:"category"`
	Style         string    `json:"style"`
	Gender        string    `json:"gender"`
	UploadBy      uint      `json:"upload_by"`
	ImageURL      string    `json:"image_url"`
	Description   string    `json:"description"`
	LocationID    uint      `json:"location_id"`
	ViewCount     int       `json:"view_count"`
	FavoriteCount int       `json:"favorite_count"`
	ReviewCount   int       `json:"review_count"`
	TrendingScore float64   `json:"trending_score"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type OutfitLocation struct {
	ID         uint `json:"id"`
	OutfitID   uint `json:"outfit_id"`
	LocationID uint `json:"location_id"`
	Score      int  `json:"score"`
}

type OutfitRepository interface {
	Create(outfit *Outfit) error

	FindAll(
		search,
		category,
		style,
		gender string,
		limit int,
	) ([]Outfit, error)

	FindPopular(limit int) ([]Outfit, error)

	FindByID(id uint) (*Outfit, error)

	FindCategories() ([]string, error)

	AttachLocation(
		outfitLocation *OutfitLocation,
	) error

	FindRecommendationByLocation(
		locationID uint,
		limit int,
	) ([]Outfit, error)

	FindRecommendationByCoordinate(
		longitude,
		latitude float64,
		radiusMeter,
		limit int,
		category,
		style,
		gender string,
	) ([]Outfit, error)
}

type OutfitUsecase interface {
	CreateOutfit(outfit *Outfit) error
	GetAllOutfits() ([]Outfit, error)
	GetOutfitByID(id uint) (*Outfit, error)
	GetOutfitsByCategory(category string) ([]Outfit, error)
	GetOutfitRecommendation(locationID uint) ([]Outfit, error)
}
