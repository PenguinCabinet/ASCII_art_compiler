package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
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

	if _, err := os.Stat("setting.json"); err == nil {
		log.Fatalln("The directory in which make a project must be empty.")
	}

	make_file("main.aasc", []byte(""))
	text, _ := json.Marshal(new_setting_file_t())
	make_file("setting.json", (text))
	os.Mkdir("font", 0777)

	my_out_path, err := os.Executable()

	if err != nil {
		log.Fatalln(err)
	}

	my_dir_path := filepath.Dir(my_out_path)

	make_file(filepath.Join("font", "font.ttf"), file_load(filepath.Join(my_dir_path, "default_font", "font.ttf")))
	make_file(filepath.Join("font", "IPA_Font_License_Agreement_v1.0.txt"), file_load(filepath.Join(my_dir_path, "default_font", "IPA_Font_License_Agreement_v1.0.txt")))
	make_file(filepath.Join("font", "Readme_ipam00303.txt"), file_load(filepath.Join(my_dir_path, "default_font", "Readme_ipam00303.txt")))
}
