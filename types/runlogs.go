package types

import "time"

type RunLog struct {
	ID        uint64
	Type      string
	Url       string
	IsSuccess string
	Desc      string
	Date      time.Time
}
