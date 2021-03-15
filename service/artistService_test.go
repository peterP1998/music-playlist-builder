package service

import (
	"github.com/stretchr/testify/assert"
	//"net/http/httptest"
	"github.com/peterP1998/music-playlist-builder/service/mocks"
	"testing"
)

func TestCreateArtistWithExistingArtist(t *testing.T) {
	artistRepositoryMock := mocks.ArtistRepositoryMock{}
	artistService := ArtistServiceCreate(artistRepositoryMock)
	err := artistService.CreateArtist("test1")
	assert.NotEqual(t, nil, err)
}

func TestCreateArtistWithNonExistingArtist(t *testing.T) {
	artistRepositoryMock := mocks.ArtistRepositoryMock{}
	artistService := ArtistServiceCreate(artistRepositoryMock)
	err := artistService.CreateArtist("test21")
	assert.Equal(t, nil, err)
}

func TestGetArtistWithNonExistingArtist(t *testing.T) {
	artistRepositoryMock := mocks.ArtistRepositoryMock{}
	artistService := ArtistServiceCreate(artistRepositoryMock)
	_, err := artistService.GetArtist("test2")
	assert.NotEqual(t, nil, err)
}

func TestGetArtistWithExistingArtist(t *testing.T) {
	artistRepositoryMock := mocks.ArtistRepositoryMock{}
	artistService := ArtistServiceCreate(artistRepositoryMock)
	_, err := artistService.GetArtist("test")
	assert.Equal(t, nil, err)
}
