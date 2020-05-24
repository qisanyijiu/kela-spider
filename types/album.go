package types

import "time"

type Album struct {
	ID        uint32
	Name      string
	Model     uint32
	Url       string
	Desc      string
	Date      string
	TitlePage string
	Pics      []*Pic
}

func (a *Album) GetUpload() time.Time {

}

func (a *Album) GetTitlePageID() int64 {

}
