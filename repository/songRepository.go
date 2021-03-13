package repository

import (
	"github.com/peterP1998/music-playlist-builder/model"
)

type SongRepository struct {
}

func (songRepository SongRepository) CreateSong(name string,length float64,genre string,artistId int) error {
	_, err := model.DB.Query("insert into Song(name,length,genre,artist_id) Values(?,?,?,?);", name,length,genre,artistId)
	if err != nil {
		return err
	}
	return nil
}

func (songRepository SongRepository) SelectSongByName(name string) (model.Song, error) {
	var song model.Song
	err := model.DB.QueryRow("SELECT * FROM Song where name=?", name).Scan(&song.Id, &song.Name,&song.Length,&song.Genre,&song.ArtistId)
	if err != nil {
		return song, err
	}
	return song, nil
}
