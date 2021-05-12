package main

import (
	"fmt"
	"net/http"
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
      };
      /*
      window.onload = function() {
          sock = new WebSocket("ws://"+location.host+"/__s");
          sock.onmessage = function(event) {
              update();
          };
      };
      */
      </script>
</head>
<body>
</body>

%s

</html>
`

func real_preview_server() {
	for k, e := range real_preview_one_html {
		cli_build("./temp/"+"out."+k, k)
		temp := e
		http.HandleFunc("/"+k, func(w http.ResponseWriter, r *http.Request) {
			one_real_preview_server(w, r, temp)
		})
	}
	fs := http.FileServer(http.Dir("./temp"))
	http.Handle("/temp/", http.StripPrefix("/temp/", fs))
	http.ListenAndServe(":8888", nil)
}

func one_real_preview_server(w http.ResponseWriter, r *http.Request, text string) {
	html := fmt.Sprintf(real_preview_html, text)
	fmt.Fprintf(w, html)
}
