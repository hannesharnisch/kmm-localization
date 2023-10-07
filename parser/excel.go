package parser

import (
	"reflect"

	"github.com/hannesharnisch/kmm-localization/core"
	"github.com/xuri/excelize/v2"
)

func LoadFromExcel(filePath *string, defaultLang string) *core.Localization {
	f, error := excelize.OpenFile(*filePath)
	core.Check(error)

	sheets := f.GetSheetList()

	languages := getLanguages(f, 2)
	localization := core.Localization{
		Languages: []core.Language{},
		Scopes:    []core.LocalizationScope{},
	}
	for _, x := range languages {
		localization.Languages = append(localization.Languages, core.Language{
			Code:      x,
			IsDefault: x == defaultLang,
		})
	}
	for _, sheet := range sheets {
		rows, error := f.GetRows(sheet)
		core.Check(error)
		scope := core.LocalizationScope{
			Name: sheet,
		}
		for _, row := range rows[1:] {
			entry := core.LocalizationEntry{}
			for x, cell := range row {
				if x == 0 {
					entry.Description = cell
				} else if x == 1 {
					entry.Key = cell
				} else {
					entry.Translations = append(entry.Translations, core.Translation{
						Language: localization.Languages[x-2].Code,
						Value:    cell,
					})
				}
			}
			scope.Entries = append(scope.Entries, entry)
		}
		localization.Scopes = append(localization.Scopes, scope)
	}
	f.Close()
	return &localization
}

func getLanguages(f *excelize.File, offset int) []string {
	languages := map[string]bool{}
	sheets := f.GetSheetList()
	for _, sheet := range sheets {
		rows, _ := f.GetRows(sheet)
		for _, lang := range rows[0][offset:] {
			languages[lang] = true
		}
	}
	keys := reflect.ValueOf(languages).MapKeys()
	strkeys := make([]string, len(keys))
	for i := 0; i < len(keys); i++ {
		strkeys[i] = keys[i].String()
	}
	return strkeys
}
