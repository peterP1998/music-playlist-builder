#  Music Playlist
> Technologies used Golang,Gin and MySql

![](gopher.jpg)

Music Playlist application is simple REST API written in golang using gin framework for educational purposes. In this application you can create artist,create song,like song and create playlist of your favourite songs.

# Instalation
1. The First step is to clone this repo on your pc 
``` https://github.com/peterP1998/music-playlist-builder.git ```
2. The Second step is to install mysql on your pc
3. You can find example configuration of db in db.yaml file. If your db runs on different port or something else is different for your db,
you should make a change in db.yaml file. 
4. Then execute db.sql file to create the database
5. To run the application you should type go run main.go. The server will run on port 8080

# Endpoints
I will list all the endpoints which this application has.

### Post Requests
/register - this endpoint register new user <br>
/login - this endpoint checks does this user was already register and returns jwt token  <br>
/artist/create - create new artist ex.AC/DC,The Beatles or Tina Turner  <br>
/song/create - create new song for artist  <br>
/song/like - add song to your list of liked songs  <br>
/playlist/create - create playlist of song  <br>
/playlist/song/add - add song to the playlist  <br>

### Get Requests
/song/like - get all of your liked songs  <br>
/playlist?name=* - get all songs for playlist  <br>

# Final Words
I created this project to learn basics of gin framework.

