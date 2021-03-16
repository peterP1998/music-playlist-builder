package controller

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	//"fmt"
	"github.com/peterP1998/music-playlist-builder/service"
	"github.com/peterP1998/music-playlist-builder/service/mocks"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestArtistCreate(t *testing.T) {
	artistRepositoryMock := mocks.ArtistRepositoryMock{}
	artistService := service.ArtistServiceCreate(artistRepositoryMock)
	artistController := ArtistControllerCreate(artistService)
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/artist/create", artistController.ArtistCreate)
	t.Run("with non existsing artist", func(t *testing.T) {
		resp := httptest.NewRecorder()
		var jsonStr = []byte(`{"name":"test21"}`)
		req, _ := http.NewRequest("POST", "/artist/create", bytes.NewBuffer(jsonStr))
		router.ServeHTTP(resp, req)
		assert.Equal(t, 201, resp.Result().StatusCode)
	})
	t.Run("with existsing artist", func(t *testing.T) {
		resp := httptest.NewRecorder()
		var jsonStr = []byte(`{"name":"test"}`)
		req, _ := http.NewRequest("POST", "/artist/create", bytes.NewBuffer(jsonStr))
		router.ServeHTTP(resp, req)
		assert.Equal(t, 400, resp.Result().StatusCode)
	})
}

