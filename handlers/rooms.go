package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"project/db"
	"project/models"

	"github.com/gorilla/mux"
)

// here I want to get the theatre and room no where a movie is being displayed - should have the option for letting us choose which theatre we want to book and the room we want to book.
// so for that movie we should get the theatre name, room no and seats info using the movie id/name

func GetMovieTheatre(w http.ResponseWriter, r *http.Request) { //return based on movie name, the theatre where we can see the movie, the room no and the seats info.
	movieName := mux.Vars(r)["name"]

	query := `
	SELECT 
		movie_name,
		c_id,
		room_no,
		total_seats,
		booked_seats
	FROM
		room
	WHERE movie_name=?`
	row := db.DB.QueryRow(query, movieName)

	var room models.Room
	err := row.Scan(&room.MovieName, &room.TheatreID, &room.RoomNo, &room.TotalSeats, &room.BookedSeats)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to fetch with movie with name %s, error : %v", movieName, err), http.StatusNotFound)
		return
	}

	w.Header().Set("Content", "application/json")
	json.NewEncoder(w).Encode(room)
}

// just return the whole room table
func GetRoomTable(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query(`
	SELECT 
		movie_name,
		booked_seats,
		r_id,
		room_no,
		c_id,
		total_seats
	FROM 
		room`)

	if err != nil {
		http.Error(w, "Failed to query rooms", http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	var rooms []models.Room
	for rows.Next() {
		var room models.Room
		if err := rows.Scan(&room.MovieName, &room.BookedSeats, &room.RoomId, &room.RoomNo, &room.TheatreID, &room.TotalSeats); err != nil {
			http.Error(w, "Error fetching the database", http.StatusInternalServerError)
			return
		}
		rooms = append(rooms, room)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(rooms)
	}
}

type BookRooms struct {
	MovID       int `json:"m_ID`
	RoomNo      int `json:"room_no"`
	SeatsBooked int `json:"seats_book"`
}

func BookRoom(w http.ResponseWriter, r *http.Request) {
	var bookroom BookRooms
	if err := json.NewDecoder(r.Body).Decode(&bookroom); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if bookroom.SeatsBooked <= 0 {
		http.Error(w, "Invalid input, please put positive quantity to book a seat", http.StatusBadRequest)
		return
	}

	tx, err := db.DB.Begin()
	if err != nil {
		http.Error(w, "Failed to start trasanction", http.StatusInternalServerError)
		return
	}

	// check whether we have enough seats to book.
	var seats_booked, seats_vacant int
	query := "SELECT (total_seats - booked_seats) AS vacant_seats seats_booked FROM room WHERE m_id = ?"
	err = tx.QueryRow(query, bookroom.MovID).Scan(&seats_vacant, &seats_booked)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Movie not found! ", http.StatusNotFound)
		return
	}

	if seats_vacant < bookroom.SeatsBooked {
		tx.Rollback()
		http.Error(w, "Not enough vacant seats", http.StatusConflict)
		return
	}

	updateQuery := `
	UPDATE room
	SET seats_booked = seats_booked + ?
	WHERE m_id = ?`
	_, err = tx.Exec(updateQuery, bookroom.SeatsBooked, bookroom.MovID)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Failed to update seats", http.StatusInternalServerError)
	}

	if err := tx.Commit(); err != nil {
		http.Error(w, "Failed to commit transaction", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Successfully booked %d seats for movie name %d", bookroom.SeatsBooked, bookroom.MovID)))
	// json.NewEncoder(w).Encode("Seats booked %d for movie %s", bookroom.SeatsBooked, bookroom.MovieName)
}
