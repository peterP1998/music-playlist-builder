create database music;
use music;
create table User(
   id INT NOT NULL AUTO_INCREMENT,
   email VARCHAR(100) NOT NULL,
   username VARCHAR(30) NOT NULL,
   password VARCHAR(100) NOT NULL,
   PRIMARY KEY ( id )
);