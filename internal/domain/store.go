package domain

import "time"

type StoreOutfit struct {
	ID        uint      `json:"id"`
	StoreID   uint      `json:"store_id"`
	OutfitID  uint      `json:"outfit_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Store struct {
	ID            uint      `json:"id"`
	Name          string    `json:"name"`
	Category      string    `json:"category"`
	Description   string    `json:"description"`
	Longitude     float64   `json:"longitude"`
	Latitude      float64   `json:"latitude"`
	Address       string    `json:"address"`
	AverageRating float64   `json:"average_rating"`
	TotalReviews  int       `json:"total_reviews"`
	CreatedBy     uint      `json:"created_by"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type StoreRepository interface {
	Create(store *Store) error
	FindNearby(longitude float64, latitude float64, radiusMeter int) ([]Store, error)
}

type StoreUsecase interface {
	CreateStore(store *Store) error
	GetNearbyStores(longitude float64, latitude float64, radiusMeter int) ([]Store, error)
}
