package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"

	"github.com/beckxie/whatismyip/internal/api"
	"github.com/beckxie/whatismyip/internal/web"
)

var (
	version   = "dev"
	commit    = "none"
	buildDate = "unknown"
)

const portDefault = 9999

func main() {
	var showVersion bool
	var port int
	flag.BoolVar(&showVersion, "v", false, "version")
	flag.IntVar(&port, "p", portDefault, "http server port")
	flag.Parse()

	if showVersion {
		fmt.Printf("whatismyip version: %s\n", version)
		fmt.Printf("commit: %s\n", commit)
		fmt.Printf("build date: %s\n", buildDate)
		return
	}

	if port <= 0 {
		slog.Info("invalid port, using default", "port", portDefault)
		port = portDefault
	}

	// Initialize handlers
	webHandler, err := web.NewHandler("./web/template/whatismyip.tmpl")
	if err != nil {
		slog.Error("failed to initialize web handler", "error", err)
		os.Exit(1)
	}

	apiHandler := api.NewHandler()

	slog.Info("starting server", "port", port, "version", version)

	// Static files
	fs := http.FileServer(http.Dir("./web/static"))
	http.Handle("/web/static/", http.StripPrefix("/web/static/", fs))

	// Routes
	http.HandleFunc("/", webHandler.Index)
	http.HandleFunc("/api/ip", apiHandler.GetIP)

	addr := ":" + strconv.Itoa(port)
	if err := http.ListenAndServe(addr, nil); err != nil {
		slog.Error("server failed", "error", err)
		os.Exit(1)
	}
}
