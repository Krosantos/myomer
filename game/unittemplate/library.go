package unittemplate

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
)

// Library -- A library of all unit templates by ID
var Library map[string]Template

var expansions []string = []string{
	"01_Core",
}

func init() {
	Library = make(map[string]Template)
	fs, err := filepath.Glob("game/unittemplate/**/**.json")
	if err != nil {
		panic("Error loading unit templates")
	}
	for _, f := range fs {
		chunk := make(map[string]Template)
		raw, err := ioutil.ReadFile(f)
		if err != nil {
			panic("Error loading unit templates")
		}
		err = json.Unmarshal(raw, &chunk)
		if err != nil {
			panic("Error loading unit templates")
		}
		for id, t := range chunk {
			Library[id] = t
		}
	}
}
