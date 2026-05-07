package entity

type Favorite struct {
	ID       int `json:"id"`
	UserID   int `json:"user_id"`
	OutfitID int `json:"outfit_id"`
}
