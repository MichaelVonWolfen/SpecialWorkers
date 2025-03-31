package main

import (
	"SpecialWorkers/models"
	"fmt"
	"github.com/gocolly/colly"
	"log"
)

func worker(data models.WorkerInformation) ([]models.MangaInformation, error) {
	//Access URL to data
	var url = fmt.Sprintf("%s%s", data.WebsiteSearchUrl, data.MangaName)
	var mangaList []models.MangaInformation
	c := colly.NewCollector()
	c.OnHTML(data.MangaSearchPattern, func(r *colly.HTMLElement) {
		manga := models.MangaInformation{
			MangaUrl:      fmt.Sprintf("%s/%s", data.BaseWebsiteUrl, r.Attr("href")),
			Status:        r.ChildText("span.status"),
			MangaName:     r.ChildText("span.block"),
			MangaImageUrl: r.ChildAttr("img", "src"),
		}
		mangaList = append(mangaList, manga)
	})
	err := c.Visit(url)
	if err != nil {
		return nil, err
	}
	return mangaList, nil
}

func extractChapters(manga models.MangaInformation) (models.MangaInformation, error) {
	c := colly.NewCollector()

	var chapterList []models.MangaChapter
	c.OnHTML("div.border a", func(e *colly.HTMLElement) {
		chapter := models.MangaChapter{
			ChapterUrl:    e.Attr("href"),
			ChapterNumber: e.ChildText("h3.flex"),
			ChapterName:   e.ChildText("h3 span"),
		}
		//fmt.Printf("%+v\n", chapter)
		chapterList = append(chapterList, chapter)
	})
	err := c.Visit(manga.MangaUrl)
	if err != nil {
		return models.MangaInformation{}, err
	}
	manga.ChapterList = chapterList
	return manga, nil
}

func main() {
	var testAsura = models.WorkerInformation{
		WebsiteSearchUrl:     "https://asuracomic.net/series?page=1&name=",
		BaseWebsiteUrl:       "https://asuracomic.net",
		MangaName:            "Solo",
		MangaId:              0,
		ChapterSearchPattern: "",
		MangaSearchPattern:   ".grid a",
		ChapterList:          nil,
	}
	mangaList, err := worker(testAsura)
	if err != nil {
		return
	}
	for _, manga := range mangaList {
		MangaWithChapter, err := extractChapters(manga)
		if err != nil {
			log.Fatal(err)
		}
		//mangaList = append(mangaList, MangaWithChapter)
		for i := range MangaWithChapter.ChapterList {
			MangaWithChapter.ChapterList[i].ChapterUrl = fmt.Sprintf("%s/series/%s", testAsura.BaseWebsiteUrl, MangaWithChapter.ChapterList[i].ChapterUrl)
		}
		//TODO: SEND RETRIEVED DATA TO THE DB???
		fmt.Printf("%+v\n", MangaWithChapter)
	}
}
