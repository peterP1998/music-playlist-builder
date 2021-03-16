package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/peterP1998/music-playlist-builder/model"
	"github.com/peterP1998/music-playlist-builder/model/requests"
	"github.com/peterP1998/music-playlist-builder/service"
)

type SongControllerInterface interface {
	SongCreate(ctx *gin.Context)
	LikeSong(ctx *gin.Context)
	GetLikedSongs(ctx *gin.Context)
}

type SongController struct {
	artistService service.ArtistServiceInterface
	songService   service.SongServiceInterface
	userService   service.UserServiceInterface
}

func SongControllerCreate(artistService service.ArtistServiceInterface,
	songService service.SongServiceInterface,
	userService service.UserServiceInterface) SongControllerInterface {
	return &SongController{
		artistService: artistService,
		songService:   songService,
		userService:   userService,
	}
}

func (songController *SongController) SongCreate(ctx *gin.Context) {
	var songCreate requests.SongCreate
	err := ctx.ShouldBindJSON(&songCreate)
	if err != nil {
		ctx.JSON(400, "Parameters are not ok.")
		return
	}
	artist, err := songController.artistService.GetArtist(songCreate.ArtistName)
	if err != nil || (artist==model.Artist{}){
		ctx.JSON(500, "Artist not exists!")
		return
	}
	err = songController.songService.CreateSong(songCreate.Name, songCreate.Length, songCreate.Genre, artist.Id)
	if err != nil && fmt.Sprint(err) == "Song already exists!" {
		ctx.JSON(400, "Song already exists!")
		return
	} else if err != nil {
		ctx.JSON(500, "Something went wrong!")
		return
	} else {
		ctx.JSON(201, "Song created!")
		return
	}
}

func (songController *SongController) LikeSong(ctx *gin.Context) {
	fmt.Println(ctx.GetString("username"))
	var likeSong requests.LikeSong
	err := ctx.ShouldBindJSON(&likeSong)
	if err != nil {
		ctx.JSON(400, "Parameters are not ok.")
		return
	}
	user, err := songController.userService.GetUser(ctx.GetString("username"))
	if err != nil {
		ctx.JSON(500, "Something went wrong!")
		return
	}
	err = songController.songService.LikeSong(user.Id, likeSong.SongName)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(500, "Something went wrong!")
		return
	} else {
		ctx.JSON(200, "Song liked!")
		return
	}
}

func (songController *SongController) GetLikedSongs(ctx *gin.Context) {
	user, err := songController.userService.GetUser(ctx.GetString("username"))
	if err != nil {
		ctx.JSON(500, "Something went wrong!")
		return
	}
	songs, err := songController.songService.GetLikedSongs(user.Id)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(500, "Something went wrong!")
		return
	} else {
		ctx.JSON(200, songs)
		return
	}
}
