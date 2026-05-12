package usecase

import "godp/internal/domain"

type OutfitUsecase interface {
	Create(outfit *domain.Outfit) error
	FindAll(search, category, style, gender string, isMember bool) ([]domain.Outfit, error)
	FindPopularGuest() ([]domain.Outfit, error)
	FindByID(id uint) (*domain.Outfit, error)
	FindCategories() ([]string, error)
	AttachLocation(outfitID, locationID uint, score int) error
	Recommend(locationID uint, isMember bool) ([]domain.Outfit, error)
	MatchByPlace(
		longitude,
		latitude float64,
		radiusMeter int,
		category,
		style,
		gender string,
		isMember bool,
	) ([]domain.Outfit, error)
}

type outfitUsecase struct {
	outfitRepo domain.OutfitRepository
}

func NewOutfitUsecase(
	outfitRepo domain.OutfitRepository,
) OutfitUsecase {
	return &outfitUsecase{
		outfitRepo: outfitRepo,
	}
}

func (u *outfitUsecase) Create(outfit *domain.Outfit) error {
	return u.outfitRepo.Create(outfit)
}

func (u *outfitUsecase) FindAll(
	search,
	category,
	style,
	gender string,
	isMember bool,
) ([]domain.Outfit, error) {

	limit := 0

	if !isMember {
		limit = 10
	}

	return u.outfitRepo.FindAll(
		search,
		category,
		style,
		gender,
		limit,
	)
}

func (u *outfitUsecase) FindPopularGuest() ([]domain.Outfit, error) {
	return u.outfitRepo.FindPopular(10)
}

func (u *outfitUsecase) FindByID(id uint) (*domain.Outfit, error) {
	return u.outfitRepo.FindByID(id)
}

func (u *outfitUsecase) FindCategories() ([]string, error) {
	return u.outfitRepo.FindCategories()
}

func (u *outfitUsecase) AttachLocation(
	outfitID,
	locationID uint,
	score int,
) error {

	if score <= 0 {
		score = 1
	}

	return u.outfitRepo.AttachLocation(
		&domain.OutfitLocation{
			OutfitID:   outfitID,
			LocationID: locationID,
			Score:      score,
		},
	)
}

func (u *outfitUsecase) Recommend(
	locationID uint,
	isMember bool,
) ([]domain.Outfit, error) {

	limit := 0

	if !isMember {
		limit = 10
	}

	return u.outfitRepo.FindRecommendationByLocation(
		locationID,
		limit,
	)
}

func (u *outfitUsecase) MatchByPlace(
	longitude,
	latitude float64,
	radiusMeter int,
	category,
	style,
	gender string,
	isMember bool,
) ([]domain.Outfit, error) {

	limit := 0

	if !isMember {
		limit = 10

		if radiusMeter > 1000 || radiusMeter <= 0 {
			radiusMeter = 1000
		}
	}

	if isMember && radiusMeter <= 0 {
		radiusMeter = 5000
	}

	return u.outfitRepo.FindRecommendationByCoordinate(
		longitude,
		latitude,
		radiusMeter,
		limit,
		category,
		style,
		gender,
	)
}
