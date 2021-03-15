package service

import (
	"errors"
	"fmt"
	"github.com/peterP1998/music-playlist-builder/model"
)

type PlaylistRepositoryInterface interface {
	CreatePlaylist(name string) error
	SelectPlaylistByName(name string) (model.Playlist, error)
	UpdatePlaylist(name string, id int, length float64, numberOfSongs int) error
	AddSongToPlaylist(playlistId int, songId int) error
	SelectAllByPlaylistId(playlistid int) ([]int, error)
}

type PlaylistServiceInterface interface {
	CreatePlaylist(name string) error
	AddSongToPlaylist(name string, songId int, songLength float64) error
	GetAllSongsFromPlaylist(playlistname string) ([]model.Song, error)
}

type PlaylistService struct {
	playlistRepository PlaylistRepositoryInterface
	songRepository     SongRepositoryInterface
}

func PlaylistServiceCreate(playlistRepository PlaylistRepositoryInterface, songRepository SongRepositoryInterface) PlaylistServiceInterface {
	return &PlaylistService{
		playlistRepository: playlistRepository,
		songRepository:     songRepository,
	}
}

func (playlistService *PlaylistService) CreatePlaylist(name string) error {
	playlist, err := playlistService.playlistRepository.SelectPlaylistByName(name)
	if err != nil && fmt.Sprint(err) != "sql: no rows in result set" {
		return err
	}
	if (playlist != model.Playlist{}) {
		return errors.New("Playlist already exists!")
	}
	err = playlistService.playlistRepository.CreatePlaylist(name)
	if err != nil {
		return err
	}
	return nil
}

func (playlistService *PlaylistService) AddSongToPlaylist(name string, songId int, songLength float64) error {
	playlist, err := playlistService.playlistRepository.SelectPlaylistByName(name)
	if err != nil {
		return err
	}
	err = playlistService.playlistRepository.UpdatePlaylist(playlist.Name, playlist.Id, playlist.Length+songLength, playlist.NumberOfSongs+1)
	if err != nil {
		return err
	}
	err = playlistService.playlistRepository.AddSongToPlaylist(playlist.Id, songId)
	if err != nil {
		return err
	}
	return nil
}

func (playlistService *PlaylistService) GetAllSongsFromPlaylist(playlistname string) ([]model.Song, error) {
	playlist, err := playlistService.playlistRepository.SelectPlaylistByName(playlistname)
	if err != nil {
		return nil, err
	}
	songIds, err := playlistService.playlistRepository.SelectAllByPlaylistId(playlist.Id)
	if err != nil {
		return nil, err
	}
	songs := make([]model.Song, 0)
	for _, id := range songIds {
		song, err := playlistService.songRepository.SelectSongById(id)
		if err != nil {
			return nil, err
		}
		songs = append(songs, song)
	}
	return songs, nil
}
