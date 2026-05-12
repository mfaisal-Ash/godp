package repository

import (
	"godp/internal/domain"

	"gorm.io/gorm"
)

type outfitRepository struct {
	db *gorm.DB
}

// FindByLocation implements [domain.OutfitRepository].
func (r *outfitRepository) FindByLocation(locationID uint) ([]domain.Outfit, error) {
	panic("unimplemented")
}

func (r *outfitRepository) Create(outfit *domain.Outfit) error {
	return r.db.Create(outfit).Error
}

func (r *outfitRepository) AttachLocation(
	outfitLocation *domain.OutfitLocation,
) error {

	return r.db.Create(outfitLocation).Error
}

func (r *outfitRepository) FindAll(
	search,
	category,
	style,
	gender string,
	limit int,
) ([]domain.Outfit, error) {

	var outfits []domain.Outfit

	query := r.db.Model(&domain.Outfit{})

	if search != "" {
		query = query.Where(
			"name ILIKE ?",
			"%"+search+"%",
		)
	}

	if category != "" {
		query = query.Where(
			"category = ?",
			category,
		)
	}

	if style != "" {
		query = query.Where(
			"style = ?",
			style,
		)
	}

	if gender != "" {
		query = query.Where(
			"gender = ?",
			gender,
		)
	}

	if limit > 0 {
		query = query.Limit(limit)
	}

	err := query.Find(&outfits).Error

	return outfits, err
}

func (r *outfitRepository) FindByID(id uint) (*domain.Outfit, error) {
	var outfit domain.Outfit

	err := r.db.First(&outfit, id).Error
	if err != nil {
		return nil, err
	}

	return &outfit, nil
}

func (r *outfitRepository) FindRecommendationByLocation(
	locationID uint,
	limit int,
) ([]domain.Outfit, error) {

	var outfits []domain.Outfit

	query := r.db.
		Table("outfits").
		Joins(`
			JOIN outfit_locations
			ON outfit_locations.outfit_id = outfits.id
		`).
		Where("outfit_locations.location_id = ?", locationID).
		Order("outfit_locations.score desc")

	if limit > 0 {
		query = query.Limit(limit)
	}

	err := query.Find(&outfits).Error

	return outfits, err
}

func (r *outfitRepository) FindCategories() ([]string, error) {
	var categories []string

	err := r.db.
		Model(&domain.Outfit{}).
		Distinct().
		Pluck("category", &categories).Error

	return categories, err
}

func (r *outfitRepository) FindPopular(limit int) ([]domain.Outfit, error) {
	var outfits []domain.Outfit

	query := r.db.Order("created_at desc")

	if limit > 0 {
		query = query.Limit(limit)
	}

	err := query.Find(&outfits).Error

	return outfits, err
}
