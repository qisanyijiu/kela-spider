package types

import "time"

type Model struct {
	ID          uint32    `json:"id" gorm:"column:id"`
	Name        string    `json:"name" gorm:"column:name"`
	Height      uint8     `json:"height" gorm:"column:height"`
	Weight      uint8     `json:"weight" gorm:"column:weight"`
	Cup         string    `json:"cup" gorm:"column:cup"`
	Bust        uint8     `json:"bust" gorm:"column:bust"`
	Waist       uint8     `json:"waist" gorm:"column:waist"`
	Hips        uint8     `json:"hips" gorm:"column:hips"`
	Fans        uint32    `json:"fans" gorm:"column:fans"`
	City        string    `json:"city" gorm:"column:city"`
	AlbumNums   uint32    `json:"album_nums" gorm:"column:album_nums"`
	VideoNums   int32     `json:"video_nums" gorm:"column:video_nums"`
	Date        time.Time `json:"date" gorm:"column:date"`
	Avatar      string    `json:"avatar" gorm:"column:avatar"`
	CoverImg    string    `json:"cover_img" gorm:"column:cover_img"`
	BannerImg   string    `json:"banner_img" gorm:"column:banner_img"`
	Job         string    `json:"job" gorm:"column:job"`
	Description string    `json:"description" gorm:"column:description"`
	Keyword     string    `json:"keyword" gorm:"column:keyword"`
}
