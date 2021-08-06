package model

type Category struct {
	Id        uint      `json:"id" gorm:"primary_key" `
	Name      string    `json:"name" gorm:"type:varchar(50);not null;unique"`
	Sort      int       `json:"sort" gorm:"type:int(11);default:1"`
	Fid       int       `json:"fid" gorm:"type:int(11);default:0"`
	CreatedAt Time      `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt Time      `json:"updated_at" gorm:"type:timestamp"`
	Articles  []Article
}
