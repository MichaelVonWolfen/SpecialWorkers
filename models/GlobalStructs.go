package models

type WorkerInformation struct {
	WebsiteSearchUrl     string
	BaseWebsiteUrl       string
	MangaName            string
	MangaId              int
	MangaUrl             string
	ChapterSearchPattern string
	MangaSearchPattern   string
	ChapterList          []MangaChapter
}
type MangaInformation struct {
	MangaId       int
	Status        string
	MangaName     string
	MangaUrl      string
	MangaImageUrl string
	ChapterList   []MangaChapter
}
type MangaChapter struct {
	ChapterId     int
	ChapterUrl    string
	ChapterName   string
	ChapterNumber string
	MangaId       int
}
