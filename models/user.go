package models

type User struct {
	Id        int    `gorm:"autoIncrement;primaryKey" json:"id"`
	Name      string `gorm:"not null" json:"name"`
	CreatedAt int    `gorm:"autoCreateTime" json:"created_at"`
	UpdateAt  int    `gorm:"autoUpdateTime" json:"update_at"`
}

func (User) TableName() string {
	return "users"
}

type UsersRequest struct {
	PageNum  int `form:"page_num"`
	PageSize int `form:"page_size"`
}

type UsersResponse struct {
	Result bool   `json:"result"`
	Users  []User `json:"users"`
}

type AddUserRequest struct {
	Name string `form:"name" binding:"required" json:"name"`
}

type AddUserResponse struct {
	Result bool `json:"result"`
	User   User `json:"user"`
}
