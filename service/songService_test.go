package service

import (
	"github.com/stretchr/testify/assert"
	//"net/http/httptest"
	"github.com/peterP1998/music-playlist-builder/service/mocks"
	"testing"
)

func TestCreateSongWithExistingSong(t *testing.T) {
	songRepositoryMock := mocks.SongRepositoryMock{}
	songService := SongServiceCreate(songRepositoryMock)
	err := songService.CreateSong("test1", 3.5,"test",3)
	assert.Equal(t, nil, err)
}

func TestCreateSongWithNonExistingSong(t *testing.T) {
	songRepositoryMock := mocks.SongRepositoryMock{}
	songService := SongServiceCreate(songRepositoryMock)
	err := songService.CreateSong("test", 3.5,"test",3)
	assert.NotEqual(t, nil, err)
}

func TestGetSongWithExistingSong(t *testing.T) {
	songRepositoryMock := mocks.SongRepositoryMock{}
	songService := SongServiceCreate(songRepositoryMock)
	song,err := songService.GetSong("test")
	assert.Equal(t, nil, err)
	assert.Equal(t, "test", song.Name)
	assert.Equal(t, 3.5, song.Length)
}

func TestGetSongWithNonExistingSong(t *testing.T) {
	songRepositoryMock := mocks.SongRepositoryMock{}
	songService := SongServiceCreate(songRepositoryMock)
	_,err := songService.GetSong("test2")
	assert.NotEqual(t, nil, err)
}

func TestAddSongWithNonExistingSong(t *testing.T) {
	songRepositoryMock := mocks.SongRepositoryMock{}
	songService := SongServiceCreate(songRepositoryMock)
	err := songService.LikeSong(2,"test1")
	assert.NotEqual(t, nil, err)
}

func TestAddSongWithExistingSong(t *testing.T) {
	songRepositoryMock := mocks.SongRepositoryMock{}
	songService := SongServiceCreate(songRepositoryMock)
	err := songService.LikeSong(2,"test")
	assert.Equal(t, nil, err)
}

func TestGetLikedSongsWithUserWithoutSongs(t *testing.T) {
	songRepositoryMock := mocks.SongRepositoryMock{}
	songService := SongServiceCreate(songRepositoryMock)
	_,err := songService.GetLikedSongs(2)
	assert.NotEqual(t, nil, err)
}

func TestGetLikedSongsWithUserWithSongs(t *testing.T) {
	songRepositoryMock := mocks.SongRepositoryMock{}
	songService := SongServiceCreate(songRepositoryMock)
	songs,err := songService.GetLikedSongs(1)
	assert.Equal(t, nil, err)
	assert.Equal(t, 1, len(songs))
}