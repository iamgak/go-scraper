package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	// node "webscraper.iamgak.com/models"
)

type application struct {
	InfoLog  *log.Logger
	ErrorLog *log.Logger
	Url      string
	Visited  *Visited
}

type Visited struct {
	List map[string]bool
}

type ListOfAnchor struct {
	Url  string
	List []string
}

// function
func main() {
	mux := http.NewServeMux()
	url := flag.String("url", "https://www.freecodecamp.com", "Given Url link")
	addr := flag.String("addr", ":4000", "Port Number")
	flag.Parse()
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		Url:      *url,
		InfoLog:  infoLog,
		ErrorLog: errorLog,
		Visited: &Visited{
			List: make(map[string]bool),
		},
	}
	mux.HandleFunc("/", app.home)

	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
