package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const version = "0.2"

func responseString() string {
	hostName, _ := os.Hostname()
	return fmt.Sprintf("Version %s Hi There from [%s]\n", version, hostName)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, responseString())
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
