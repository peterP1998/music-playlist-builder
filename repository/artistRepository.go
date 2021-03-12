package repository

import (
	"github.com/peterP1998/music-playlist-builder/model"
)

type ArtistRepository struct {
}

func (artistRepository ArtistRepository) CreateArtist(name string) error {
	_, err := model.DB.Query("insert into Artist(name) Values(?);", name)
	if err != nil {
		return err
	}
	return nil
}

func (artistRepository ArtistRepository) SelectArtistByName(name string) (model.Artist, error) {
	var artist model.Artist
	err := model.DB.QueryRow("SELECT * FROM Artist where name=?", name).Scan(&artist.Id, &artist.Name)
	if err != nil {
		return artist, err
	}
	return artist, nil
}
