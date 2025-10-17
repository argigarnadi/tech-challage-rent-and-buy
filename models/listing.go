package models

type Listing struct {
	Id          int    `gorm:"autoIncrement;primaryKey" json:"id"`
	UserId      int    `gorm:"not null" json:"user_id"`
	Price       int    `gorm:"not null" json:"price"`
	ListingType string `gorm:"not null" json:"listing_type"`
	CreatedAt   int    `gorm:"autoCreateTime" json:"created_at"`
	UpdateAt    int    `gorm:"autoUpdateTime" json:"update_at"`
}

func (Listing) TableName() string {
	return "listing"
}

type ListingRequest struct {
	PageNum  int `form:"page_num"`
	PageSize int `form:"page_size"`
	UserId   int `form:"user_id"`
}

type ListingResponse struct {
	Result   bool      `json:"result"`
	Listings []Listing `json:"listings"`
}

type AddListingRequest struct {
	UserId      int    `form:"user_id" binding:"required" json:"user_id"`
	Price       int    `form:"price" binding:"required" json:"price"`
	ListingType string `form:"listing_type" binding:"required" json:"listing_type"`
}

type AddListingResponse struct {
	Result  bool    `json:"result"`
	Listing Listing `json:"listing"`
}
