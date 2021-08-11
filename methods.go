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

package anilistWrapGO

import (
	"github.com/Dank-del/anilistWrapGO/anilistWrapGo/anime"
	characters "github.com/Dank-del/anilistWrapGO/anilistWrapGo/character"
	"github.com/Dank-del/anilistWrapGO/anilistWrapGo/media"
	"github.com/Dank-del/anilistWrapGO/anilistWrapGo/user"
)

/* --------------------------------------------------------------------------------------- */

/*
  Anime methods
*/

func GetAnimeByID(ID int64) (*anime.AnilistAnime, error) {
	return anime.IdRequest(ID)
}

func SearchAnime(query string) (*anime.AnilistAnime, error) {
	return anime.SearchRequest(query)
}

/*
  Character methods
*/

func GetCharacterByID(ID int64) (*characters.AnilistCharacter, error) {
	return characters.RequestById(ID)
}

func SearchCharacter(query string) (*characters.AnilistCharacter, error) {
	return characters.SearchRequest(query)
}

/*
  Media methods
*/

func GetMediaByID(ID int64) (*media.AnilistMedia, error) {
	return media.RequestByID(ID)
}

func SearchMedia(query string) (*media.AnilistMedia, error) {
	return media.SearchRequest(query)
}

/*
  User methods
*/

func GetUserByID(ID int64) (*user.AnilistUser, error) {
	return user.IDRequest(ID)
}

func GetUserByUsername(Username string) (*user.AnilistUser, error) {
	return user.UsernameRequest(Username)
}

/* --------------------------------------------------------------------------------------- */
