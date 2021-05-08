package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

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
						Name:  "type",
						Value: "image",
						Usage: "The type of output file.",
					},
				},
				Action: func(c *cli.Context) error {
					setting_file := new_setting_file_t()
					temp_setting := (file_load("./setting.json"))
					if err := json.Unmarshal(temp_setting, &setting_file); err != nil {
						log.Fatal(err)
					}

					source := string(file_load("./main.aasc"))

					switch c.String("type") {
					case "image":
						image_bytes := image_build(setting_file, source)
						make_file("output.png", image_bytes)
					case "html":
						bytes := html_build(setting_file, source)
						make_file("output.html", bytes)

					}
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
