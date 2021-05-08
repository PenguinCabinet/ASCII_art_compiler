package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func make_file(fname string, text []byte) {
	f, _ := os.Create(fname)
	f.Write(text)
	f.Close()
}

func file_load(fileName string) []byte {

	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	return (bytes)
}

func new_project() {
	make_file("./main.aasc", []byte(""))
	text, _ := json.Marshal(new_setting_file_t())
	make_file("./setting.json", (text))
	os.Mkdir("./font", 0777)
}
