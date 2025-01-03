package routes

import (
	"github.com/AyushA9rawal/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(r *mux.Router) {
	r.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	r.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	r.HandleFunc("/book/{bookid}", controllers.GetBookById).Methods("GET")
	r.HandleFunc("/book/{bookid}", controllers.UpdateBook).Methods("PUT")
	r.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")

}
