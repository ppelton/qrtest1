package web

import (
	"io"
	"log"
	"net/http"

	"qrtest1.peltonium.net/types"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, `
	<html>
		<title>Root page</title>
		<body>This is the root page</body>
	</html>`)
}

func StartServer(contentProviders []types.ContentProvider) {
	http.HandleFunc("/", rootHandler)
	for _, cp := range contentProviders {
		http.HandleFunc(cp.Path, func(w http.ResponseWriter, r *http.Request) {
			var bytes, error = cp.Provide()
			if error == nil {
				io.WriteString(w, string(bytes))
			} else {
				io.WriteString(w, "Error: "+error.Error())
			}
		})
	}
	log.Fatal(http.ListenAndServe(":8080", nil))
}
