CREATE DATABASE movie_cinema;
USE movie_cinema;

CREATE TABLE movie (
	m_id int primary key,
    movie_name varchar(20),
    movie_des varchar(50)
    );

INSERT INTO movie VALUES 
	(1, 'Veer Zara', 'Achi hai bro dekh le'),
    (2, 'Tere Naam', 'bhot bhadia'),
    (3, 'Wake Up Sid', 'peak ranbir'),
    (4, 'kahona pyaar hai', 'tunu tunu walaa dance krna hai');



SELECT * FROM movie;


CREATE TABLE cinema_hall (
	m_id int,
    c_id int PRIMARY KEY,
    theatre_name varchar(20),
    theatre_location varchar(20),
    seats_available int,
    seats_booked int,
    total_seats int,
    room_no int,
    foreign key (m_id) REFERENCES movie (m_id)
);


-- DROP TABLE cinema_hall;
ALTER TABLE cinema_hall ADD c_id int;
-- INSERT INTO cinema_hall (c_id) VALUES 
-- 	(1),     (2),    (3);


INSERT INTO cinema_hall VALUES 
	(1,1,'theatre 1', 'location 1',10,30,50, 2),
    (1,2, 'theatre 2', 'location 2', 20, 10, 30, 3),
    (3,3, 'theatre 3', 'location 3', 15, 15, 30, 5);


INSERT INTO cinema_hall VALUES
	(2,4,'theatre 1', 'location 2', 20, 25, 45, 2),
    (4,6, 'theatre 4', 'location 5', 30, 40, 70, 5),
    (4, 5, 'theatre 6', 'location 6', 30, 20, 50, 10);


SELECT * FROM cinema_hall;

CREATE TABLE status ( -- this will be the transactional table that will be used to update the details of booked seats vs vacant seats
    screening_id int PRIMARY KEY,
    m_id int,
	c_id int,
    room_no int,
    seats_booked int,
    seats_vacant int,
    total_seats int,
    foreign key (m_id) REFERENCES movie (m_id),
    foreign key (c_id) REFERENCES cinema_hall (c_id)
);

INSERT INTO status VALUES (1, 2,2,3,10,20,30);
INSERT INTO status VALUES 
	(2, 1,1,2,20,10,30),
    (3,2,2,3, 30, 20, 50),
    (4, 3, 4, 4,10, 10, 20);
	
SELECT 
	ms.screening_id,
    m.movie_name,
    ch.theatre_name,
    ms.room_no,
    ms.seats_booked,
    ms.seats_vacant,
    ms.total_seats
FROM
	status ms
JOIN
	movie m ON ms.m_id = m.m_id
JOIN
	cinema_hall ch ON ms.c_id = ch.c_id;

UPDATE status 
SET seats_booked = seats_booked + 1, seats_vacant = seats_vacant- 1
WHERE screening_id = 1;

UPDATE status SET seats_vacant = 0 WHERE screening_id = 1;
SELECT * FROM status;

CREATE table credentials (
	username varchar(50),
    pass varchar(50),
    initial_login varchar(50),
    last_login varchar(50)
    );

INSERT INTO credentials Values
	("user1","pass1","10-01-2025 7:20:13", "15-01-2025 12:13:19"),
    ("user2","pass2","14-02-2025 17:06:04", "14-03-2025 10:10:28"),
    ("user2","pass2","16-07-2025 23:21:43", "17-07-2025 08:22:15");
    
SELECT * FROM credentials;

ALTER TABLE cinema_hall DROP COLUMN created_on;

ALTER TABLE cinema_hall ADD (
      created_by VARCHAR(60),
      created_on TIMESTAMP
);
-- ALTER TABLE cinema_hall
-- MODIFY COLUMN c_id INT AUTO_INCREMENT;

ALTER TABLE status
DROP FOREIGN KEY status_ibfk_2;

ALTER TABLE cinema_hall
MODIFY COLUMN c_id INT AUTO_INCREMENT;

ALTER TABLE status
ADD CONSTRAINT fk_status_cinema_hall
FOREIGN KEY (c_id) REFERENCES cinema_hall(c_id);

-- Check if c_id is AUTO_INCREMENT
SHOW CREATE TABLE cinema_hall;

-- Check if the foreign key is recreated
SHOW CREATE TABLE status;

SELECT * FROM cinema_hall;

UPDATE cinema_hall SET created_on = CURRENT_TIMESTAMP WHERE c_id;

SET SQL_SAFE_UPDATES = 1;


SELECT * FROM movie WHERE m_id = 2;



-- Creating room table
CREATE TABLE room (
	r_id int PRIMARY KEY AUTO_INCREMENT, -- this will be different ffrom the room_no bcz this help to identify rows/cols in the table, whereas room no can be same across diff theatres.
    room_no int,
    c_id int,
    movie_name varchar(20),
    total_seats int,
    booked_seats int,
    
    FOREIGN KEY (c_id) REFERENCES cinema_hall (c_id),
    FOREIGN KEY (movie_name) REFERENCES movie (movie_name)
); 
-- make the refrenced col either unique or primary key in the parent table bcz only then it can be used as foreign key
Alter table movie change movie_name movie_name varchar(20) unique; -- to be abel to use it a FK, making it unique
Alter table room ALTER booked_seats set default '0';
Alter table room ADD m_id int;
ALTER table room ADD foreign key (m_id) REFERENCES movie(m_id)  ;
UPDATE room SET m_id = 4 WHERE r_id=5;
SELECT * FROM room;

INSERT INTO room VALUES
	(1,1,1,"veer zara",50,0),
    (2,4,1,"Tere Naam",100,0),
	(3,6,2,"Wake Up Sid", 150,0),
    (4,3,3,"Veer zara", 200,0),
    (5,3,2,"Kahona pyaar hai", 80,0);

SELECT * FROM room;
desc room;


-- now we want to modify our cinemas table, we want to the only keep total rooms in a cinema and remove any other seats/room related info
ALTER TABLE cinema_hall 
	DROP COLUMN seats_booked, 
    DROP COLUMN seats_available,
    DROP COLUMN total_seats;

ALTER TABLE cinema_hall ALTER created_by SET DEFAULT "SYSTEM"; -- query is running but not doing what I want it to.
UPDATE cinema_hall SET created_by = "System" WHERE c_id;
ALTER TABLE cinema_hall ADD COLUMN total_rooms int; 
UPDATE cinema_hall SET total_rooms = 9 WHERE c_id = 10;
DELETE FROM cinema_hall WHERE c_id=9;
SELECT * FROM cinema_hall;    


