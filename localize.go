package main

import (
	"flag"

	"github.com/hannesharnisch/kmm-localization/exporter"
	"github.com/hannesharnisch/kmm-localization/parser"
)

func main() {
	excelgen := flag.Bool("excelgen", false, "tells the script to generate from excel")
	ressourcePath := flag.String("ressource", "./ressources/i18n", "Set the ressource folder path")
	kotlinLocalizationPath := flag.String("kotlin", "./src/commonMain/kotlin/i18n", "Set the kotlin localization folder path")
	kotlinPackageName := flag.String("kotlin-package", "i18n", "Set the kotlin localization package name")
	defaultLang := flag.String("base-lang", "en", "Sets the default language that is used")
	genFromJson := flag.Bool("rev", false, "Reverses the generation process and generates an excel file from intermediate representation")
	flag.Parse()

	excelFilePath := *ressourcePath + "/Localization.xlsx"
	jsonPath := *ressourcePath + "/Localization.json"

	if *excelgen {
		if *genFromJson {
			localization := parser.LoadfromJson(&jsonPath)
			exporter.ExportAsExcel(&excelFilePath, localization)
		} else {
			localization := parser.LoadFromExcel(&excelFilePath, *defaultLang)
			exporter.ExportAsJson(&jsonPath, localization)
		}
	}
	localization := parser.LoadfromJson(&jsonPath)
	exporter.ExportAsKMMRessources(ressourcePath, localization)
	exporter.ExportAsKotlinObjects(kotlinLocalizationPath, *kotlinPackageName, localization)
}
