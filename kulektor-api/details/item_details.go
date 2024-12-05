package details

type MetaType string

const (
	MetaTypeUnkown MetaType = "unknown"
	MetaTypeString MetaType = "string"
	MetaTypeNumber MetaType = "number"
	MetaTypeDate   MetaType = "date"
)

var AllMetaType = [...]MetaType{
	MetaTypeUnkown,
	MetaTypeString,
	MetaTypeNumber,
	MetaTypeDate,
}

func (m MetaType) IsValid() bool {
	for _, t := range AllMetaType {
		if m == t {
			return true
		}
	}
	return false
}

type MetaField struct {
	Name string   `json:"name"`
	Type MetaType `json:"type"`
}

type Meta struct {
	Fields []MetaField
}

type Details map[string]any
