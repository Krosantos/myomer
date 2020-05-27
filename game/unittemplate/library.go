package unittemplate

import (
	"path/filepath"
)

// Library -- A library of all unit templates by ID
var Library map[string]Template

var expansions []string = []string{
	"01_Core",
}

func init() {
	fs, err := filepath.Glob("**.json")
	if err != nil {
		println(err.Error())
	}
	for _, f := range fs {
		println(f)
	}
}
