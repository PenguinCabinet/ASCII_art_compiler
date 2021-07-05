package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
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
		log.Fatalln(err)
	}

	return (bytes)
}

func new_project() {
	make_file("./main.aasc", []byte(""))
	text, _ := json.Marshal(new_setting_file_t())
	make_file("./setting.json", (text))
	os.Mkdir("./font", 0777)

	/*
		my_out_path, err := os.Executable()

		if err != nil {
			log.Fatalln(err)
		}

		my_dir_path := filepath.Dir(my_out_path)

		make_file("./font/font.ttf", file_load(filepath.Join(my_dir_path, "font", "font.ttf")))
	*/
}
