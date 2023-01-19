package models

type Links struct {
	Collection []*Link `json:"collection"`
}

func (links *Links) AddLink(link *Link) {
	links.Collection = append(links.Collection, link)
}
