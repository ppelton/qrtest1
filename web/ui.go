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

func corsHandler(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Print("preflight detected: ", r.Header)
		w.Header().Add("Connection", "keep-alive")
		w.Header().Add("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Add("Access-Control-Allow-Methods", "POST, OPTIONS, GET, DELETE, PUT")
		w.Header().Add("Access-Control-Allow-Headers", "content-type")
		w.Header().Add("Access-Control-Max-Age", "86400")

		// continue with my method
		h(w, r)
	}
}

func StartServer(contentProviders []types.ContentProvider) {
	http.HandleFunc("/", rootHandler)
	for _, cp := range contentProviders {
		http.HandleFunc(cp.Path, corsHandler(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("content/type", "image/png")
			var bytes, error = cp.Provide()
			if error == nil {
				io.WriteString(w, string(bytes))
			} else {
				io.WriteString(w, "Error: "+error.Error())
			}
		}))
	}
	log.Fatal(http.ListenAndServe(":8080", nil))
}
