package main

import (
	"fmt"
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

var real_preview_html string = `
<html>
<head>

      <title>Real time preview</title>

      <script type="text/javascript">
      var sock = null;
      var myData = "";
      function update() {
	      location.reload();
      };
      window.onload = function() {
          sock = new WebSocket("ws://"+location.host+"/__s");
          sock.onmessage = function(event) {
              update();
          };
      };
      </script>
</head>
<body>
</body>

%s

</html>
`

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
	http.ListenAndServe(":8888", nil)
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
