package model

type Article struct {
	Id         uint     `json:"id" gorm:"primary_key" `
	CategoryId uint     `json:"category_id" gorm:"default:1"`
	Title      string   `json:"name" gorm:"type:varchar(50);not null;unique"`
	Content    string   `json:"content" gorm:"type text"`
	CreatedAt  Time     `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt  Time     `json:"updated_at" gorm:"type:timestamp"`
}
