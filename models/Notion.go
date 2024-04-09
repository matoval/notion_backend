package models

type External struct {
	Url string
}

type File struct {
	Url string
	ExpiryTime string
}

type Cover struct {
	Type string
	External External
	File File 
}

type CreatedBy struct {
	Id string
	Object string
}

type Icon struct {
	Type string
	Url string
	ExpiryTime string
	Emoji string
}

type LastEditedBy struct {
	Id string
	Object string
}

type Parent struct {
	DatabaseId string
	Type string
}

type NotionData struct {
	Archived bool
	Cover File
	CreatedBy CreatedBy
	CreatedTime string
	Icon Icon
	Id string
	InTrash bool
	LastEditedBy LastEditedBy
	LastEditedTime string
	Object string
	Parent Parent
	Properties any
	PublicUrl string
	Url string
}
