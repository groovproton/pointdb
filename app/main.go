package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	DDMMYYYYhhmmss = "2006-01-02 15:04:05"
)

func HelloServer(w http.ResponseWriter, r *http.Request) {
	tm:= time.Now().UTC()

	fmt.Fprintf(w, "<b>Start work PointDB. Url: %s! time:%q, </b>", r.URL.Path[1:], tm.Format(DDMMYYYYhhmmss))
}


func main() {
	port:= ":8080"
	println("Start server "+ port)

	http.HandleFunc("/", HelloServer)

	http.ListenAndServe(port, nil)
}
