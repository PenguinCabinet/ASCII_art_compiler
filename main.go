package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func cli_build(out_fname string, type_data string) error {
	fname := out_fname

	setting_file := new_setting_file_t()
	temp_setting := (file_load("./setting.json"))
	if err := json.Unmarshal(temp_setting, &setting_file); err != nil {
		log.Fatal(err)
	}

	source := string(file_load("./main.aasc"))

	switch type_data {
	case "image":
		if fname == "____default" {
			fname = "output.png"
		}
		image_bytes := image_build(setting_file, source)
		make_file(fname, image_bytes)
	case "html":
		if fname == "____default" {
			fname = "output.html"
		}
		bytes := html_build(setting_file, source)
		make_file(fname, bytes)
	case "pdf":
		if fname == "____default" {
			fname = "output.pdf"
		}
		bytes := pdf_build(setting_file, source)
		make_file(fname, bytes)

	}

	return nil
}

func main() {

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "new",
				Aliases: []string{"n"},
				Usage:   "New project",
				Action: func(c *cli.Context) error {
					//fmt.Println("added task: ", c.Args().First())
					new_project()
					return nil
				},
			},
			{
				Name:    "build",
				Aliases: []string{"build"},
				Usage:   "Build the project ASCII art",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "o",
						Value: "____default",
						Usage: "The output file name.",
					},
					&cli.StringFlag{
						Name:  "type",
						Value: "image",
						Usage: "The type of output file.",
					},
				},
				Action: func(c *cli.Context) error {

					fname := c.String("o")
					cli_build(fname, c.String("type"))

					//image_bytes := image_build(setting_file, source)
					//make_file("output.png", image_bytes)
					//TODO switch
					//fmt.Println("completed task: ", c.Args().First())
					return nil
				},
			},
			{
				Name:    "real-time-preview",
				Aliases: []string{"rtp"},
				Usage:   "Build the project ASCII art",
				Flags:   []cli.Flag{},
				Action: func(c *cli.Context) error {

					real_preview_server()
					//image_bytes := image_build(setting_file, source)
					//make_file("output.png", image_bytes)
					//TODO switch
					//fmt.Println("completed task: ", c.Args().First())
					return nil
				},
			},
		},
	}

	app.Run(os.Args)

}
