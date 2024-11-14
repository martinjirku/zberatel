package page

import (
	"jirku.sk/zberatel/model"
	"jirku.sk/zberatel/template/components"
)

type CollectionsNewModel struct {
	Form        model.CollectionInput
	TypeOptions []components.Option
	Errors      map[string][]error
	CsrfToken   string
	GlobalError error
}

type CollectionsListModel struct {
	Collections model.PagingResponse[model.Collection]
	Error       error
	Paging      model.PagingRequest
}
