package main

import (
	"encoding/csv"
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

func (app *application) dfs(url string) []string {
	queue := &ListOfAnchor{
		List: []string{url},
	}

	CsvOutput := []string{}
	for len(queue.List) > 0 && len(app.Visited) < app.Maxlimit {
		pop := queue.List[0]
		queue.List = queue.List[1:]
		if has := app.Visited[pop]; pop != "" && !has && !app.Ignore[pop] {
			data := app.scrape(pop)
			if data == nil || len(data.List) == 0 {
				app.Ignore[pop] = true
			} else {
				queue.List = append(queue.List, data.List...)
				app.Visited[pop] = true
				CsvOutput = append(CsvOutput, pop)
			}
		}
	}

	return CsvOutput
}

func (app *application) scrape(url string) *ListOfAnchor {
	links := &ListOfAnchor{
		List: make([]string, 0),
	}

	re, err := regexp.Compile(`^(?:(https?:\/\/)([^:/$]{1,})(?::(\d{1,}))?(?:($|\/(?:[^?#]{0,}))?((?:\?(?:[^#]{1,}))?)?(?:(#(?:.*)?)?|$)))$`)
	if err != nil {
		app.InfoLog.Print("Invalid url")
		return links
	}

	valid := re.MatchString(url)
	if !valid {
		return links
	}

	app.InfoLog.Printf("Scraping URL: %s\n", url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		app.InfoLog.Print(err)
		return nil
	}

	resp, err := app.Client.Do(req)
	if err != nil {
		if os.IsTimeout(err) {
			app.InfoLog.Println("Error: scraping timed out")
		} else {
			app.InfoLog.Print(err)
		}

		return nil
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		app.InfoLog.Printf("Error: status code %d\n", resp.StatusCode)
		return nil
	}

	doc, err := html.Parse(resp.Body)

	if err != nil {
		app.InfoLog.Print(err)
		return nil
	}

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					link := attr.Val
					if strings.HasPrefix(link, "http") && !app.Visited[link] && len(links.List) < app.TotalLinkPerUrl {
						links.List = append(links.List, link)
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	return links
}

func (app *application) writeCSV(links []string, outputFile string) {
	file, err := os.OpenFile("./csv/"+outputFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		app.ErrorLog.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, link := range links {
		if err := writer.Write([]string{link}); err != nil {
			app.ErrorLog.Fatal(err)
		}
	}

	app.InfoLog.Printf("Successfully written %d links to %s\n", len(links), outputFile)
}

func filename() string {
	return fmt.Sprintf("links%s.csv", time.Now().Format("20060102_150405"))
}
