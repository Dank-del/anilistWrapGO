/*
 * This file is part of anilistWrapGo (https://github.com/Dank-del/anilistWrapGo).
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

package media

import "github.com/Dank-del/anilistWrapGo/anilistWrapGo"

type AnilistMedia struct {
	Data   *Data                  `json:"data"`
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
	Title          *Title          `json:"title"`
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
	AiringSchedule *AiringSchedule `json:"airingSchedule"`
	SiteURL        string          `json:"siteUrl"`
}

type Data struct {
	Media *Media `json:"Media"`
}
