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
}

type PlaylistServiceInterface interface {
	CreatePlaylist(name string) error
	AddSongToPlaylist(name string, songId int,songLength float64) error
}

type PlaylistService struct {
	playlistRepository PlaylistRepositoryInterface
}

func PlaylistServiceCreate(playlistRepository PlaylistRepositoryInterface) PlaylistServiceInterface {
	return &PlaylistService{
		playlistRepository: playlistRepository,
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

func (playlistService *PlaylistService) AddSongToPlaylist(name string, songId int,songLength float64) error {
	playlist, err := playlistService.playlistRepository.SelectPlaylistByName(name)
	if err != nil {
		return err
	}
	err = playlistService.playlistRepository.UpdatePlaylist(playlist.Name,playlist.Id,playlist.Length+songLength,playlist.NumberOfSongs+1)
	if err != nil {
		return err
	}
	err = playlistService.playlistRepository.AddSongToPlaylist(playlist.Id,songId)
	if err != nil {
		return err
	}
	return nil
}
