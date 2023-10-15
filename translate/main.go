package translate

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/fatih/color"
)

type Translate map[string]interface{}

func Load(language string) Translate {
	if language == "" {
		color.Yellow("Language param invalid or empty: %s", language)
		color.Yellow("Loading default language: en-US")
		language = "en-US"
	}

	jsonPath := filepath.Join("translate", language+".json")
	file, err := os.Open(jsonPath)

	if err != nil {
		language = "en-Us"
		color.Yellow("Error loading specified language %s :: %v\n", language, err)
		color.Yellow("Loading default language: en-US")
	}

	defer file.Close()

	var data map[string]interface{}
	decoder := json.NewDecoder(file)

	if err := decoder.Decode(&data); err != nil {
		color.Red("Error while decoding JSON:", err)
		panic(err)
	}

	return data
}
