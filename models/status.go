package models

type Status struct {
	ScreeningID int    `json:"screening_id"`
	MovieName   string `json:"movie_name"`
	TheatreName string `json:"theatre_name"`
	RoomNo      int    `json:"room_no"`
	SeatsBooked int    `json:"seats_booked"`
	SeatsVacant int    `json:"seats_vacant"`
	TotalSeats  int    `json:"total_seats"`
}
