package models

type PublicApiResponse struct {
	Result   bool       `json:"result"`
	Listings []Listings `json:"listings"`
}

type Listings struct {
	Id          int    `json:"id"`
	ListingType string `json:"listing_type"`
	Price       int    `json:"price"`
	CreatedAt   int    `json:"created_at"`
	UpdateAt    int    `json:"update_at"`
	User        User   `json:"user"`
}
