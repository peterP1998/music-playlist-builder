package mocks

import (
	"errors"
	"github.com/peterP1998/music-playlist-builder/model"
)

type SongRepositoryMock struct {
}

func (ur SongRepositoryMock) SelectSongByName(name string) (model.Song, error) {
	var song model.Song
	if name == "test" {
		return getSong(1, "test", "test", 3.5), nil
	} else if name == "test2" {
		return song, errors.New("Song not found")
	} else {
		return model.Song{}, nil
	}
}

func (ur SongRepositoryMock) CreateSong(name string, length float64, genre string, artistId int) error {
	if name == "test1" {
		return nil
	}
	return errors.New("Song is not created")
}

func (ur SongRepositoryMock) SelectSongById(id int) (model.Song, error) {
	var song model.Song
	if id == 1 {
		return getSong(1, "test", "test", 3.5), nil
	} else if id == 2 {
		return song, errors.New("Song not found")
	} else {
		return model.Song{}, nil
	}
}

func (ur SongRepositoryMock) SelectAllByUserId(userid int) ([]int, error) {
	songs := make([]int, 0)
	if userid == 1 {
		songs=append(songs,1)
		return songs, nil
	} else if userid == 2 {
		return songs, errors.New("Song not found")
	} else {
		return songs, nil
	}
}

func (ur SongRepositoryMock) AddLikedSong(userid int, songid int) error  {
	if songid == 1{
		return nil
	}
	return errors.New("Not added")
}

func getSong(id int, name string, genre string, length float64) model.Song {
	var song model.Song
	song.Id = id
	song.Name = name
	song.Length = length
	song.Genre = genre
	return song
}
