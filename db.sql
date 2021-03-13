create database music;
use music;
create table User(
   id INT NOT NULL AUTO_INCREMENT,
   email VARCHAR(100) NOT NULL,
   username VARCHAR(30) NOT NULL,
   password VARCHAR(100) NOT NULL,
   PRIMARY KEY ( id )
);
create table Artist(
   id INT NOT NULL AUTO_INCREMENT,
   name VARCHAR(100) NOT NULL,
   PRIMARY KEY ( id )
);
create table Playlist(
   id INT NOT NULL AUTO_INCREMENT,
   name VARCHAR(100) NOT NULL,
   numberOfSongs int NOT NULL,
   length float NOT NULL,
   PRIMARY KEY ( id )
);
create table Song(
   id INT NOT NULL AUTO_INCREMENT,
   name VARCHAR(100) NOT NULL,
   length FLOAT NOT NULL,
   genre VARCHAR(100) NOT NULL,
   artist_id INT NOT NULL,
   PRIMARY KEY ( id ),
   FOREIGN KEY (artist_id) REFERENCES Artist(id)
);
create table PlaylistSong(
   id INT NOT NULL AUTO_INCREMENT,
   playlist_id INT NOT NULL,
   song_id INT NOT NULL,
   primary key(id),
   FOREIGN KEY (playlist_id) REFERENCES Playlist(id),
   FOREIGN KEY (song_id) REFERENCES Song(id)
);