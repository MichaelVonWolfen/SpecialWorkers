package main

import "time"

type WorkerInformation struct {
	websiteSearchUrl     string
	mangaName            string
	mangaId              int
	mangaUrl             string
	websiteSearchPattern string
	mangaSearchPattern   string
	chapterList          []MangaChapter
}
type MangaInformation struct {
	mangaId     int
	mangaName   string
	mangaUrl    string
	chapterList []MangaChapter
}
type MangaChapter struct {
	chapterId     int
	chapterUrl    string
	chapterName   string
	chapterNumber float64
	mangaId       int
}
type ReaperResponseStruct struct {
	Meta struct {
		Total           int         `json:"total"`
		PerPage         int         `json:"per_page"`
		CurrentPage     int         `json:"current_page"`
		LastPage        int         `json:"last_page"`
		FirstPage       int         `json:"first_page"`
		FirstPageUrl    string      `json:"first_page_url"`
		LastPageUrl     string      `json:"last_page_url"`
		NextPageUrl     interface{} `json:"next_page_url"`
		PreviousPageUrl interface{} `json:"previous_page_url"`
	} `json:"meta"`
	Data []struct {
		Id               int       `json:"id"`
		Title            string    `json:"title"`
		Description      string    `json:"description"`
		AlternativeNames string    `json:"alternative_names"`
		SeriesType       string    `json:"series_type"`
		SeriesSlug       string    `json:"series_slug"`
		Thumbnail        string    `json:"thumbnail"`
		Status           string    `json:"status"`
		CreatedAt        time.Time `json:"created_at"`
		Badge            string    `json:"badge"`
		Latest           string    `json:"latest"`
		Rating           float64   `json:"rating"`
		ReleaseSchedule  struct {
			Wed bool `json:"wed"`
		} `json:"release_schedule"`
		NuLink        interface{} `json:"nu_link"`
		IsComingSoon  bool        `json:"is_coming_soon"`
		DiscordRoleId string      `json:"discord_role_id"`
		Acronym       interface{} `json:"acronym"`
		Color         interface{} `json:"color"`
		FreeChapters  []struct {
			Id                int           `json:"id"`
			ChapterName       string        `json:"chapter_name"`
			ChapterSlug       string        `json:"chapter_slug"`
			CreatedAt         time.Time     `json:"created_at"`
			SeriesId          int           `json:"series_id"`
			Index             string        `json:"index"`
			ChaptersToBeFreed []interface{} `json:"chapters_to_be_freed"`
			NovelChapters     []interface{} `json:"novel_chapters"`
			Excerpt           interface{}   `json:"excerpt"`
			Meta              struct {
				AdonisGroupLimitCounter string `json:"adonis_group_limit_counter"`
			} `json:"meta"`
		} `json:"free_chapters"`
		PaidChapters []interface{} `json:"paid_chapters"`
		Meta         struct {
			Background interface{} `json:"background"`
			Metadata   struct {
			} `json:"metadata"`
			ChaptersCount string `json:"chapters_count"`
		} `json:"meta"`
	} `json:"data"`
}
type ReaperMangaData struct {
	Meta struct {
		Total           int         `json:"total"`
		PerPage         int         `json:"per_page"`
		CurrentPage     int         `json:"current_page"`
		LastPage        int         `json:"last_page"`
		FirstPage       int         `json:"first_page"`
		FirstPageUrl    string      `json:"first_page_url"`
		LastPageUrl     string      `json:"last_page_url"`
		NextPageUrl     interface{} `json:"next_page_url"`
		PreviousPageUrl interface{} `json:"previous_page_url"`
	} `json:"meta"`
	Data []struct {
		Id               int         `json:"id"`
		ChapterSlug      string      `json:"chapter_slug"`
		ChapterName      string      `json:"chapter_name"`
		ChapterTitle     interface{} `json:"chapter_title"`
		SeriesId         int         `json:"series_id"`
		Price            int         `json:"price"`
		Index            string      `json:"index"`
		Public           bool        `json:"public"`
		ChapterThumbnail string      `json:"chapter_thumbnail"`
		ChapterType      string      `json:"chapter_type"`
		CreatedAt        time.Time   `json:"created_at"`
		UpdatedAt        *time.Time  `json:"updated_at"`
		Series           struct {
			SeriesSlug string `json:"series_slug"`
			Id         int    `json:"id"`
			Meta       struct {
			} `json:"meta"`
		} `json:"series"`
		ChaptersToBeFreed []interface{} `json:"chapters_to_be_freed"`
		NovelChapters     []interface{} `json:"novel_chapters"`
		Excerpt           interface{}   `json:"excerpt"`
		Meta              struct {
		} `json:"meta"`
	} `json:"data"`
}
