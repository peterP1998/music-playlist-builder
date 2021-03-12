package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/peterP1998/music-playlist-builder/model/requests"
	"github.com/peterP1998/music-playlist-builder/service"
)

type SongControllerInterface interface {
	SongCreate(ctx *gin.Context)
}

type SongController struct {
	artistService service.ArtistServiceInterface
	songService service.SongServiceInterface
}

func SongControllerCreate(artistService service.ArtistServiceInterface,songService service.SongServiceInterface) SongControllerInterface {
	return &SongController{
		artistService: artistService,
		songService:songService,
	}
}

func (songController *SongController) SongCreate(ctx *gin.Context) {
	var songCreate requests.SongCreate
	err := ctx.ShouldBindJSON(&songCreate)
	if err != nil {
		ctx.JSON(400, "Parameters are not ok.")
		return
	}
	artist,err:=songController.artistService.GetArtist(songCreate.ArtistName)
	if err != nil {
		ctx.JSON(500, "Artist not exists!")
		return
	}
	err=songController.songService.CreateSong(songCreate.Name,songCreate.Length,songCreate.Genre,artist.Id)
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