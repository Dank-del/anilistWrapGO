package Media

import "anilistWrapGo/anilistWrapGo"

type AnilistMedia struct {
	Data   Data                   `json:"data"`
	Errors []anilistWrapGo.Errors `json:"errors"`
}
type Title struct {
	Romaji  string `json:"romaji"`
	English string `json:"english"`
	Native  string `json:"native"`
}
type ExternalLinks struct {
	URL string `json:"url"`
}
type AiringSchedule struct {
	Nodes []interface{} `json:"nodes"`
}
type Media struct {
	ID             int             `json:"id"`
	Title          Title           `json:"title"`
	Type           string          `json:"type"`
	Format         string          `json:"format"`
	Status         string          `json:"status"`
	Description    string          `json:"description"`
	Episodes       int             `json:"episodes"`
	BannerImage    string          `json:"bannerImage"`
	ExternalLinks  []ExternalLinks `json:"externalLinks"`
	Duration       int             `json:"duration"`
	Chapters       interface{}     `json:"chapters"`
	Volumes        interface{}     `json:"volumes"`
	Genres         []string        `json:"genres"`
	Synonyms       []string        `json:"synonyms"`
	AverageScore   int             `json:"averageScore"`
	AiringSchedule AiringSchedule  `json:"airingSchedule"`
	SiteURL        string          `json:"siteUrl"`
}
type Data struct {
	Media Media `json:"Media"`
}
