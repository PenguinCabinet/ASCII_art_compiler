package main

import (
	"bytes"
	_ "embed"
	"encoding/base64"
	"html/template"
	"io/ioutil"
	"strings"
)

//go:embed template.html
var html_template_text []byte

func html_build(setting setting_file_t, source string) []byte {
	ftBinary, err := ioutil.ReadFile("font/font.ttf")

	type html_args_t struct {
		Font string
		Data []string
	}

	args := html_args_t{}
	args.Font = base64.StdEncoding.EncodeToString(ftBinary)
	args.Data = strings.Split(source, "\n")

	html_template, err := template.New("output").Parse(string(html_template_text))

	if err != nil {
		panic(err)
	}

	var A *bytes.Buffer = new(bytes.Buffer)

	err = html_template.Execute(A, args)
	if err != nil {
		panic(err)
	}

	A_buf := A.Bytes()

	return A_buf
}
