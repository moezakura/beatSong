package song

import (
	"testing"
)

func TestDir_GetDirList(t *testing.T) {
	t.Run(":", func(t *testing.T) {
		d := &SongParser{}
		d.GetList("C:\\Program Files (x86)\\Steam\\steamapps\\common\\Beat Saber\\Beat Saber_Data\\CustomLevels\\")
	})
}
