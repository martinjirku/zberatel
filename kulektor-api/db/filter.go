package db

type FilterType string

const (
	FilterTypeEQ FilterType = "eq"
)

type Valuer interface {
	Value() any
}

type Filter struct {
	Column string
	Type   FilterType
	Value  any
}
