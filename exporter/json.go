package exporter

import (
	"encoding/json"
	"os"

	"github.com/hannesharnisch/kmm-localization/core"
)

func ExportAsJson(filePath *string, localization *core.Localization) {
	json, err := json.MarshalIndent(localization, "", "  ")
	core.Check(err)
	os.WriteFile(*filePath, json, 0644)
}
