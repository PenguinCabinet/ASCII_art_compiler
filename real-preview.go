package main

import (
	_ "embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/websocket"
)

var real_preview_one_html map[string]string = map[string]string{
	"image": `<img border="0" src="./temp/out.image" width="80%" >`,
	"html":  `<iframe src="./temp/out.html" width="100%" height="100%"></iframe>`,
	"pdf":   `<embed src="./temp/out.pdf" type="application/pdf" width="100%" height="100%">`,
}

//go:embed real_time_preview_index.html
var real_preview_html_index string

//go:embed real_time_preview.html
var real_preview_html string

func real_preview_server() {
	datach := make(chan bool, 1)
	for k, e := range real_preview_one_html {
		cli_build("./temp/"+"out."+k, k)
		temp_k := k
		temp_e := e
		fname := "./temp/" + "out." + k
		http.HandleFunc("/"+k, func(w http.ResponseWriter, r *http.Request) {
			one_real_preview_server(w, r, temp_e, "main.aasc", fname, temp_k, &datach)
		})
	}

	http.Handle("/__s", websocket.Handler(func(ws *websocket.Conn) { websocket_func(ws, &datach) }))

	fs := http.FileServer(http.Dir("./temp"))
	http.Handle("/temp/", http.StripPrefix("/temp/", fs))

	http.HandleFunc("/", index_real_preview_server)

	fmt.Printf("http://localhost:8994\n")
	http.ListenAndServe(":8994", nil)

}
func index_real_preview_server(w http.ResponseWriter, r *http.Request) {
	html_template, _ := template.New("rpt_index").Parse(string(real_preview_html_index))
	html_template.Execute(w, real_preview_one_html)
}

func one_real_preview_server(w http.ResponseWriter, r *http.Request, text, src_fname, out_fname, file_type string, datach *chan bool) {
	cli_build(out_fname, file_type)
	go func() {
		f_info_l, err1 := os.Stat(src_fname)
		if err1 != nil {
			log.Println(err1)
		}
		fmod_l := f_info_l.ModTime()
		dsec_l := fmod_l.Unix()

		for {
			f_info_n, err2 := os.Stat(src_fname)
			if err2 != nil {
				continue
			}
			if f_info_n == nil {
				continue
			}
			fmod_n := f_info_n.ModTime()

			dsec_n := fmod_n.Unix()

			if dsec_l < dsec_n {
				cli_build(out_fname, file_type)
				dsec_l = dsec_n
				*datach <- true
			}
		}

	}()
	html := fmt.Sprintf(real_preview_html, text)
	fmt.Fprintf(w, html)
}

func websocket_func(ws *websocket.Conn, datach *chan bool) {
	for data := range *datach {
		err := websocket.JSON.Send(ws, data)
		if err != nil {
			//log.Printf("error sending data: %v\n", err)
			return
		}
	}

}
