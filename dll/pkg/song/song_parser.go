package song

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"

	"golang.org/x/xerrors"
)

type SongParser struct {
}

func NewSongParser() *SongParser {
	return &SongParser{}
}

func (s *SongParser) GetList(root string) ([]*Song, error) {
	songDir, err := s.getDirs(root)
	if err != nil {
		return nil, xerrors.Errorf("failed to get dir: %w", err)
	}

	mu := &sync.Mutex{}
	wg := new(sync.WaitGroup)
	st := make(chan struct{}, 30)
	songsInfo := make([]*Song, 0, len(songDir))
	sir := regexp.MustCompile("^[1-9a-f][0-9a-f]*$")

	for _, f := range songDir {
		st <- struct{}{}
		wg.Add(1)
		go func(f string) {
			defer func() {
				wg.Done()
				<-st
			}()
			i := filepath.Join(f, "info.dat")
			if !s.exists(i) {
				log.Printf("not found 'info.dat': %s", i)
				return
			}

			infoData, err := os.ReadFile(i)
			if err != nil {
				log.Printf("failed to read info.dat: %+v", err)
				return
			}

			var ij SongInfoJSON
			err = json.Unmarshal(infoData, &ij)
			if err != nil {
				log.Printf("failed to parse info.dat: %+v", err)
				return
			}

			var standardDifficulties []*DifficultyBeatmapsJSON
			for _, d := range ij.DifficultyBeatmapSets {
				if strings.ToLower(d.BeatmapCharacteristicName) != "standard" {
					continue
				}
				standardDifficulties = d.DifficultyBeatmaps
			}

			sds := make([]SongDifficulty, len(standardDifficulties))
			for i, d := range standardDifficulties {
				switch strings.ToLower(d.Difficulty) {
				case "easy":
					sds[i] = SongDifficultyEasy
				case "normal":
					sds[i] = SongDifficultyNormal
				case "hard":
					sds[i] = SongDifficultyHard
				case "expert":
					sds[i] = SongDifficultyExpert
				case "expertplus":
					sds[i] = SongDifficultyExpertPlus
				}
			}

			dirName := filepath.Base(f)
			id := strings.Split(dirName, " ")[0]
			isValidID := sir.MatchString(id)
			if !isValidID {
				id = "UNKNOWN"
			}

			si := &Song{
				ID:         id,
				IsValidID:  isValidID,
				DirPath:    f,
				ImagePath:  filepath.Join(f, ij.CoverImageFilename),
				Name:       ij.SongName,
				Difficulty: sds,
			}

			mu.Lock()
			songsInfo = append(songsInfo, si)
			mu.Unlock()
		}(f)
	}

	wg.Wait()
	close(st)

	return songsInfo, nil
}

func (s *SongParser) getDirs(root string) ([]string, error) {
	files, err := os.ReadDir(root)
	if err != nil {
		return nil, xerrors.Errorf("failed to read dir: %w", err)
	}

	songDirs := make([]string, 0, len(files))
	for _, f := range files {
		if !f.IsDir() {
			continue
		}
		p := filepath.Join(root, f.Name())
		songDirs = append(songDirs, p)
	}

	return songDirs, nil
}

func (s *SongParser) exists(i string) bool {
	_, err := os.Stat(i)
	return err == nil
}
