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
func TestPlaylistController(t *testing.T) {
	songRepository:=mocks.SongRepositoryMock{}
	songService:=service.SongServiceCreate(songRepository)
	playlistRepository:=mocks.PlaylistRepositoryMock{}
	playlistService:=service.PlaylistServiceCreate(playlistRepository,songRepository)
	playlistController:=PlaylistControllerCreate(playlistService,songService)
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/playlist/create", playlistController.PlaylistCreate)
	router.POST("/playlist/song/add", playlistController.AddSongToPlaylist)
	router.GET("/playlist/song", playlistController.GetSongsFromPlaylist)
	t.Run("playlist create with non existsing playlist", func(t *testing.T) {
		resp := httptest.NewRecorder()
		var jsonStr = []byte(`{"name":"test1"}`)
		req, _ := http.NewRequest("POST", "/playlist/create", bytes.NewBuffer(jsonStr))
		router.ServeHTTP(resp, req)
		assert.Equal(t, 201, resp.Result().StatusCode)
	})
	t.Run("playlist create with existsing playlist", func(t *testing.T) {
		resp := httptest.NewRecorder()
		var jsonStr = []byte(`{"name":"test"}`)
		req, _ := http.NewRequest("POST", "/playlist/create", bytes.NewBuffer(jsonStr))
		router.ServeHTTP(resp, req)
		assert.Equal(t, 400, resp.Result().StatusCode)
	})
	t.Run("playlist add song with existsing playlist and song", func(t *testing.T) {
		resp := httptest.NewRecorder()
		var jsonStr = []byte(`{"playlistname":"test","songname":"test"}`)
		req, _ := http.NewRequest("POST", "/playlist/song/add", bytes.NewBuffer(jsonStr))
		router.ServeHTTP(resp, req)
		assert.Equal(t, 200, resp.Result().StatusCode)
	})
	t.Run("playlist add songwith non existsing playlist", func(t *testing.T) {
		resp := httptest.NewRecorder()
		var jsonStr = []byte(`{"playlistname":"test2","songname":"test"}`)
		req, _ := http.NewRequest("POST", "/playlist/song/add", bytes.NewBuffer(jsonStr))
		router.ServeHTTP(resp, req)
		assert.Equal(t, 500, resp.Result().StatusCode)
	})
	t.Run("playlist add song with non existsing song", func(t *testing.T) {
		resp := httptest.NewRecorder()
		var jsonStr = []byte(`{"playlistname":"test","songname":"test1"}`)
		req, _ := http.NewRequest("POST", "/playlist/song/add", bytes.NewBuffer(jsonStr))
		router.ServeHTTP(resp, req)
		assert.Equal(t, 400, resp.Result().StatusCode)
	})
	t.Run("get song from playlist with non existsing playlist", func(t *testing.T) {
		resp := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/playlist/song?name=test1", nil)
		router.ServeHTTP(resp, req)
		assert.Equal(t, 500, resp.Result().StatusCode)
	})
	t.Run("get song from playlist with existsing playlist", func(t *testing.T) {
		resp := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/playlist/song?name=test", nil)
		router.ServeHTTP(resp, req)
		assert.Equal(t, 200, resp.Result().StatusCode)
	})
}