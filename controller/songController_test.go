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
func TestSongController(t *testing.T) {
	songRepository:=mocks.SongRepositoryMock{}
	artistRepository:=mocks.ArtistRepositoryMock{}
	artistService:=service.ArtistServiceCreate(artistRepository)
	userRepository:=mocks.UserRepositoryMock{}
	userService:=service.UserServiceCreate(userRepository)
	songService:=service.SongServiceCreate(songRepository)
	songController:=SongControllerCreate(artistService,songService,userService)
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/song/create", songController.SongCreate)
	router.POST("/song/like" ,songController.LikeSong)
	router.GET("/song/like",songController.GetLikedSongs)
	t.Run("song create with non existsing song", func(t *testing.T) {
		resp := httptest.NewRecorder()
		var jsonStr = []byte(`{"artistname":"test","name":"test1","length":4.15,"genre":"ROCK"}`)
		req, _ := http.NewRequest("POST", "/song/create", bytes.NewBuffer(jsonStr))
		router.ServeHTTP(resp, req)
		assert.Equal(t, 201, resp.Result().StatusCode)
	})
	t.Run("song create with existsing song", func(t *testing.T) {
		resp := httptest.NewRecorder()
		var jsonStr = []byte(`{"artistname":"test","name":"test","length":4.15,"genre":"ROCK"}`)
		req, _ := http.NewRequest("POST", "/song/create", bytes.NewBuffer(jsonStr))
		router.ServeHTTP(resp, req)
		assert.Equal(t, 400, resp.Result().StatusCode)
	})
	t.Run("like song with non existsing song", func(t *testing.T) {
		resp := httptest.NewRecorder()
		var jsonStr = []byte(`{"songname":"test1"}`)
		req, _ := http.NewRequest("POST", "/song/like", bytes.NewBuffer(jsonStr))
		router.ServeHTTP(resp, req)
		assert.Equal(t, 500, resp.Result().StatusCode)
	})
	t.Run("like song with existsing song", func(t *testing.T) {
		resp := httptest.NewRecorder()
		var jsonStr = []byte(`{"songname":"test"}`)
		req, _ := http.NewRequest("POST", "/song/like", bytes.NewBuffer(jsonStr))
		router.ServeHTTP(resp, req)
		assert.Equal(t, 200, resp.Result().StatusCode)
	})
	t.Run("get liked songs", func(t *testing.T) {
		resp := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/song/like", nil)
		router.ServeHTTP(resp, req)
		assert.Equal(t, 200, resp.Result().StatusCode)
	})
}