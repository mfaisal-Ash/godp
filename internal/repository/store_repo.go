package repository

import (
	domain "godp/internal/domain"

	"gorm.io/gorm"
)

type StoreRepository interface {
	Create(store *domain.Store) error
	AttachOutfit(link *domain.StoreOutfit) error
	FindNearby(longitude, latitude float64, radiusMeter, limit int, outfitID uint) ([]domain.Store, error)
}

type storeRepository struct{ db *gorm.DB }

func NewStoreRepository(db *gorm.DB) StoreRepository { return &storeRepository{db: db} }

func (r *storeRepository) Create(store *domain.Store) error {
	if err := r.db.Create(store).Error; err != nil {
		return err
	}
	return r.db.Exec(`UPDATE stores SET geom = ST_SetSRID(ST_MakePoint(longitude, latitude), 4326)::geography WHERE id = ?`, store.ID).Error
}

func (r *storeRepository) AttachOutfit(link *domain.StoreOutfit) error {
	return r.db.Where(domain.StoreOutfit{StoreID: link.StoreID, OutfitID: link.OutfitID}).FirstOrCreate(link).Error
}

func (r *storeRepository) FindNearby(longitude, latitude float64, radiusMeter, limit int, outfitID uint) ([]domain.Store, error) {
	var stores []domain.Store
	q := r.db.Table("stores").Select("DISTINCT stores.id, stores.name, stores.category, stores.brand, stores.address, stores.longitude, stores.latitude, stores.description, stores.created_at, stores.updated_at")
	if outfitID > 0 {
		q = q.Joins("JOIN store_outfits ON store_outfits.store_id = stores.id").Where("store_outfits.outfit_id = ?", outfitID)
	}
	q = q.Where(`ST_DWithin(stores.geom, ST_SetSRID(ST_MakePoint(?, ?), 4326)::geography, ?)`, longitude, latitude, radiusMeter)
	q = q.Order("id ASC")
	if limit > 0 {
		q = q.Limit(limit)
	}
	return stores, q.Scan(&stores).Error
}
