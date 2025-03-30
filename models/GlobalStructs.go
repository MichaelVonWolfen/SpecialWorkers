package models

type WorkerInformation struct {
	WebsiteSearchUrl     string
	MangaName            string
	MangaId              int
	MangaUrl             string
	WebsiteSearchPattern string
	MangaSearchPattern   string
	ChapterList          []MangaChapter
}
type MangaInformation struct {
	MangaId     int
	Status      string
	MangaName   string
	MangaUrl    string
	ChapterList []MangaChapter
}
type MangaChapter struct {
	ChapterId     int
	ChapterUrl    string
	ChapterName   string
	ChapterNumber float64
	MangaId       int
}
