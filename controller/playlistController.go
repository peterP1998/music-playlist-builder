package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/peterP1998/music-playlist-builder/model"
	"github.com/peterP1998/music-playlist-builder/model/requests"
	"github.com/peterP1998/music-playlist-builder/service"
)

type PlaylistControllerInterface interface {
	PlaylistCreate(ctx *gin.Context)
	AddSongToPlaylist(ctx *gin.Context)
	GetSongsFromPlaylist(ctx *gin.Context)
}

type PlaylistController struct {
	playlistService service.PlaylistServiceInterface
	songService     service.SongServiceInterface
}

func PlaylistControllerCreate(playlistService service.PlaylistServiceInterface, songService service.SongServiceInterface) PlaylistControllerInterface {
	return &PlaylistController{
		playlistService: playlistService,
		songService:     songService,
	}
}

func (playlistController *PlaylistController) PlaylistCreate(ctx *gin.Context) {
	var playlist model.Playlist
	err := ctx.ShouldBindJSON(&playlist)
	if err != nil {
		ctx.JSON(400, "Parameters are not ok.")
		return
	}
	err = playlistController.playlistService.CreatePlaylist(playlist.Name)
	if err != nil && fmt.Sprint(err) == "Playlist already exists!" {
		ctx.JSON(400, "Playlist already exists!")
		return
	} else if err != nil {
		ctx.JSON(500, "Something went wrong!")
		return
	} else {
		ctx.JSON(201, "Playlist created!")
		return
	}
}

func (playlistController *PlaylistController) AddSongToPlaylist(ctx *gin.Context) {
	var playlistSong requests.PlaylistSong
	err := ctx.ShouldBindJSON(&playlistSong)
	if err != nil {
		ctx.JSON(400, "Parameters are not ok.")
		return
	}
	song, err := playlistController.songService.GetSong(playlistSong.SongName)
	if err != nil || (model.Song{}==song) {
		ctx.JSON(400, "Song doesnt exists!")
		return
	}
	err = playlistController.playlistService.AddSongToPlaylist(playlistSong.PlaylistName, song.Id, song.Length)
	if err != nil {
		ctx.JSON(500, "Something went wrong!")
		return
	}
	ctx.JSON(200, "Song added!")
}

func (playlistController *PlaylistController) GetSongsFromPlaylist(ctx *gin.Context) {
	var playlist requests.PlaylistName
	err := ctx.BindQuery(&playlist)
	if err != nil {
		ctx.JSON(400, "Parameters are not ok.")
		return
	}
	songs, err := playlistController.playlistService.GetAllSongsFromPlaylist(ctx.Query("name"))
	if err != nil {
		ctx.JSON(500, "Something went wrong!")
		return
	}
	ctx.JSON(200, songs)
}
