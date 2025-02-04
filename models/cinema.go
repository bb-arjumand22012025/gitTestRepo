package models

type Cinema struct {
	CinemaID        int    `json:"c_id"`
	MovieID         int    `json:"m_id"`
	TheatreName     string `json:"theatre_name"`
	TheatreLocation string `json:"theatre_location"`
	SeatsBooked     int    `json:"seats_booked"`
	SeatsVacant     int    `json:"seats_available"`
	TotalSeats      int    `json:"total_seats"`
	RoomNo          int    `json:"room_no"`
}
