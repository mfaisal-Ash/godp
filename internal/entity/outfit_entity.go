package entity

type Outfit struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Category    string `json:"category"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Style       string `json:"style"`

	IsPopuler bool `json:"is_popular"`
	IsViral   bool `json:"is_viral"`
}
