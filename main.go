package main

import (
	"encoding/json"
	"fmt"
	"os"
	"text/template"
)

func main() {
	var (
		templateFile string
		vars         = make(map[string]interface{})
		err          error
		tpl          *template.Template
		decoder      *json.Decoder
	)

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <template>\nVariables should be sent as JSON from STDIN\n", os.Args[0])
		os.Exit(1)
	}

	templateFile = os.Args[1]

	if _, err = os.Stat(templateFile); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "%s not found\n", templateFile)
		os.Exit(1)
	}

	tpl = template.Must(template.ParseGlob(templateFile))

	decoder = json.NewDecoder(os.Stdin)

	if err = decoder.Decode(&vars); err != nil {
		fmt.Fprintf(os.Stderr, "error reading parsing stdin: %s", err)
		os.Exit(1)
	}

	err = tpl.Execute(os.Stdout, vars)

	if err != nil {
		fmt.Fprintf(os.Stderr, "error rendering template %v: %v", templateFile, err)
		os.Exit(1)
	}
}