package routes

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/tobywiedenhoefer/linkholder-go/pkg/firebase/firestore"
)

func Links(w http.ResponseWriter, r *http.Request, ctx *context.Context) {
	switch r.Method {
	case "GET":
		links, err := firestore.GetLinks(ctx)
		if err != nil {
			w.WriteHeader(404)
			log.Fatalf("Ran into a problem getting links from firestore: %v \nStruct: %v \n", err, links)
		}
		b, err := json.Marshal(links)
		if err != nil {
			w.WriteHeader(404)
			log.Fatalf("Ran into an issue while marshalling links struct: %v \nStruct: %v \n", err, links)
		}
		if err := json.NewEncoder(w).Encode(b); err != nil {
			w.WriteHeader(404)
			log.Fatalf("Ran into an issue while encoding links struct: %v \nStruct: %v \n", err, links)
		}
	default:
		w.WriteHeader(404)
		log.Fatalln("Requested unsupported method")
	}
}
