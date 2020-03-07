package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"text/template"
)

//Tmpl template file struct
type Tmpl struct {
	IPInfo        string
	RequestHeader *http.Header
}

const (
	version     = "1.0.0 (2020-03-08)"
	portDefault = 9999
	tmplDefault = "./web/template/whereismyip.tmpl"
)

var (
	port        int
	tmplDir     string
	showVersion bool
	//LogInfo log level-INFO
	LogInfo *log.Logger
	//LogWarn log level-WARN
	LogWarn *log.Logger
	//LogError log level-ERROR
	LogError *log.Logger
)

func init() {

	flag.BoolVar(&showVersion, "v", false, "version")
	flag.StringVar(&tmplDir, "tmpl", tmplDefault, "tmpl file")
	flag.IntVar(&port, "p", portDefault, "http server port")
	flag.Parse()

	logFile, err := os.OpenFile("whereismyip.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Open log file fail:", err)
	}

	LogInfo = log.New(logFile, "INFO:", log.LstdFlags|log.Ltime|log.Llongfile)
	LogWarn = log.New(logFile, "WARN:", log.LstdFlags|log.Ltime|log.Llongfile)
	LogError = log.New(logFile, "ERROR:", log.LstdFlags|log.Ltime|log.Llongfile)
}

func main() {
	switch {
	case showVersion:
		fmt.Println("version:" + version)
		return
	case port <= 0:
		LogInfo.Println("use default port:", portDefault)
		port = portDefault
	}

	LogInfo.Println("Listen port:", port)

	http.HandleFunc("/", indexHandler)
	LogError.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	var ip string

	t := &Tmpl{
		IPInfo:        "Can't identify your IP address.",
		RequestHeader: &req.Header,
	}

	if len(req.Header.Get("X-Forwarded-For")) > 0 {
		ip = net.ParseIP(strings.Split(req.Header.Get("X-Forwarded-For"), ",")[0]).String()
		if len(ip) > 0 {
			t.IPInfo = "Your IP address:" + ip
		}
	}

	tmpl, err := template.ParseFiles(tmplDir)
	if err != nil {
		LogError.Fatalf("Parse: %v", err)
	}

	err = tmpl.Execute(w, t)
	if err != nil {
		LogError.Fatalf("Parse: %v", err)
	}

}
