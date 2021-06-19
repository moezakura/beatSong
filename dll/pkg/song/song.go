package song

type SongDifficulty string

const (
	SongDifficultyEasy       SongDifficulty = "easy"
	SongDifficultyNormal     SongDifficulty = "normal"
	SongDifficultyHard       SongDifficulty = "hard"
	SongDifficultyExpert     SongDifficulty = "expert"
	SongDifficultyExpertPlus SongDifficulty = "expert+"
)

type Song struct {
	DirPath    string           `json:"dirPath"`
	ImagePath  string           `json:"imagePath"`
	Name       string           `json:"name"`
	Difficulty []SongDifficulty `json:"difficulty"`
}
