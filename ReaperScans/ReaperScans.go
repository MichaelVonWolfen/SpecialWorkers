package main

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"io"
	"log"
	"net/http"
)

func getMangaChapters(url string) ([]MangaChapter, error) {
	//resp, err := http.Get(url)
	//if err != nil {
	//	return nil, err
	//}
	c := colly.NewCollector()
	c.OnError(func(r *colly.Response, err error) {
		log.Println("Error:", r.Request.URL, err)
	})
	c.OnHTML("div[id$='chapters_list'] a", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})
	err := c.Visit(url)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
func worker(data WorkerInformation) {
	getMangaChapters(data.mangaUrl)
	return
	resp, err := http.Get(data.websiteSearchUrl + data.mangaName)
	if err != nil {
		log.Fatal(err)
	}
	if resp != nil {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		var response ReaperResponseStruct
		err = json.Unmarshal(body, &response)
		if err != nil {
			log.Fatalln(err)
		}
		for _, data := range response.Data {
			fmt.Println(data.Title)
		}
		err = resp.Body.Close()
		if err != nil {
			return
		}
	} else {
		fmt.Println("BAD")
	}
}

func main() {
	var testReaper = WorkerInformation{
		websiteSearchUrl:     "https://api.reaperscans.com/query?adult=true&query_string=",
		mangaName:            "Solo",
		mangaId:              0,
		mangaUrl:             "https://reaperscans.com/series/solo-leveling-ragnarok",
		websiteSearchPattern: "",
		mangaSearchPattern:   "",
		chapterList:          nil,
	}
	worker(testReaper)
}
