package service

import (
	"fmt"
	"github.com/peterP1998/music-playlist-builder/model"
	"errors"
)

type SongRepositoryInterface interface {
	CreateSong(name string,length float64,genre string,artistId int64) error
	SelectSongByName(name string) (model.Song, error)
}

type SongServiceInterface interface {
	CreateSong(name string,length float64,genre string,artistId int64) error
}

type SongService struct {
	songRepository SongRepositoryInterface
}

func SongServiceCreate(songRepository SongRepositoryInterface) SongServiceInterface {
	return &SongService{
		songRepository: songRepository,
	}
}

func (songService *SongService) CreateSong(name string,length float64,genre string,artistId int64) error{
	song, err := songService.songRepository.SelectSongByName(name)
	if err != nil && fmt.Sprint(err) != "sql: no rows in result set" {
		return err
	}
	if (song!=model.Song{}) {
		return errors.New("Song already exists!")
	}
	err = songService.songRepository.CreateSong(name,length,genre,artistId)
	if err != nil {
		return err
	}
	return nil
}