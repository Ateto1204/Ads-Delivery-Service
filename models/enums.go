package models

type Gender string

const (
	Male   Gender = "M"
	Female Gender = "F"
)

type Platform string

const (
	Android Platform = "android"
	IOS     Platform = "ios"
	Web     Platform = "web"
)

type Country string

const (
	TW Country = "TW"
	JP Country = "JP"
	US Country = "US"
	KR Country = "KR"
)
