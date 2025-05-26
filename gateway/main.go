package main

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/paulnasdaq/fms-v2/common"
	"log"
	"net/http"
)

type ItemResource interface {
	GET(r *http.Request) (any, int, error)
	DELETE(r *http.Request) (any, int, error)
	PUT(r *http.Request) (any, int error)
}
type ListResource interface {
	ItemResource
	POST(r *http.Request) (any, int, error)
}

type API struct {
	router chi.Router
}

func NewApiServer() *API {
	r := chi.NewRouter()
	return &API{r}
}

func (a *API) addResource(path string, resource ItemResource) {
	a.router.Get(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Could not do stuff", http.StatusUnprocessableEntity)
			return
		}
		res, status, _ := resource.GET(r)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		if err := json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		}
	})
}
func (a *API) run(port int) {
	log.Fatal(http.ListenAndServe(":8080", a.router))
}

type TodoResource struct {
}

type User struct {
	FirstName string
	LastName  string
}

func (r2 *TodoResource) GET(r *http.Request) (any, int) {
	return &User{
		FirstName: "hehe",
		LastName:  "Hehe",
	}, 201
}

func main() {
	server := Server{common.NewClientManager()}
	server.RegisterRoutes()
	log.Fatal(http.ListenAndServe(":5000", nil))
}
