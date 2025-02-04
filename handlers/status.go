package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"project/db"
)

type BookingRequest struct {
	ScreeningID int `json:"screening_id"`
	SeatsToBook int `json:"seats_to_book"`
}

// BookSeats handles seat booking
func BookSeats(w http.ResponseWriter, r *http.Request) {
	var bookingRequest BookingRequest
	if err := json.NewDecoder(r.Body).Decode(&bookingRequest); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if bookingRequest.SeatsToBook <= 0 {
		http.Error(w, "Seats to book must be greater than 0", http.StatusBadRequest)
		return
	}

	tx, err := db.DB.Begin()
	if err != nil {
		http.Error(w, "Failed to start transaction", http.StatusInternalServerError)
		return
	}

	var seatsVacant, seatsBooked int
	query := `SELECT seats_vacant, seats_booked FROM status WHERE screening_id = ?`
	err = tx.QueryRow(query, bookingRequest.ScreeningID).Scan(&seatsVacant, &seatsBooked)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Screening not found", http.StatusNotFound)
		return
	}

	if seatsVacant < bookingRequest.SeatsToBook {
		tx.Rollback()
		http.Error(w, "Not enough vacant seats", http.StatusConflict)
		return
	}

	updateQuery := `
		UPDATE status 
		SET seats_booked = seats_booked + ?, 
		seats_vacant = seats_vacant - ? 
		WHERE screening_id = ?`
	_, err = tx.Exec(updateQuery, bookingRequest.SeatsToBook, bookingRequest.SeatsToBook, bookingRequest.ScreeningID)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Failed to update seats", http.StatusInternalServerError)
		return
	}

	if err := tx.Commit(); err != nil {
		http.Error(w, "Failed to commit transaction", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Successfully booked %d seats for screening ID %d", bookingRequest.SeatsToBook, bookingRequest.ScreeningID)))
}
