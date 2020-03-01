package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(":9999", nil))
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		panic(err)
	}
	var ip string

	if len(req.Header.Get("X-Forwarded-For")) > 0 {
		ip = net.ParseIP(strings.Split(req.Header.Get("X-Forwarded-For"), ",")[0]).String()
	}

	if len(ip) > 0 {
		fmt.Fprintf(w, "Your IP address: %s\n", ip)
	} else {
		fmt.Fprintf(w, "Can't identify your IP address.\n")
	}

}
