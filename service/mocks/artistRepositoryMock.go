package mocks

import (
	"errors"
	"github.com/peterP1998/music-playlist-builder/model"
)

type ArtistRepositoryMock struct {
}

func (ur ArtistRepositoryMock) SelectArtistByName(name string) (model.Artist, error) {
	var artist model.Artist
	if name == "test" {
		return getArtist(1, "test"), nil
	} else if name == "test2" {
		return artist, errors.New("Artist not found")
	} else {
		return model.Artist{}, nil
	}
}

func (ur ArtistRepositoryMock) CreateArtist(name string) error {
	if name == "test21" {
		return nil
	}
	return errors.New("Artist is not created")
}

func getArtist(id int, name string) model.Artist {
	var artist model.Artist
	artist.Id = id
	artist.Name = name
	return artist
}
