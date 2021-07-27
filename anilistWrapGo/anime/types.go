/*
 * This file is part of anilistWrapGo (https://github.com/Dank-del/anilistWrapGO).
 * Copyright (c) 2021 Sayan Biswas, ALiwoto.
 *
 * This library is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, version 3.
 *
 * This library is distributed in the hope that it will be useful, but
 * WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU
 * General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this library. If not, see <http://www.gnu.org/licenses/>.
 */

package anime

import "github.com/Dank-del/anilistWrapGO/anilistWrapGo"

type AnilistAnime struct {
	Data   Data                   `json:"data"`
	Errors []anilistWrapGo.Errors `json:"errors"`
}

type Data struct {
	Media Media `json:"Media"`
}

type Media struct {
	ID                      int64              `json:"id"`
	Title                   Title              `json:"title"`
	CoverImage              CoverImage         `json:"coverImage"`
	BannerImage             string             `json:"bannerImage"`
	StartDate               Date               `json:"startDate"`
	EndDate                 Date               `json:"endDate"`
	Description             string             `json:"description"`
	Season                  string             `json:"season"`
	SeasonYear              int64              `json:"seasonYear"`
	Type                    string             `json:"type"`
	Format                  string             `json:"format"`
	Status                  string             `json:"status"`
	Episodes                int64              `json:"episodes"`
	Duration                int64              `json:"duration"`
	Chapters                interface{}        `json:"chapters"`
	Volumes                 interface{}        `json:"volumes"`
	Genres                  []string           `json:"genres"`
	Synonyms                []string           `json:"synonyms"`
	Source                  string             `json:"source"`
	IsAdult                 bool               `json:"isAdult"`
	IsLocked                bool               `json:"isLocked"`
	MeanScore               int64              `json:"meanScore"`
	AverageScore            int64              `json:"averageScore"`
	Popularity              int64              `json:"popularity"`
	Favourites              int64              `json:"favourites"`
	Hashtag                 string             `json:"hashtag"`
	CountryOfOrigin         string             `json:"countryOfOrigin"`
	IsLicensed              bool               `json:"isLicensed"`
	IsFavourite             bool               `json:"isFavourite"`
	IsRecommendationBlocked bool               `json:"isRecommendationBlocked"`
	NextAiringEpisode       interface{}        `json:"nextAiringEpisode"`
	Relations               Relations          `json:"relations"`
	CharacterPreview        CharacterPreview   `json:"characterPreview"`
	StaffPreview            StaffPreview       `json:"staffPreview"`
	Studios                 Studios            `json:"studios"`
	ReviewPreview           ReviewPreview      `json:"reviewPreview"`
	Recommendations         Recommendations    `json:"recommendations"`
	ExternalLinks           []ExternalLink     `json:"externalLinks"`
	StreamingEpisodes       []StreamingEpisode `json:"streamingEpisodes"`
	Trailer                 Trailer            `json:"trailer"`
	Rankings                []Ranking          `json:"rankings"`
	Tags                    []Tag              `json:"tags"`
	MediaListEntry          interface{}        `json:"mediaListEntry"`
	Stats                   Stats              `json:"stats"`
}

type CharacterPreview struct {
	Edges []CharacterPreviewEdge `json:"edges"`
}

type CharacterPreviewEdge struct {
	ID          int64             `json:"id"`
	Role        string            `json:"role"`
	Name        interface{}       `json:"name"`
	VoiceActors []VoiceActorClass `json:"voiceActors"`
	Node        VoiceActorClass   `json:"node"`
}

type VoiceActorClass struct {
	ID       int64     `json:"id"`
	Name     Name      `json:"name"`
	Image    Image     `json:"image"`
	Language *Language `json:"language,omitempty"`
}

type Image struct {
	Large string `json:"large"`
}

type Name struct {
	UserPreferred string `json:"userPreferred"`
}

type CoverImage struct {
	ExtraLarge string `json:"extraLarge"`
	Large      string `json:"large"`
}

type Date struct {
	Year  int64 `json:"year"`
	Month int64 `json:"month"`
	Day   int64 `json:"day"`
}

type ExternalLink struct {
	Site string `json:"site"`
	URL  string `json:"url"`
}

type Ranking struct {
	ID      int64   `json:"id"`
	Rank    int64   `json:"rank"`
	Type    string  `json:"type"`
	Format  string  `json:"format"`
	Year    *int64  `json:"year"`
	Season  *string `json:"season"`
	AllTime bool    `json:"allTime"`
	Context string  `json:"context"`
}

type Recommendations struct {
	PageInfo PageInfo              `json:"pageInfo"`
	Nodes    []RecommendationsNode `json:"nodes"`
}

type RecommendationsNode struct {
	ID                  int64               `json:"id"`
	Rating              int64               `json:"rating"`
	UserRating          string              `json:"userRating"`
	MediaRecommendation MediaRecommendation `json:"mediaRecommendation"`
	User                User                `json:"user"`
}

type MediaRecommendation struct {
	ID          int64   `json:"id"`
	Title       Name    `json:"title"`
	Format      string  `json:"format"`
	Type        string  `json:"type"`
	Status      string  `json:"status"`
	BannerImage *string `json:"bannerImage"`
	CoverImage  Image   `json:"coverImage"`
}

type User struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Avatar Image  `json:"avatar"`
}

type PageInfo struct {
	Total int64 `json:"total"`
}

type Relations struct {
	Edges []RelationsEdge `json:"edges"`
}

type RelationsEdge struct {
	ID           int64               `json:"id"`
	RelationType string              `json:"relationType"`
	Node         MediaRecommendation `json:"node"`
}

type ReviewPreview struct {
	PageInfo PageInfo            `json:"pageInfo"`
	Nodes    []ReviewPreviewNode `json:"nodes"`
}

type ReviewPreviewNode struct {
	ID           int64  `json:"id"`
	Summary      string `json:"summary"`
	Rating       int64  `json:"rating"`
	RatingAmount int64  `json:"ratingAmount"`
	User         User   `json:"user"`
}

type StaffPreview struct {
	Edges []StaffPreviewEdge `json:"edges"`
}

type StaffPreviewEdge struct {
	ID   int64           `json:"id"`
	Role string          `json:"role"`
	Node VoiceActorClass `json:"node"`
}

type Stats struct {
	StatusDistribution []StatusDistribution `json:"statusDistribution"`
	ScoreDistribution  []ScoreDistribution  `json:"scoreDistribution"`
}

type ScoreDistribution struct {
	Score  int64 `json:"score"`
	Amount int64 `json:"amount"`
}

type StatusDistribution struct {
	Status string `json:"status"`
	Amount int64  `json:"amount"`
}

type StreamingEpisode struct {
	Site      Site   `json:"site"`
	Title     string `json:"title"`
	Thumbnail string `json:"thumbnail"`
	URL       string `json:"url"`
}

type Studios struct {
	Edges []StudiosEdge `json:"edges"`
}

type StudiosEdge struct {
	IsMain bool       `json:"isMain"`
	Node   PurpleNode `json:"node"`
}

type PurpleNode struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Tag struct {
	ID               int64  `json:"id"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	Rank             int64  `json:"rank"`
	IsMediaSpoiler   bool   `json:"isMediaSpoiler"`
	IsGeneralSpoiler bool   `json:"isGeneralSpoiler"`
}

type Title struct {
	UserPreferred string `json:"userPreferred"`
	Romaji        string `json:"romaji"`
	English       string `json:"english"`
	Native        string `json:"native"`
}

type Trailer struct {
	ID   string `json:"id"`
	Site string `json:"site"`
}

type Language string

type Site string
