package mocks

import (
	"errors"
	"github.com/peterP1998/music-playlist-builder/model"
)

type PlaylistRepositoryMock struct {
}

func (ur PlaylistRepositoryMock) SelectPlaylistByName(name string) (model.Playlist, error) {
	var playlist model.Playlist
	if name == "test" {
		return getPlaylist(1, "test", 3, 3.5), nil
	} else if name == "test2" {
		return playlist, errors.New("Song not found")
	} else {
		return model.Playlist{}, nil
	}
}

func (ur PlaylistRepositoryMock) CreatePlaylist(name string) error {
	if name == "test1" {
		return nil
	}
	return errors.New("Playlist is not created")
}

func (ur PlaylistRepositoryMock) UpdatePlaylist(name string, id int, length float64, numberOfSongs int) error {
	if name == "test1" {
		return nil
	}
	return errors.New("Playlist is not updated")
}

func (ur PlaylistRepositoryMock) SelectAllByPlaylistId(playlistid int) ([]int, error) {
	songs := make([]int, 0)
	if playlistid == 1 {
		songs=append(songs,1)
		return songs, nil
	} else if playlistid == 2 {
		return songs, errors.New("Playlist not found")
	} else {
		return songs, nil
	}
}

func (ur PlaylistRepositoryMock) AddSongToPlaylist(playlistId int, songid int) error  {
	if songid == 1{
		return nil
	}
	return errors.New("Not added")
}

func getPlaylist(id int, name string, numberOfSongs int, length float64) model.Playlist {
	var playlist model.Playlist
	playlist.Id = id
	playlist.Name = name
	playlist.Length = length
	playlist.NumberOfSongs = numberOfSongs
	return playlist
}
