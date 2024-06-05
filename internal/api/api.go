package api

import (
	"fmt"
	"net/http"
	"simple-server/internal/models"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func Init() {
	fmt.Println("Init API")
	addr := ":8081"

	r := initRouter()

	http.ListenAndServe(addr, r)
}

func initRouter() *mux.Router {
	r := mux.NewRouter()
	r.Use(handlers.CORS(handlers.AllowedOrigins([]string{"http://localhost:8081/posts"})))

	r.HandleFunc("/", handlerHello).Methods("GET")

	r.HandleFunc("/posts", handlerPosts).Methods("GET")
	r.HandleFunc("/posts/{id}", handlerPost).Methods("GET")
	r.HandleFunc("/posts", handlerPostCreate).Methods("POST")
	r.HandleFunc("/posts/{id}", handlerPostUpdate).Methods("PATCH")
	r.HandleFunc("/posts/{id}", handlerPostDelete).Methods("DELETE")

	return r
}

func handlerHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello server"))
}

//////// Posts

func handlerPosts(w http.ResponseWriter, r *http.Request) {
	articles := []models.Article{
		{
			ID:          01,
			Title:       "test ti",
			Body:        "some lorem ipsum text",
			Description: "a litle desc",
			Author:      "dev",
		},
		{
			ID:          02,
			Title:       "test ti1",
			Body:        "some1 lorem ipsum text",
			Description: "a1 litle desc",
			Author:      "dev",
		},
	}
	writeJSON(w, http.StatusOK, articles)
}

func handlerPost(w http.ResponseWriter, r *http.Request) {
	article := models.Article{
		ID:          01,
		Title:       "test ti",
		Body:        "some lorem ipsum text",
		Description: "a litle desc",
		Author:      "dev",
	}
	writeJSON(w, http.StatusOK, article)
}

func handlerPostCreate(w http.ResponseWriter, r *http.Request) {
	article := models.Article{
		ID:          01,
		Title:       "test ti",
		Body:        "some lorem ipsum text",
		Description: "a litle desc",
		Author:      "dev",
	}
	writeJSON(w, http.StatusCreated, article)
}

func handlerPostUpdate(w http.ResponseWriter, r *http.Request) {
	article := models.Article{
		ID:          01,
		Title:       "test ti",
		Body:        "some lorem ipsum text",
		Description: "a litle desc",
		Author:      "dev",
	}
	writeJSON(w, http.StatusOK, article)
}

func handlerPostDelete(w http.ResponseWriter, r *http.Request) {
	res := models.Mess{
		Message: "post deleted",
	}
	writeJSON(w, http.StatusOK, res)
}
