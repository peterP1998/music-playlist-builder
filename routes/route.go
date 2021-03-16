package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/peterP1998/music-playlist-builder/controller"
	"github.com/peterP1998/music-playlist-builder/service"
	"github.com/peterP1998/music-playlist-builder/repository"
	"github.com/peterP1998/music-playlist-builder/middleware"
)

func initLoginController() controller.LoginControllerInterface{
	userRepository:=repository.UserRepository{}
	userService:=service.UserServiceCreate(userRepository)
	loginService:=service.LoginServiceAuth()
	loginController:=controller.LoginControllerCreate(loginService,userService)
    return loginController
}

func initArtistController() controller.ArtistControllerInterface{
	artistRepository:=repository.ArtistRepository{}
	artistService:=service.ArtistServiceCreate(artistRepository)
	artistController:=controller.ArtistControllerCreate(artistService)
    return artistController
}

func initSongController() controller.SongControllerInterface{
	songRepository:=repository.SongRepository{}
	artistRepository:=repository.ArtistRepository{}
	artistService:=service.ArtistServiceCreate(artistRepository)
	userRepository:=repository.UserRepository{}
	userService:=service.UserServiceCreate(userRepository)
	songService:=service.SongServiceCreate(songRepository)
	songController:=controller.SongControllerCreate(artistService,songService,userService)
    return songController
}

func initPlaylistController() controller.PlaylistControllerInterface{
	songRepository:=repository.SongRepository{}
	songService:=service.SongServiceCreate(songRepository)
	playlistRepository:=repository.PlaylistRepository{}
	playlistService:=service.PlaylistServiceCreate(playlistRepository,songRepository)
	playlistController:=controller.PlaylistControllerCreate(playlistService,songService)
    return playlistController
}

func InitRouter() *gin.Engine{
	server := gin.New()
	loginController:=initLoginController()
	songController:=initSongController()
	artistController:=initArtistController()
	playlistController:=initPlaylistController()
	server.POST("/login", loginController.Login)
	server.POST("/register", loginController.Register)
	artistRoutes := server.Group("/artist") 
	{
		artistRoutes.POST("/create", middleware.AuthorizeJWT(),artistController.ArtistCreate)
	}
	songRoutes := server.Group("/song") 
	{
		songRoutes.POST("/create", middleware.AuthorizeJWT(),songController.SongCreate)
		songRoutes.POST("/like", middleware.AuthorizeJWT(),songController.LikeSong)
		songRoutes.GET("/like", middleware.AuthorizeJWT(),songController.GetLikedSongs)
	}
	playlistRoutes := server.Group("/playlist") 
	{
		playlistRoutes.POST("/create", middleware.AuthorizeJWT(),playlistController.PlaylistCreate)
		playlistRoutes.POST("/song/add", middleware.AuthorizeJWT(),playlistController.AddSongToPlaylist)
		playlistRoutes.GET("/song", middleware.AuthorizeJWT(),playlistController.GetSongsFromPlaylist)
	}
	return server
}