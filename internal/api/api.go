package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"simple-server/internal/database"
	"simple-server/internal/models"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var fakeStorage *database.FakeStorage

func Init() {
	fmt.Println("Init API")
	addr := ":8081"
	fakeStorage = database.New()
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
	articles, err := fakeStorage.Get()
	if err != nil {
		writeJSON(w, http.StatusNotFound, err)
		return
	}
	writeJSON(w, http.StatusOK, articles)
}

func handlerPost(w http.ResponseWriter, r *http.Request) {
	id, err := strToInt(mux.Vars(r)["id"])
	if err != nil {
		writeJSON(w, http.StatusBadRequest, err)
		return
	}
	article, err := fakeStorage.GetBy(id)
	if err != nil {
		writeJSON(w, http.StatusNotFound, err)
		return
	}
	writeJSON(w, http.StatusOK, article)
}

func handlerPostCreate(w http.ResponseWriter, r *http.Request) {
	article := models.ArticleDTO{}
	err := json.NewDecoder(r.Body).Decode(&article)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, err)
		return
	}
	err = fakeStorage.Create(article)
	if err != nil {
		writeJSON(w, http.StatusForbidden, err)
		return
	}
	writeJSON(w, http.StatusCreated, article)
}

func handlerPostUpdate(w http.ResponseWriter, r *http.Request) {
	id, err := strToInt(mux.Vars(r)["id"])
	if err != nil {
		writeJSON(w, http.StatusBadRequest, err)
		return
	}
	article := models.Article{}
	err = json.NewDecoder(r.Body).Decode(&article)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, err)
		return
	}
	article.ID = id
	article, err = fakeStorage.Update(article)
	if err != nil {
		writeJSON(w, http.StatusForbidden, err)
		return
	}
	writeJSON(w, http.StatusOK, article)
}

func handlerPostDelete(w http.ResponseWriter, r *http.Request) {
	id, err := strToInt(mux.Vars(r)["id"])
	if err != nil {
		writeJSON(w, http.StatusBadRequest, err)
		return
	}
	err = fakeStorage.Delete(id)
	if err != nil {
		writeJSON(w, http.StatusNotFound, err)
		return
	}
	res := models.Mess{
		Message: "post deleted",
	}
	writeJSON(w, http.StatusOK, res)
}
