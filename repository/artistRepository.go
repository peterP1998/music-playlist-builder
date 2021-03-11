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