package main

import (
	"SpecialWorkers/models"
	"fmt"
	"github.com/gocolly/colly"
	"log"
)

func extractMangaChapters(url string, mangaDefaultURL string) ([]models.MangaChapter, error) {
	c := colly.NewCollector()
	var chapterList []models.MangaChapter
	c.OnHTML("#chapters-list a", func(c *colly.HTMLElement) {
		chapter := models.MangaChapter{
			ChapterUrl:    fmt.Sprintf("%s%s", mangaDefaultURL, c.Attr("href")),
			ChapterNumber: c.Text,
			ChapterName:   c.Attr("title"),
		}
		chapterList = append(chapterList, chapter)
	})
	err := c.Visit(url)
	if err != nil {
		return nil, err
	}
	return chapterList, nil
}
func worker(data models.WorkerInformation) {
	c := colly.NewCollector()
	var mangaList []models.MangaInformation
	c.OnHTML(data.MangaSearchPattern, func(c *colly.HTMLElement) {
		manga := models.MangaInformation{
			MangaUrl:      fmt.Sprintf("%s%s", data.BaseWebsiteUrl, c.Attr("href")),
			MangaImageUrl: c.ChildAttr("img", "src"),
			MangaName:     c.ChildText(".flex div:first-child"),
		}
		chapters, err := extractMangaChapters(manga.MangaUrl, data.BaseWebsiteUrl)
		if err != nil {
			log.Panic(err)
		}
		manga.ChapterList = chapters
		mangaList = append(mangaList, manga)
	})
	c.Visit(fmt.Sprintf("%s%s", data.WebsiteSearchUrl, data.MangaName))

	//TODO: SEND DATA TO DB
	fmt.Printf("%+v\n", mangaList)
}
func main() {
	var testDemonScans = models.WorkerInformation{
		WebsiteSearchUrl:     "https://demonicscans.org/search.php?manga=",
		BaseWebsiteUrl:       "https://demonicscans.org",
		MangaName:            "Solo",
		MangaId:              0,
		ChapterSearchPattern: "",
		MangaSearchPattern:   "a",
		ChapterList:          nil,
	}
	worker(testDemonScans)
}
