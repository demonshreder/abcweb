package cmd

import (
	"text/template"

	"github.com/nullbio/abcweb/strmangle"
)

var templateFuncs = template.FuncMap{
	"randString": strmangle.RandString,
	"envAppName": strmangle.EnvAppName,
	"dbAppName":  strmangle.DBAppName,
}
