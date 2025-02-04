package main

import (
	"log"
	"net/http"

	"project/db"
	"project/handlers"
	"project/middleware"

	"github.com/gorilla/mux"
)

func main() {
	db.Init()

	r := mux.NewRouter()

	// Routes with authorization
	r.Handle("/movies", middleware.BasicAuthMiddleware(http.HandlerFunc(handlers.GetMovies))).Methods("GET")
	r.Handle("/cinema", middleware.BasicAuthMiddleware(http.HandlerFunc(handlers.GetAllCinemas))).Methods("GET")
	r.Handle("/seat", middleware.BasicAuthMiddleware(http.HandlerFunc(handlers.BookSeats))).Methods("POST")

	r.Handle("/addcinema", middleware.BasicAuthMiddleware(http.HandlerFunc(handlers.AddTheatre))).Methods("POST")
	r.Handle("/updatecinema", middleware.BasicAuthMiddleware(http.HandlerFunc((handlers.UpdateCinema)))).Methods("PUT")
	r.Handle("/deletecinema", middleware.BasicAuthMiddleware(http.HandlerFunc((handlers.DeleteTheatre)))).Methods("DELETE")
	r.Handle("/getmovie/{ID}", middleware.BasicAuthMiddleware(http.HandlerFunc((handlers.GetByMovieId)))).Methods("GET")


	// routes without authorization
	// r.HandleFunc("/movies", handlers.GetMovies).Methods(("GET"))
	// r.HandleFunc("/cinema", handlers.GetAllCinemas).Methods(("GET"))
	// r.HandleFunc("/seat", handlers.BookSeats).Methods(("POST"))

	log.Println("Server is running on :8084")
	log.Fatal(http.ListenAndServe(":8084", r))
}
