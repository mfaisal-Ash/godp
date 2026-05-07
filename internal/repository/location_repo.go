package repository

import (
	"godp/internal/entity"

	"gorm.io/gorm"
)

type LocationRepository interface {
	Create(location *entity.Location) error
	FindAll(limit int) ([]entity.Location, error)
	FindNearby(longitude, latitude float64, radiusMeter, limit int) ([]entity.Location, error)
}

type locationRepository struct{ db *gorm.DB }

func NewLocationRepository(db *gorm.DB) LocationRepository { return &locationRepository{db: db} }

func (r *locationRepository) Create(location *entity.Location) error {
	if err := r.db.Create(location).Error; err != nil {
		return err
	}
	return r.db.Exec(`UPDATE locations SET geom = ST_SetSRID(ST_MakePoint(longitude, latitude), 4326)::geography WHERE id = ?`, location.ID).Error
}

func (r *locationRepository) FindAll(limit int) ([]entity.Location, error) {
	var locations []entity.Location
	q := r.db.Order("created_at DESC")
	if limit > 0 {
		q = q.Limit(limit)
	}
	return locations, q.Find(&locations).Error
}

func (r *locationRepository) FindNearby(longitude, latitude float64, radiusMeter, limit int) ([]entity.Location, error) {
	var locations []entity.Location
	q := r.db.Table("locations").Select("id, name, category, address, description, longitude, latitude, created_at, updated_at").Where(`ST_DWithin(geom, ST_SetSRID(ST_MakePoint(?, ?), 4326)::geography, ?)`, longitude, latitude, radiusMeter).Order("id ASC")
	if limit > 0 {
		q = q.Limit(limit)
	}
	return locations, q.Scan(&locations).Error
}
