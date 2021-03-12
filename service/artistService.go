package service

import (
	"fmt"
	"github.com/peterP1998/music-playlist-builder/model"
	"errors"
)

type ArtistRepositoryInterface interface {
	CreateArtist(name string) error
	SelectArtistByName(name string) (model.Artist, error)
}

type ArtistServiceInterface interface {
	CreateArtist(name string) error
	GetArtist(name string) (model.Artist,error)
}

type ArtistService struct {
	artistRepository ArtistRepositoryInterface
}

func ArtistServiceCreate(artistRepository ArtistRepositoryInterface) ArtistServiceInterface {
	return &ArtistService{
		artistRepository: artistRepository,
	}
}
func (artistService *ArtistService) CreateArtist(name string) error {
	artist, err := artistService.artistRepository.SelectArtistByName(name)
	if err != nil && fmt.Sprint(err) != "sql: no rows in result set" {
		return err
	}
	if (artist!=model.Artist{}) {
		return errors.New("Artist already exists!")
	}
	err = artistService.artistRepository.CreateArtist(name)
	if err != nil {
		return err
	}
	return nil
}

func (artistService *ArtistService) GetArtist(name string) (model.Artist,error) {
	artist, err := artistService.artistRepository.SelectArtistByName(name)
	if err != nil {
		return artist,err
	}
	return artist,nil
}