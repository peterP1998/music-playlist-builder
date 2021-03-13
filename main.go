package main

import (
	"github.com/gin-gonic/gin"
	"github.com/peterP1998/music-playlist-builder/controller"
	"github.com/peterP1998/music-playlist-builder/service"
	"github.com/peterP1998/music-playlist-builder/repository"
	"github.com/peterP1998/music-playlist-builder/middleware"
	"github.com/peterP1998/music-playlist-builder/model"
	"database/sql"
)

func main() {
	model.DB,_=sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/music")
	server := gin.New()
	var userRepository repository.UserRepository=repository.UserRepository{}
	var userService service.UserServiceInterface=service.UserServiceCreate(userRepository)
	var loginService service.LoginServiceInterface=service.LoginServiceAuth()
	var loginController controller.LoginControllerInterface=controller.LoginControllerCreate(loginService,userService)
	artistRepository:=repository.ArtistRepository{}
	artistService:=service.ArtistServiceCreate(artistRepository)
	artistController:=controller.ArtistControllerCreate(artistService)
	songRepository:=repository.SongRepository{}
	songService:=service.SongServiceCreate(songRepository)
	songController:=controller.SongControllerCreate(artistService,songService)
	playlistRepository:=repository.PlaylistRepository{}
	playlistService:=service.PlaylistServiceCreate(playlistRepository)
	playlistController:=controller.PlaylistControllerCreate(playlistService,songService)
	server.POST("/login", loginController.Login)
	server.POST("/register", loginController.Register)
	artistRoutes := server.Group("/artist") 
	{
		artistRoutes.POST("/create", middleware.AuthorizeJWT(),artistController.ArtistCreate)
	}
	songRoutes := server.Group("/song") 
	{
		songRoutes.POST("/create", middleware.AuthorizeJWT(),songController.SongCreate)
	}
	playlistRoutes := server.Group("/playlist") 
	{
		playlistRoutes.POST("/create", middleware.AuthorizeJWT(),playlistController.PlaylistCreate)
		playlistRoutes.POST("/song/add", middleware.AuthorizeJWT(),playlistController.AddSongToPlaylist)
	}
	port := "8080"
	server.Run(":" + port)

}