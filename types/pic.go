package types

type Pic struct {
	ID    uint32
	Album uint32
	Name  string
	Url   string
}

func (p *Pic)Save() error {
	return nil
}
