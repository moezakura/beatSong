package jsio

import "github.com/moezakura/beatSong/dll/pkg/song"

type SongList struct {
	Status  bool         `json:"status"`
	Error   interface{}  `json:"error"`
	Payload []*song.Song `json:"payload"`
}
