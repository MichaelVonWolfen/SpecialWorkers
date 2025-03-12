package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"runtime/debug"
	"slices"
	"strconv"
	"strings"
)

func getMangaChapters(url string) ([]MangaChapter, error) {
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
	var chapterList []MangaChapter
	for _, data := range mangaData.Data {
		var chapterNB, err = strconv.ParseFloat(data.Index, 64)
		if err != nil {
			log.Println(err)
			debug.PrintStack()
			continue
		}
		var erlSeriesName = strings.ToLower(data.Series.SeriesSlug)
		erlSeriesName = strings.ReplaceAll(erlSeriesName, " ", "-")
		var urlChapterName = strings.ReplaceAll(data.ChapterName, " ", "-")
		urlChapterName = strings.ToLower(urlChapterName)
		var chapter = MangaChapter{
			chapterUrl:    fmt.Sprintf("https://reaperscans.com/series/%s/%s", erlSeriesName, urlChapterName),
			chapterName:   data.ChapterName,
			chapterNumber: chapterNB,
		}
		//fmt.Printf("%+v\n", chapter)
		chapterList = append(chapterList, chapter)
	}
	return chapterList, nil
}
func getMangaUrl(searchUrl string, mangaName string) ([]MangaInformation, error) {
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
	var mangaList []MangaInformation
	for _, data := range mangaRawData.Data {
		manga := MangaInformation{
			mangaName:   data.Title,
			mangaUrl:    fmt.Sprintf("https://api.reaperscans.com/chapters/%d?page=1&perPage=1000&query=&order=desc", data.Id),
			chapterList: nil,
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

func updateMangasAndChapters(firstMangaList []MangaInformation, secondMangaList []MangaInformation) []MangaInformation {

	for i := 0; i < len(firstMangaList); i++ {
		for j := 0; j < len(secondMangaList); j++ {
			if firstMangaList[i].mangaName == secondMangaList[j].mangaName {
				firstMangaList[i].chapterList = secondMangaList[j].chapterList
				firstMangaList[i].mangaId = int(math.Max(float64(firstMangaList[i].mangaId), float64(secondMangaList[j].mangaId)))
				firstMangaList[i].mangaUrl = secondMangaList[j].mangaUrl
			} else {
				contains := slices.ContainsFunc(firstMangaList, func(m MangaInformation) bool { return m.mangaName == secondMangaList[j].mangaName })
				if !contains {
					firstMangaList = append(firstMangaList, secondMangaList[j])
				}
			}
		}
	}

	return slices.DeleteFunc(firstMangaList, func(m MangaInformation) bool { return m.mangaUrl == "" })
}
func worker(data WorkerInformation) []MangaInformation {
	isUrlValid := true
	var mangaList []MangaInformation
	if data.mangaUrl != "" && len(data.mangaUrl) > 0 {
		chapterList, err := getMangaChapters(data.mangaUrl)
		if err != nil {
			if err.Error() == "url seems invalid" {
				log.Printf("url seems invalid for %s \n", data.mangaName)
				isUrlValid = false
			} else {
				log.Fatal(err)
			}
		}
		manga := MangaInformation{
			mangaName:   data.mangaName,
			mangaUrl:    data.mangaUrl,
			chapterList: chapterList,
		}
		mangaList = updateMangasAndChapters([]MangaInformation{manga}, []MangaInformation{
			{
				mangaId:     data.mangaId,
				mangaName:   data.mangaName,
				mangaUrl:    data.mangaUrl,
				chapterList: data.chapterList,
			},
		})
	} else {
		isUrlValid = false
	}
	if !isUrlValid {
		mangas, err := getMangaUrl(data.websiteSearchUrl, data.mangaName)
		if err != nil {
			log.Panic(err)
		}
		fmt.Printf("%+v\n", mangas)

		log.Printf("Trying to read manga again after url update!")
		for i := 0; i < len(mangas); i++ {
			chapterList, err := getMangaChapters(mangas[i].mangaUrl)
			if err != nil {
				log.Panic(err)
			}
			mangas[i].chapterList = chapterList
		}
		mangaList = updateMangasAndChapters([]MangaInformation{{
			mangaId:     data.mangaId,
			mangaName:   data.mangaName,
			mangaUrl:    data.mangaUrl,
			chapterList: data.chapterList,
		}}, mangas)
	}
	fmt.Printf("%+v\n", mangaList)
	return mangaList
}

func main() {
	var testReaper = WorkerInformation{
		websiteSearchUrl: "https://api.reaperscans.com/query?adult=true&query_string=",
		mangaName:        "Solo",
		mangaId:          0,
		//mangaUrl:             "https://api.reaperscans.com/chapters/100?page=1&perPage=1000&query=&order=desc",
		websiteSearchPattern: "",
		mangaSearchPattern:   "",
		chapterList:          nil,
	}
	worker(testReaper)
	//TODO: SEND RETRIEVED DATA TO THE DB
}
