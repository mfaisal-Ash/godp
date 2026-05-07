package domain

import "time"

type Location struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Category    string    `json:"category"`
	Description string    `json:"description"`
	Longitude   float64   `json:"longitude"`
	Latitude    float64   `json:"latitude"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type LocationRepository interface {
	Create(location *Location) error
	FindAll() ([]Location, error)
	FindByID(id uint) (*Location, error)
	FindNearby(longitude float64, latitude float64, radiusMeter int) ([]Location, error)
}

type LocationUsecase interface {
	CreateLocation(location *Location) error
	GetAllLocations() ([]Location, error)
	GetLocationByID(id uint) (*Location, error)
	GetNearbyLocations(longitude float64, latitude float64, radiusMeter int) ([]Location, error)
}
