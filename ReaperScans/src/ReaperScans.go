package main

import (
	"SpecialWorkers/models"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"runtime/debug"
	"slices"
	"strings"
)

func getMangaChapters(url string) ([]models.MangaChapter, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		debug.PrintStack()
		return nil, fmt.Errorf("url seems invalid")
	}
	if resp == nil {
		debug.PrintStack()
		log.Println("resp is null")
		return nil, fmt.Errorf("url seems invalid")
	}
	if resp.StatusCode != 200 {
		debug.PrintStack()
		log.Println("resp is not ok")
		return nil, fmt.Errorf("url seems invalid")
	}
	var mangaData ReaperMangaData
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		debug.PrintStack()
		return nil, err
	}
	err = json.Unmarshal(body, &mangaData)
	if err != nil {
		debug.PrintStack()
		return nil, err
	}
	if mangaData.Data == nil || len(mangaData.Data) == 0 {
		log.Printf("no chapters found for %s \n", url)
		debug.PrintStack()
		return nil, fmt.Errorf("url seems invalid")

	}
	var chapterList []models.MangaChapter
	for _, data := range mangaData.Data {
		var chapterNB = data.Index
		var erlSeriesName = strings.ToLower(data.Series.SeriesSlug)
		erlSeriesName = strings.ReplaceAll(erlSeriesName, " ", "-")
		var urlChapterName = strings.ReplaceAll(data.ChapterName, " ", "-")
		urlChapterName = strings.ToLower(urlChapterName)
		var chapter = models.MangaChapter{
			ChapterUrl:    fmt.Sprintf("https://reaperscans.com/series/%s/%s", erlSeriesName, urlChapterName),
			ChapterName:   data.ChapterName,
			ChapterNumber: chapterNB,
		}
		//fmt.Printf("%+v\n", chapter)
		chapterList = append(chapterList, chapter)
	}
	return chapterList, nil
}
func getMangaUrl(searchUrl string, mangaName string) ([]models.MangaInformation, error) {
	resp, err := http.Get(searchUrl + mangaName)
	if err != nil {
		log.Fatal(err)
	}
	if resp == nil {
		fmt.Println("BAD")
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	var mangaRawData ReaperResponseStruct

	err = json.Unmarshal(body, &mangaRawData)
	if err != nil {
		log.Fatalln(err)
	}
	var mangaList []models.MangaInformation
	for _, data := range mangaRawData.Data {
		manga := models.MangaInformation{
			MangaName:   data.Title,
			MangaUrl:    fmt.Sprintf("https://api.reaperscans.com/chapters/%d?page=1&perPage=1000&query=&order=desc", data.Id),
			ChapterList: nil,
		}
		mangaList = append(mangaList, manga)
	}
	err = resp.Body.Close()
	if err != nil {
		log.Println(err)
		//return nil, err
	}
	return mangaList, nil
}

func updateMangasAndChapters(firstMangaList []models.MangaInformation, secondMangaList []models.MangaInformation) []models.MangaInformation {

	for i := 0; i < len(firstMangaList); i++ {
		for j := 0; j < len(secondMangaList); j++ {
			if firstMangaList[i].MangaName == secondMangaList[j].MangaName {
				firstMangaList[i].ChapterList = secondMangaList[j].ChapterList
				firstMangaList[i].MangaId = int(math.Max(float64(firstMangaList[i].MangaId), float64(secondMangaList[j].MangaId)))
				firstMangaList[i].MangaUrl = secondMangaList[j].MangaUrl
			} else {
				contains := slices.ContainsFunc(firstMangaList, func(m models.MangaInformation) bool { return m.MangaName == secondMangaList[j].MangaName })
				if !contains {
					firstMangaList = append(firstMangaList, secondMangaList[j])
				}
			}
		}
	}

	return slices.DeleteFunc(firstMangaList, func(m models.MangaInformation) bool { return m.MangaUrl == "" })
}
func worker(data models.WorkerInformation) []models.MangaInformation {
	isUrlValid := true
	var mangaList []models.MangaInformation
	if data.MangaUrl != "" && len(data.MangaUrl) > 0 {
		chapterList, err := getMangaChapters(data.MangaUrl)
		if err != nil {
			if err.Error() == "url seems invalid" {
				log.Printf("url seems invalid for %s \n", data.MangaName)
				isUrlValid = false
			} else {
				log.Fatal(err)
			}
		}
		manga := models.MangaInformation{
			MangaName:   data.MangaName,
			MangaUrl:    data.MangaUrl,
			ChapterList: chapterList,
		}
		mangaList = updateMangasAndChapters([]models.MangaInformation{manga}, []models.MangaInformation{
			{
				MangaId:     data.MangaId,
				MangaName:   data.MangaName,
				MangaUrl:    data.MangaUrl,
				ChapterList: data.ChapterList,
			},
		})
	} else {
		isUrlValid = false
	}
	if !isUrlValid {
		mangas, err := getMangaUrl(data.WebsiteSearchUrl, data.MangaName)
		if err != nil {
			log.Panic(err)
		}
		fmt.Printf("%+v\n", mangas)

		log.Printf("Trying to read manga again after url update!")
		for i := 0; i < len(mangas); i++ {
			chapterList, err := getMangaChapters(mangas[i].MangaUrl)
			if err != nil {
				log.Panic(err)
			}
			mangas[i].ChapterList = chapterList
		}
		mangaList = updateMangasAndChapters([]models.MangaInformation{{
			MangaId:     data.MangaId,
			MangaName:   data.MangaName,
			MangaUrl:    data.MangaUrl,
			ChapterList: data.ChapterList,
		}}, mangas)
	}
	fmt.Printf("%+v\n", mangaList)
	return mangaList
}

func main() {
	var testReaper = models.WorkerInformation{
		WebsiteSearchUrl: "https://api.reaperscans.com/query?adult=true&query_string=",
		MangaName:        "Solo",
		MangaId:          0,
		//mangaUrl:             "https://api.reaperscans.com/chapters/100?page=1&perPage=1000&query=&order=desc",
		ChapterSearchPattern: "",
		MangaSearchPattern:   "",
		ChapterList:          nil,
	}
	worker(testReaper)
	//TODO: SEND RETRIEVED DATA TO THE DB
}
