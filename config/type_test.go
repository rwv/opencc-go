package config

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	var config Config
	filename := "../raw-data/config/hk2s.json"

	data, err := os.ReadFile(filename)

	assert.Nil(t, err)

	err = json.Unmarshal(data, &config)

	assert.Nil(t, err)

	assert.Equal(t, "Traditional Chinese (Hong Kong standard) to Simplified Chinese", config.Name)
	assert.Equal(t, "mmseg", config.Segmentation.Type)
	assert.Equal(t, "TSPhrases.txt", config.Segmentation.Dict.File)

	assert.Equal(t, "HKVariantsRevPhrases.txt", config.ConversionChain[0].Dict.Dicts[0].File)
	assert.Equal(t, "HKVariantsRev.txt", config.ConversionChain[0].Dict.Dicts[1].File)

	assert.Equal(t, "TSPhrases.txt", config.ConversionChain[1].Dict.Dicts[0].File)
	assert.Equal(t, "TSCharacters.txt", config.ConversionChain[1].Dict.Dicts[1].File)
}
