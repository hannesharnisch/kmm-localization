package exporter

import (
	"strconv"

	"github.com/hannesharnisch/kmm-localization/core"
	"github.com/xuri/excelize/v2"
)

func ExportAsExcel(filePath *string, localization *core.Localization) {
	letters := []rune{}
	for ch := 'A'; ch <= 'Z'; ch++ {
		letters = append(letters, ch)
	}
	f := excelize.NewFile()
	for _, scope := range localization.Scopes {
		f.NewSheet(scope.Name)
		f.SetCellValue(scope.Name, "A1", "Beschreibung")
		f.SetCellValue(scope.Name, "B1", "Key")
		for i, lang := range localization.Languages {
			f.SetCellValue(scope.Name, string(letters[2:][i])+"1", lang.Code)
		}
	}
	f.DeleteSheet("Sheet1")
	for _, scope := range localization.Scopes {
		i := 2
		for _, value := range scope.Entries {
			core.Check(f.SetCellValue(scope.Name, "B"+strconv.Itoa(i), value.Key))
			core.Check(f.SetCellValue(scope.Name, "A"+strconv.Itoa(i), value.Description))
			a := 2
			for _, trans := range value.Translations {
				core.Check(f.SetCellValue(scope.Name, string(letters[a])+strconv.Itoa(i), trans.Value))
				a += 1
			}
			i += 1
		}
	}
	error := f.SaveAs(*filePath)
	core.Check(error)
	f.Close()
}
