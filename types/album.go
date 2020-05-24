package types

import "time"

type Album struct {
	ID       uint32
	Name     string
	Model    uint32
	Url      string
	Date     string
	CoverPic string
	Pics     []*Pic
}

func (a *Album) GetUpload() time.Time {
	return time.Time{}
}

func (a *Album) GetTitlePageID() int64 {
	return 0
}
