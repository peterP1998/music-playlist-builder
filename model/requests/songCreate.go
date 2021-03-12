package requests

type SongCreate struct {
	Name       string  `json:"name"`
	Length     float64 `json:"length"`
	Genre      string  `json:"genre"`
	ArtistName string  `json:"artistname"`
}
