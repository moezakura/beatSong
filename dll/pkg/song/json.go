package song

type SongInfoJSON struct {
	Version               string                       `json:"_version"`
	SongName              string                       `json:"_songName"`
	SongSubName           string                       `json:"_songSubName"`
	SongAuthorName        string                       `json:"_songAuthorName"`
	LevelAuthorName       string                       `json:"_levelAuthorName"`
	BeatsPerMinute        float64                      `json:"_beatsPerMinute"`
	SongTimeOffset        float64                      `json:"_songTimeOffset"`
	Shuffle               float64                      `json:"_shuffle"`
	ShufflePeriod         float64                      `json:"_shufflePeriod"`
	PreviewStartTime      float64                      `json:"_previewStartTime"`
	PreviewDuration       float64                      `json:"_previewDuration"`
	SongFilename          string                       `json:"_songFilename"`
	CoverImageFilename    string                       `json:"_coverImageFilename"`
	EnvironmentName       string                       `json:"_environmentName"`
	CustomData            *CustomDataJSON              `json:"_customData"`
	DifficultyBeatmapSets []*DifficultyBeatmapSetsJSON `json:"_difficultyBeatmapSets"`
}
type CustomDataJSON struct {
	Contributors          []interface{} `json:"_contributors"`
	CustomEnvironment     string        `json:"_customEnvironment"`
	CustomEnvironmentHash string        `json:"_customEnvironmentHash"`
}
type SongCustomdataJSON struct {
	DifficultyLabel string        `json:"_difficultyLabel"`
	EditorOffset    float64       `json:"_editorOffset"`
	EditorOldOffset float64       `json:"_editorOldOffset"`
	Warnings        []interface{} `json:"_warnings"`
	Information     []interface{} `json:"_information"`
	Suggestions     []interface{} `json:"_suggestions"`
	Requirements    []interface{} `json:"_requirements"`
}
type DifficultyBeatmapsJSON struct {
	Difficulty              string              `json:"_difficulty"`
	DifficultyRank          int                 `json:"_difficultyRank"`
	BeatmapFilename         string              `json:"_beatmapFilename"`
	NoteJumpMovementSpeed   float64             `json:"_noteJumpMovementSpeed"`
	NoteJumpStartBeatOffset float64             `json:"_noteJumpStartBeatOffset"`
	CustomData              *SongCustomdataJSON `json:"_customData"`
}
type DifficultyBeatmapSetsJSON struct {
	BeatmapCharacteristicName string                    `json:"_beatmapCharacteristicName"`
	DifficultyBeatmaps        []*DifficultyBeatmapsJSON `json:"_difficultyBeatmaps"`
}
