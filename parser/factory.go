package parser

import "errors"

const (
	ALBUM = "album"
	MODEL = "model"
)
func FactoryNew(id uint32, html string, env string) Parser  {
	switch env {
	case ALBUM:
		return NewAlbumParser(id, html)
	case MODEL:
		return NewModelParser(id, html)
	default:
		panic(errors.New("un support parser"))
	}
}
