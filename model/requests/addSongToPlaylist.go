package requests

type PlaylistSong struct {
	PlaylistName string `json:"playlistname"`
	SongName     string `json:"songname"`
}
