package grid

type Paging struct {
	Limit  *int `json:"limit" gqlgen:"limit"`
	Offset *int `json:"offset" gqlgen:"offset"`
}

func NewPaging(limit *int, offset *int) Paging {
	p := Paging{Limit: limit, Offset: offset}
	if p.Limit == nil {
		l := 20
		p.Limit = &l
	}
	if p.Offset == nil {
		o := 0
		p.Offset = &o
	}
	return p
}

func (p *Paging) GetLimit() int {
	if p.Limit == nil {
		return 20
	}
	return *p.Limit
}
func (p *Paging) GetOffset() int {
	if p.Offset == nil {
		return 0
	}
	return *p.Offset
}

func (p Paging) NextPage(total int64) *Paging {
	out := NewPaging(p.Limit, p.Offset)
	nextOffset := (p.GetLimit() + p.GetOffset())
	if int(total) <= nextOffset {
		return nil
	}
	out.Offset = &nextOffset
	return &out
}

func (p Paging) PrevPage() *Paging {
	out := NewPaging(p.Limit, p.Offset)
	nextOffset := p.GetOffset() - p.GetLimit()
	if nextOffset < 0 {
		return nil
	}
	out.Offset = &nextOffset
	return &out
}
func (p Paging) CurrentPage() *Paging {
	out := NewPaging(p.Limit, p.Offset)
	return &out
}
