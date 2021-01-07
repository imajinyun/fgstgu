package hello

import (
	"encoding/json"
	"fgstgu/go/src/b/grpc/b14/gen"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// HelloServiceInterface interface.
type HelloServiceInterface interface {
	Hello(in *gen.String, out *gen.String) error
}

// HelloServiceHandler func.
func HelloServiceHandler(svc HelloServiceInterface) http.Handler {
	var router = httprouter.New()
	handleHelloServiceHelloGet(router, svc)
	return router
}

func handleHelloServiceHelloGet(router httprouter.Router, svc HelloServiceInterface) {
	router.Handle("GET", "/hello/:value", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		var request, response gen.String
		err := gen.PopulateFieldFromPath(&request, fieldPath, ps.ByName("value"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err := svc.Hello(request, response); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err := json.NewEncoder(w).Encode(&response); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}
