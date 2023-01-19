package firestore

import (
	"context"
	"github.com/tobywiedenhoefer/linkholder-go/pkg/firebase"
	"github.com/tobywiedenhoefer/linkholder-go/pkg/firebase/firestore/models"
	"google.golang.org/api/iterator"
	"log"
)

func GetLinks(ctx *context.Context) (*models.Links, error) {
	client, err := firebase.GetClient(ctx)
	if err != nil {
		return nil, err
	}
	iter := client.Collection("links").Documents(*ctx)
	var links models.Links
	defer iter.Stop()
	for {
		doc, err := iter.Next()
		log.Printf("Received doc from links collection: %v", doc)

		if err == iterator.Done {
			break
		} else if err != nil {
			return nil, err
		}

		var link models.Link
		if err := doc.DataTo(&link); err != nil {
			return nil, err
		}
		link.Id = doc.Ref.ID
		log.Printf("Link struct: %v \n", link)
		links.AddLink(&link)
	}
	return &links, nil
}
