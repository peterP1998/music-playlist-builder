package service

import (
	"github.com/stretchr/testify/assert"
	//"net/http/httptest"
	"github.com/peterP1998/music-playlist-builder/service/mocks"
	"testing"
)
func TestCreateSongWithExistingPlaylist(t *testing.T) {
	songRepositoryMock := mocks.SongRepositoryMock{}
	playlistRepository := mocks.PlaylistRepositoryMock{}
	playlistService := PlaylistServiceCreate(playlistRepository,songRepositoryMock)
	err := playlistService.CreatePlaylist("test")
	assert.NotEqual(t, nil, err)
}

func TestCreateSongWithNonExistingPlaylist(t *testing.T) {
	songRepositoryMock := mocks.SongRepositoryMock{}
	playlistRepository := mocks.PlaylistRepositoryMock{}
	playlistService := PlaylistServiceCreate(playlistRepository,songRepositoryMock)
	err := playlistService.CreatePlaylist("test1")
	assert.Equal(t, nil, err)
}

func TestAddSongToPlaylistWithNonExistingPlaylist(t *testing.T) {
	songRepositoryMock := mocks.SongRepositoryMock{}
	playlistRepository := mocks.PlaylistRepositoryMock{}
	playlistService := PlaylistServiceCreate(playlistRepository,songRepositoryMock)
	err := playlistService.AddSongToPlaylist("test1",2,3.5)
	assert.NotEqual(t, nil, err)
}

func TestAddSongToPlaylistWithExistingPlaylist(t *testing.T) {
	songRepositoryMock := mocks.SongRepositoryMock{}
	playlistRepository := mocks.PlaylistRepositoryMock{}
	playlistService := PlaylistServiceCreate(playlistRepository,songRepositoryMock)
	err := playlistService.AddSongToPlaylist("test",2,3.5)
	assert.NotEqual(t, nil, err)
}

func TestGetAllSongsFromPlaylistWithExistingPlaylist(t *testing.T) {
	songRepositoryMock := mocks.SongRepositoryMock{}
	playlistRepository := mocks.PlaylistRepositoryMock{}
	playlistService := PlaylistServiceCreate(playlistRepository,songRepositoryMock)
	err := playlistService.AddSongToPlaylist("test",2,3.5)
	assert.NotEqual(t, nil, err)
}

func TestGetAllSongsFromPlaylistWithNotExistingPlaylist(t *testing.T) {
	songRepositoryMock := mocks.SongRepositoryMock{}
	playlistRepository := mocks.PlaylistRepositoryMock{}
	playlistService := PlaylistServiceCreate(playlistRepository,songRepositoryMock)
	err := playlistService.AddSongToPlaylist("test1",2,3.5)
	assert.NotEqual(t, nil, err)
}
