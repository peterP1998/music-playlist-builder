package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/peterP1998/music-playlist-builder/model"
	"github.com/peterP1998/music-playlist-builder/service"
)

type ArtistControllerInterface interface {
	ArtistCreate(ctx *gin.Context)
}

type ArtistController struct {
	artistService service.ArtistServiceInterface
}

func ArtistControllerCreate(artistService service.ArtistServiceInterface) ArtistControllerInterface {
	return &ArtistController{
		artistService: artistService,
	}
}

func (artistController *ArtistController) ArtistCreate(ctx *gin.Context) {
	var artist model.Artist
	err := ctx.ShouldBindJSON(&artist)
	if err != nil {
		ctx.JSON(400, "Parameters are not ok.")
		return
	}
	err = artistController.artistService.CreateArtist(artist.Name)
	if err != nil && fmt.Sprint(err) == "Artist already exists!" {
		ctx.JSON(400, "Artist already exists!")
		return
	} else if err != nil {
		ctx.JSON(500, "Something went wrong!")
		return
	} else {
		ctx.JSON(201, "Artist created!")
		return
	}
}
