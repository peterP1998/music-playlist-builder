package service

import (
	"errors"
	"fmt"
	"github.com/peterP1998/music-playlist-builder/model"
)

type SongRepositoryInterface interface {
	CreateSong(name string, length float64, genre string, artistId int) error
	SelectSongByName(name string) (model.Song, error)
	AddLikedSong(userid int, songid int) error
	SelectSongById(id int) (model.Song, error)
	SelectAllByUserId(userid int) ([]int, error)
}

type SongServiceInterface interface {
	CreateSong(name string, length float64, genre string, artistId int) error
	GetSong(name string) (model.Song, error)
	LikeSong(userId int, song string) error
	GetLikedSongs(userId int) ([]model.Song, error)
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

func (songService *SongService) LikeSong(userId int, songname string) error {
	song, err := songService.songRepository.SelectSongByName(songname)
	if err != nil {
		return err
	}
	err = songService.songRepository.AddLikedSong(userId, song.Id)
	if err != nil {
		return err
	}
	return nil
}

func (songService *SongService) GetLikedSongs(userId int) ([]model.Song, error) {
	songIds, err := songService.songRepository.SelectAllByUserId(userId)
	if err != nil {
		return nil, err
	}
	songs := make([]model.Song, 0)
	for _, id := range songIds {
		song, err := songService.songRepository.SelectSongById(id)
		if err != nil {
			return nil, err
		}
		songs = append(songs, song)
	}
	return songs, nil
}

