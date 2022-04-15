package models

type MovieSynopsis struct {
	BannerWidget BannerWidget
	Meta         Meta
	Seo          Seo
}

type BannerWidget struct {
	PageCta []Page
}

type Page struct {
	Text string
}

type Meta struct {
	Event Event
}

type Event struct {
	EventName string
}

type Seo struct {
	MetaProperties []MetaProperties
}

type MetaProperties struct {
	Value string
}
