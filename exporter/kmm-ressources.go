package exporter

import (
	"encoding/json"
	"os"

	"github.com/flosch/pongo2/v6"
	"github.com/hannesharnisch/kmm-localization/core"
)

const kotlinTemplate = `package {{ package_name }}

sealed class Strings(val scope: String) {
	val defaultLanguage = "{{ default_language }}"
	{% for name, keys in scopes %}
	data object {{ name | capfirst }}: Strings(scope = "{{ name }}") {
		{% for key in keys %}
		val {{ key }} by I18n(){% endfor %}
	}{% endfor %}
}
`

func ExportAsKMMRessources(filesPath *string, localization *core.Localization) {
	for _, lang := range localization.Languages {
		for _, scope := range localization.Scopes {
			translations := make(map[string]string)
			for _, entry := range scope.Entries {
				for _, trans := range entry.Translations {
					if trans.Language == lang.Code {
						translations[entry.Key] = trans.Value
					}
				}
			}
			json, err := json.MarshalIndent(translations, "", "  ")
			core.Check(err)
			os.MkdirAll(*filesPath+"/"+lang.Code, 0777)
			core.Check(os.WriteFile(*filesPath+"/"+lang.Code+"/"+scope.Name+".json", json, 0777))
		}
	}
}

func ExportAsKotlinObjects(filesPath *string, package_name string, localization *core.Localization) {
	var template = pongo2.Must(pongo2.FromString(kotlinTemplate))
	scopes := make(map[string][]string)
	default_language := ""
	for _, lang := range localization.Languages {
		if lang.IsDefault {
			default_language = lang.Code
		}
	}
	for _, scope := range localization.Scopes {
		keys := []string{}
		for _, entry := range scope.Entries {
			keys = append(keys, entry.Key)
		}
		scopes[scope.Name] = keys
	}
	res, error := template.Execute(pongo2.Context{"package_name": package_name, "default_language": default_language, "scopes": scopes})
	core.Check(error)
	core.Check(os.WriteFile(*filesPath+"/Strings.kt", []byte(res), 0777))
}
