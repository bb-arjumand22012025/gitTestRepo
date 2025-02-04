package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"project/db"
	"project/models"
)

func GetAllCinemas(w http.ResponseWriter, r *http.Request) {
	// Query all cinemas from the cinema_hall table
	rows, err := db.DB.Query(`
		SELECT 
			c_id, 
			theatre_name, 
			seats_booked, 
			seats_available, 
			total_seats, 
			room_no 
		FROM 
			cinema_hall
	`)
	if err != nil {
		http.Error(w, "Failed to query cinemas", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Slice to store all cinema details
	var cinemas []models.Cinema
	for rows.Next() {
		var cinema models.Cinema
		if err := rows.Scan(&cinema.CinemaID, &cinema.TheatreName, &cinema.SeatsBooked, &cinema.SeatsVacant, &cinema.TotalSeats, &cinema.RoomNo); err != nil {
			http.Error(w, "Error reading database", http.StatusInternalServerError)
			return
		}
		cinemas = append(cinemas, cinema)
	}

	// Set response header and encode data to JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cinemas)
}

func AddTheatre(w http.ResponseWriter, r *http.Request) {
	var cinemas models.Cinema
	if err := json.NewDecoder(r.Body).Decode(&cinemas); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// now we want to insert values into the table - so we write an insert query
	query := `INSERT INTO cinema_hall (m_id, theatre_name, theatre_location, seats_available, seats_booked, total_seats, room_no ) VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err := db.DB.Exec(query, cinemas.MovieID, cinemas.TheatreName, cinemas.TheatreLocation, cinemas.SeatsVacant, cinemas.SeatsBooked, cinemas.TotalSeats, cinemas.RoomNo)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to add theatre: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Cinema '%s' added successfully! ", cinemas.TheatreName)))
}

func UpdateCinema(w http.ResponseWriter, r *http.Request) {
	var cinema models.Cinema
	if err := json.NewDecoder(r.Body).Decode(&cinema); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	query := `UPDATE cinema_hall SET theatre_name = ?, total_seats = ? WHERE c_id = ?`
	_, err := db.DB.Exec(query, cinema.TheatreName, cinema.TotalSeats, cinema.CinemaID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to update theatre: %v", err), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Cinema name '%s' updated successfully ", cinema.TheatreName)))
}

func DeleteTheatre(w http.ResponseWriter, r *http.Request) {
	var cinema models.Cinema
	if err := json.NewDecoder(r.Body).Decode(&cinema); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	query := `DELETE FROM cinema_hall WHERE c_id = ?`
	_, err := db.DB.Exec(query, cinema.CinemaID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to update theatre : %v", err), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Cinema id '%d' deleted successfully ", cinema.CinemaID)))
}

