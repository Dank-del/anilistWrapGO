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

package user

import "github.com/Dank-del/anilistWrapGO/anilistWrapGo"

type AnilistUser struct {
	Data   *Data                  `json:"data,omitempty"`
	Errors []anilistWrapGo.Errors `json:"errors"`
}

type Data struct {
	User *User `json:"User,omitempty"`
}

type User struct {
	ID               *int64            `json:"id,omitempty"`
	Name             *string           `json:"name,omitempty"`
	PreviousNames    []interface{}     `json:"previousNames,omitempty"`
	Avatar           *Avatar           `json:"avatar,omitempty"`
	BannerImage      *string           `json:"bannerImage,omitempty"`
	About            *string           `json:"about,omitempty"`
	IsFollowing      *bool             `json:"isFollowing,omitempty"`
	IsFollower       *bool             `json:"isFollower,omitempty"`
	DonatorTier      *int64            `json:"donatorTier,omitempty"`
	DonatorBadge     *string           `json:"donatorBadge,omitempty"`
	CreatedAt        *int64            `json:"createdAt,omitempty"`
	ModeratorRoles   interface{}       `json:"moderatorRoles"`
	IsBlocked        *bool             `json:"isBlocked,omitempty"`
	Bans             []interface{}     `json:"bans,omitempty"`
	Options          *Options          `json:"options,omitempty"`
	MediaListOptions *MediaListOptions `json:"mediaListOptions,omitempty"`
	Statistics       *Statistics       `json:"statistics,omitempty"`
	Stats            *Stats            `json:"stats,omitempty"`
	Favourites       *Favourites       `json:"favourites,omitempty"`
}

type Avatar struct {
	Large *string `json:"large,omitempty"`
}

type Favourites struct {
	Anime      *MangaClass `json:"anime,omitempty"`
	Manga      *MangaClass `json:"manga,omitempty"`
	Characters *Characters `json:"characters,omitempty"`
	Staff      *MangaClass `json:"staff,omitempty"`
	Studios    *MangaClass `json:"studios,omitempty"`
}

type MangaClass struct {
	Edges []AnimeEdge `json:"edges,omitempty"`
}

type AnimeEdge struct {
	FavouriteOrder *int64      `json:"favouriteOrder,omitempty"`
	Node           *PurpleNode `json:"node,omitempty"`
}

type PurpleNode struct {
	ID          *int64     `json:"id,omitempty"`
	Type        *string    `json:"type,omitempty"`
	Status      *string    `json:"status,omitempty"`
	Format      *string    `json:"format,omitempty"`
	IsAdult     *bool      `json:"isAdult,omitempty"`
	BannerImage *string    `json:"bannerImage,omitempty"`
	Title       *Title     `json:"title,omitempty"`
	CoverImage  *Avatar    `json:"coverImage,omitempty"`
	StartDate   *StartDate `json:"startDate,omitempty"`
}

type StartDate struct {
	Year *int64 `json:"year,omitempty"`
}

type Title struct {
	UserPreferred *string `json:"userPreferred,omitempty"`
}

type Characters struct {
	Edges []CharactersEdge `json:"edges,omitempty"`
}

type CharactersEdge struct {
	FavouriteOrder *int64      `json:"favouriteOrder,omitempty"`
	Node           *FluffyNode `json:"node,omitempty"`
}

type FluffyNode struct {
	ID    *int64  `json:"id,omitempty"`
	Name  *Title  `json:"name,omitempty"`
	Image *Avatar `json:"image,omitempty"`
}

type MediaListOptions struct {
	ScoreFormat *string `json:"scoreFormat,omitempty"`
}

type Options struct {
	ProfileColor *string `json:"profileColor,omitempty"`
}

type Statistics struct {
	Anime *StatisticsAnime `json:"anime,omitempty"`
	Manga *Manga           `json:"manga,omitempty"`
}

type StatisticsAnime struct {
	Count             *int64         `json:"count,omitempty"`
	MeanScore         *float64       `json:"meanScore,omitempty"`
	StandardDeviation *float64       `json:"standardDeviation,omitempty"`
	MinutesWatched    *int64         `json:"minutesWatched,omitempty"`
	EpisodesWatched   *int64         `json:"episodesWatched,omitempty"`
	GenrePreview      []GenrePreview `json:"genrePreview,omitempty"`
}

type GenrePreview struct {
	Genre *string `json:"genre,omitempty"`
	Count *int64  `json:"count,omitempty"`
}

type Manga struct {
	Count             *int64         `json:"count,omitempty"`
	MeanScore         *float64       `json:"meanScore,omitempty"`
	StandardDeviation *float64       `json:"standardDeviation,omitempty"`
	ChaptersRead      *int64         `json:"chaptersRead,omitempty"`
	VolumesRead       *int64         `json:"volumesRead,omitempty"`
	GenrePreview      []GenrePreview `json:"genrePreview,omitempty"`
}

type Stats struct {
	ActivityHistory []ActivityHistory `json:"activityHistory,omitempty"`
}

type ActivityHistory struct {
	Date   *int64 `json:"date,omitempty"`
	Amount *int64 `json:"amount,omitempty"`
	Level  *int64 `json:"level,omitempty"`
}
