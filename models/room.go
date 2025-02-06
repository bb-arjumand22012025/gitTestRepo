package models

type Room struct {
	RoomId      int    `json:"r_id"`
	RoomNo      int    `json:"room_no"`
	TheatreID   int    `json:"c_id"`
	MovieName   string `json:"movie_name"`
	TotalSeats  int    `json:"total_seats"`
	BookedSeats int    `json:"booked_seats"`
	MovieID     int    `json:"movie_id`
}

// 		r_id int
// 		room_no int,
//     	c_id int,
//     	movie_name varchar(20),
//     	total_seats int,
//     	booked_seats int,
