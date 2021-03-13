package service

import (
	"errors"
	"fmt"
	"github.com/peterP1998/music-playlist-builder/model"
)

type SongRepositoryInterface interface {
	CreateSong(name string, length float64, genre string, artistId int) error
	SelectSongByName(name string) (model.Song, error)
}

type SongServiceInterface interface {
	CreateSong(name string, length float64, genre string, artistId int) error
	GetSong(name string) (model.Song, error)
}

type SongService struct {
	songRepository SongRepositoryInterface
}

func SongServiceCreate(songRepository SongRepositoryInterface) SongServiceInterface {
	return &SongService{
		songRepository: songRepository,
	}
}

func (songService *SongService) CreateSong(name string, length float64, genre string, artistId int) error {
	song, err := songService.songRepository.SelectSongByName(name)
	if err != nil && fmt.Sprint(err) != "sql: no rows in result set" {
		return err
	}
	if (song != model.Song{}) {
		return errors.New("Song already exists!")
	}
	err = songService.songRepository.CreateSong(name, length, genre, artistId)
	if err != nil {
		return err
	}
	return nil
}

func (songService *SongService) GetSong(name string) (model.Song, error) {
	song, err := songService.songRepository.SelectSongByName(name)
	if err != nil {
		return song, err
	}
	return song, nil
}
