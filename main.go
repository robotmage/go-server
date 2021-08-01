package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const version = "0.1"

func handler(w http.ResponseWriter, r *http.Request) {
	hostName, _ := os.Hostname()
	fmt.Fprintf(w, "Version %s Hi There from [%s], I love %s!\n",version, hostName, r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
