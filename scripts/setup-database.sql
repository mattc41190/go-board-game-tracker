CREATE DATABASE IF NOT EXISTS board_game_tracker_db;

USE board_game_tracker_db;

CREATE TABLE IF NOT EXISTS users(
    id INT AUTO_INCREMENT,
    email VARCHAR(200) NOT NULL,
    password VARCHAR(100) NOT NULL,
    first_name VARCHAR(50) DEFAULT "Lamey",
    last_name VARCHAR(50) DEFAULT "McLameface",
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS games(
    id INT AUTO_INCREMENT,
    user_id INT NOT NULL,
    title VARCHAR(255) NOT NULL,
    played BOOLEAN NOT NULL,
    link VARCHAR(50) DEFAULT "https://www.reddit.com/r/boardgames/wiki/index",
    rating INT DEFAULT 5,
    PRIMARY KEY(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);