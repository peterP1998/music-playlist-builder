package main

import (
	"github.com/peterP1998/music-playlist-builder/routes"
	"github.com/peterP1998/music-playlist-builder/model"
)

func main() {
	model.InitDb()
	server := routes.InitRouter()
	port := "8080"
	server.Run(":" + port)

}