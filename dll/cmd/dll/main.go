package main

import "C"
import (
	"encoding/json"

	"github.com/moezakura/beatSong/dll/pkg/jsio"
	"github.com/moezakura/beatSong/dll/pkg/song"
)

func main() {

}

//export songList
func songList(root *C.char) *C.char {
	sp := song.NewSongParser()

	res := func() *jsio.SongList {
		result, err := sp.GetList(C.GoString(root))

		if err != nil {
			return &jsio.SongList{
				Status: false,
				Error:  err,
			}
		}
		return &jsio.SongList{Status: true,
			Payload: result}
	}()

	j, err := json.Marshal(res)
	if err != nil {
		return C.CString(err.Error())
	}

	return C.CString(string(j))
}
