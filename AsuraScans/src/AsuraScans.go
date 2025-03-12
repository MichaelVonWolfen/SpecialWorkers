package main

import "SpecialWorkers/models"

func worker(data models.WorkerInformation) []models.MangaInformation {
	return nil
}
func main() {
	var testAsura = models.WorkerInformation{
		WebsiteSearchUrl: "https://api.reaperscans.com/query?adult=true&query_string=",
		MangaName:        "Solo",
		MangaId:          0,
		//mangaUrl:             "https://api.reaperscans.com/chapters/100?page=1&perPage=1000&query=&order=desc",
		WebsiteSearchPattern: "",
		MangaSearchPattern:   "",
		ChapterList:          nil,
	}
	worker(testAsura)
	//TODO: SEND RETRIEVED DATA TO THE DB
}
