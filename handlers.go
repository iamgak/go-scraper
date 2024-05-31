package main

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	visited := &Visited{
		List: make(map[string]bool),
	}

	stackData := app.anchorList(app.Url)
	i := 0
	// i = 0
	for len(stackData.List) > 0 {
		pop := stackData.List[0]
		stackData.List = stackData.List[1:]
		if _, has := visited.List[pop]; !has && i < 100 {
			visited.List[pop] = true
			log.Print(pop)
			stackData.List = append(stackData.List, app.anchorList(pop).List...)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(visited)
}

func (app *application) anchorList(main_url string) ListOfAnchor {
	links := ListOfAnchor{
		Url:  main_url,
		List: make([]string, 0),
	}

	re, err := regexp.Compile(`^(?:(https?:\/\/)([^:/$]{1,})(?::(\d{1,}))?(?:($|\/(?:[^?#]{0,}))?((?:\?(?:[^#]{1,}))?)?(?:(#(?:.*)?)?|$)))$`)
	if err != nil {
		log.Fatal(err)
	}

	valid := re.MatchString(main_url)
	if !valid {
		return links
	}

	response, err := http.Get(main_url)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists {
			valid := re.MatchString(href)
			if valid {
				links.List = append(links.List, href)
			}
		}
	})

	return links
}
