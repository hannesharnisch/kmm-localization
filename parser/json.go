package parser

import (
	"encoding/json"
	"os"

	"github.com/hannesharnisch/kmm-localization/core"
)

func LoadfromJson(filePath *string) *core.Localization {
	jsonData, err := os.ReadFile(*filePath)
	core.Check(err)
	var localization core.Localization
	error := json.Unmarshal(jsonData, &localization)
	core.Check(error)
	return &localization
}
