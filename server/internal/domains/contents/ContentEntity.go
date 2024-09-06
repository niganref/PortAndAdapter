package domain

type Content struct {
	title  string
	content string
}

type ContentsList struct {
	contents []Content
	maxPage int32
}
