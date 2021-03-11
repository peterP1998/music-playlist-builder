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
create table Album(
   id INT NOT NULL AUTO_INCREMENT,
   name VARCHAR(100) NOT NULL,
   year int NOT NULL,
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