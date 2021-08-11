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

package characters

const (
	CharacterQuery = `
query ($query: String, $id: Int) {
  Character(search: $query, id: $id) {
    id
    name {
      first
      last
      full
      native
      alternative
      userPreferred
    }
    siteUrl
    favourites
    image {
      large
    }
    description(asHtml: false)
    gender
    age
    dateOfBirth {
      year
      month
      day
    }
    media {
      nodes {
        title {
          romaji
          english
          native
        }
        type
        format
        siteUrl
      }
    }
  }
}

`
)

const (
	contentType    = "application/json"
	variablesValue = "variables"
	queryKey       = "query"
	idKey          = "id"
)
