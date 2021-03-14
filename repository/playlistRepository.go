package repository

import (
	"github.com/peterP1998/music-playlist-builder/model"
)

type PlaylistRepository struct {
}

func (playlistRepository PlaylistRepository) CreatePlaylist(name string) error {
	_, err := model.DB.Query("insert into Playlist(name,length,numberOfSongs) Values(?,?,?);", name,0.0,0)
	if err != nil {
		return err
	}
	return nil
}

func (playlistRepository PlaylistRepository) AddSongToPlaylist(playlistId int,songId int) error {
	_, err := model.DB.Query("insert into PlaylistSong(playlist_id,song_id) Values(?,?);", playlistId,songId)
	if err != nil {
		return err
	}
	return nil
}

func (songRepository SongRepository) SelectAllByPlaylistId(playlistid int)  ([]int,error){
	res, err := model.DB.Query("SELECT song_id FROM PlaylistSong where playlist_id=?",playlistid)
	if err != nil {
		return nil, err
	}
	songs := make([]int, 0)
	for res.Next() {
		var songId int
		res.Scan(&songId)
		songs = append(songs, songId)
	}
	return songs, nil
}

func (playlistRepository PlaylistRepository) UpdatePlaylist(name string,id int,length float64,numberOfSongs int) error {
	_, err := model.DB.Query("update Playlist SET name=?,length=?,numberOfSongs=? where id=?;", name,length,numberOfSongs,id)
	if err != nil {
		return err
	}
	return nil
}

func (playlistRepository PlaylistRepository) SelectPlaylistByName(name string) (model.Playlist, error) {
	var playlist model.Playlist
	err := model.DB.QueryRow("SELECT * FROM Playlist where name=?", name).Scan(&playlist.Id, &playlist.Name,&playlist.Length,&playlist.NumberOfSongs)
	if err != nil {
		return playlist, err
	}
	return playlist, nil
}