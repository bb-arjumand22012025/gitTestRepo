package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"project/db"
	"project/models"

	"github.com/gorilla/mux"
)

func GetMovies(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT m_id, movie_name, movie_des FROM movie")

	if err != nil {
		log.Println("Database query failed:", err)
		http.Error(w, "Failed to query movies", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var movies []models.Movie // a slice of type modes.Movies which is basically a struct.
	for rows.Next() {         //next() is a method in the sql.rows that allows us to iterate through the rows returned by the SQL quesry, it gives us true if there is a row else it returns false.
		var movie models.Movie //to hold the data of the row
		if err := rows.Scan(&movie.ID, &movie.Name, &movie.Description); err != nil {
			http.Error(w, "Error reading database", http.StatusInternalServerError)
			return
		}
		movies = append(movies, movie) // adding the data from the variable to the slice
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func GetByMovieId(w http.ResponseWriter, r *http.Request) {
	// first we want to extract the movie ID from the URL path
	movieID := mux.Vars(r)["ID"]
	// movieName := mux.Vars(r)["name"]
	fmt.Println(movieID)
	// now we want to query the database for the movie - sql query
	query := `SELECT m_id, movie_name, movie_des FROM movie WHERE m_id = ?`
	row := db.DB.QueryRow(query, movieID)

	// then we store this result in a struct
	var movie models.Movie
	err := row.Scan(&movie.ID, &movie.Name, &movie.Description)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to fetch movie with ID %s, error: %v", movieID, err), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movie)
}
