package main

import (
	// "context"
	"flag"
	"log"
	"net/http"
	"os"
	"time"
)

type application struct {
	InfoLog         *log.Logger
	ErrorLog        *log.Logger
	Url             string
	Visited         map[string]bool
	Ignore          map[string]bool
	Maxlimit        int
	TotalLinkPerUrl int
	Client          *http.Client
}

type ListOfAnchor struct {
	List []string
}

func main() {
	url := flag.String("url", "https://www.github.com/iamgak", "URL to scrape")

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	timeout := flag.Int("timeout", 10, "Timeout in seconds for the scraping process")
	maxlimit := flag.Int("maxlimit", 10, "Maximum no. of data scrape")
	totalLinkPerUrl := flag.Int("totalLinkPerUrl", 10, "Maximum no. of data send to queue")
	fileName := flag.String("filename", filename(), "Filename for Csv file")
	flag.Parse()
	client := &http.Client{
		Timeout: time.Duration(*timeout) * time.Second,
	}

	app := &application{
		InfoLog:         infoLog,
		ErrorLog:        errorLog,
		Visited:         make(map[string]bool),
		Ignore:          make(map[string]bool),
		Maxlimit:        *maxlimit,
		TotalLinkPerUrl: *totalLinkPerUrl,
		Client:          client,
	}

	links := app.dfs(*url)
	if len(links) != 0 {
		app.writeCSV(links, *fileName)
		app.InfoLog.Print("Successfully, Saved csv file")
	} else {
		app.InfoLog.Print("!written Links")
		app.InfoLog.Print(app.Ignore)
		app.Visited = make(map[string]bool)
		app.Ignore = make(map[string]bool)
		app.InfoLog.Print("No Data Saved")
	}
}
