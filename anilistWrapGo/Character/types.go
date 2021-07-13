package Character

import "anilistWrapGo/anilistWrapGo"

type AnilistCharacter struct {
	Data   Data                   `json:"data"`
	Errors []anilistWrapGo.Errors `json:"errors"`
}
type Name struct {
	First         string   `json:"first"`
	Last          string   `json:"last"`
	Full          string   `json:"full"`
	Native        string   `json:"native"`
	Alternative   []string `json:"alternative"`
	UserPreferred string   `json:"userPreferred"`
}
type Image struct {
	Large string `json:"large"`
}
type DateOfBirth struct {
	Year  interface{} `json:"year"`
	Month int         `json:"month"`
	Day   int         `json:"day"`
}
type Title struct {
	Romaji  string      `json:"romaji"`
	English interface{} `json:"english"`
	Native  string      `json:"native"`
}
type Nodes struct {
	Title   Title  `json:"title"`
	Type    string `json:"type"`
	Format  string `json:"format"`
	SiteURL string `json:"siteUrl"`
}
type Media struct {
	Nodes []Nodes `json:"nodes"`
}
type Character struct {
	ID          int         `json:"id"`
	Name        Name        `json:"name"`
	SiteURL     string      `json:"siteUrl"`
	Favourites  int         `json:"favourites"`
	Image       Image       `json:"image"`
	Description string      `json:"description"`
	Gender      string      `json:"gender"`
	Age         interface{} `json:"age"`
	DateOfBirth DateOfBirth `json:"dateOfBirth"`
	Media       Media       `json:"media"`
}
type Data struct {
	Character Character `json:"Character"`
}
