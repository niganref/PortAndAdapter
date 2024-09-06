package domain

type ContentPort interface {
	GetContent()(content Content,error)
	GetContents(page int32)(contents ContentsList,error)
}
