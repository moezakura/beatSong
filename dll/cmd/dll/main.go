package main

import "C"
import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	"os"

	"github.com/moezakura/beatSong/dll/pkg/jsio"
	"github.com/moezakura/beatSong/dll/pkg/song"
	"golang.org/x/image/draw"
	"golang.org/x/xerrors"
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
		return &jsio.SongList{
			Status:  true,
			Payload: result,
		}
	}()

	j, err := json.Marshal(res)
	if err != nil {
		return C.CString(err.Error())
	}

	return C.CString(string(j))
}

//export getImage
func getImage(imagePath *C.char) *C.char {
	ip := C.GoString(imagePath)

	img, err := func() (string, error) {
		f, err := os.Open(ip)
		if err != nil {
			return "", xerrors.Errorf("failed to open image: %w", err)
		}
		defer f.Close()

		img, _, err := image.Decode(f)
		if err != nil {
			return "", xerrors.Errorf("failed to decode image: %w", err)
		}

		ss := img.Bounds()
		ds := 80

		imgDist := image.NewRGBA64(image.Rect(0, 0, ds, ds))
		draw.CatmullRom.Scale(imgDist, imgDist.Bounds(), img, ss, draw.Over, nil)

		var buf bytes.Buffer
		if err := png.Encode(&buf, imgDist); err != nil {
			return "", xerrors.Errorf("failed to encode image to png: %w", err)
		}

		b64 := base64.StdEncoding.EncodeToString(buf.Bytes())

		return fmt.Sprintf("data:image/png;base64,%s", b64), nil
	}()

	if err != nil {
		return C.CString(err.Error())
	}

	return C.CString(img)
}
