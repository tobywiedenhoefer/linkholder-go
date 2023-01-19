package models

type Link struct {
	Id        string
	Href      bool   `firestore:"href"`
	URL       string `firestore:"url"`
	ShortName string `firestore:"shortName"`
	UserId    string `firestore:"userId"`
}
