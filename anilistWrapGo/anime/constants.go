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

const (
	contentType    = "application/json"
	variablesValue = "variables"
	queryKey       = "query"
	searchKey      = "search"
)

const (
	AnimeGraphQL = `
query ($id: Int,$search: String) {
	 Media (id: $id, type: ANIME,search: $search) { 
		id
		title {
			userPreferred
			romaji
			english
			native
		}
		coverImage {
			extraLarge
			large
		}
		bannerImage
		startDate {
			year
			month
			day
		}
		endDate {
			year
			month
			day
		}
		description
		season
		seasonYear
		type
		format
		status(version: 2)
		episodes
		duration
		chapters
		volumes
		genres
		synonyms
		source(version: 2)
		isAdult
		isLocked
		meanScore
		averageScore
		popularity
		favourites
		hashtag
		countryOfOrigin
		isLicensed
		isFavourite
		isRecommendationBlocked
		nextAiringEpisode {
			airingAt
			timeUntilAiring
			episode
		}
		relations {
			edges {
				id
				relationType(version: 2)
				node {
					id
					title {
						userPreferred
					}
					format
					type
					status(version: 2)
					bannerImage
					coverImage {
						large
					}
				}
			}
		}
		characterPreview: characters(perPage: 6, sort: [ROLE, RELEVANCE, ID]) {
			edges {
				id
				role
				name
				voiceActors(language: JAPANESE, sort: [RELEVANCE, ID]) {
					id
					name {
						userPreferred
					}
					language: languageV2
					image {
						large
					}
				}
				node {
					id
					name {
						userPreferred
					}
					image {
						large
					}
				}
			}
		}
		staffPreview: staff(perPage: 8, sort: [RELEVANCE, ID]) {
			edges {
				id
				role
				node {
					id
					name {
						userPreferred
					}
					language: languageV2
					image {
						large
					}
				}
			}
		}
		studios {
			edges {
				isMain
				node {
					id
					name
				}
			}
		}
		reviewPreview: reviews(perPage: 2, sort: [RATING_DESC, ID]) {
			pageInfo {
				total
			}
			nodes {
				id
				summary
				rating
				ratingAmount
				user {
					id
					name
					avatar {
						large
					}
				}
			}
		}
		recommendations(perPage: 7, sort: [RATING_DESC, ID]) {
			pageInfo {
				total
			}
			nodes {
				id
				rating
				userRating
				mediaRecommendation {
					id
					title {
						userPreferred
					}
					format
					type
					status(version: 2)
					bannerImage
					coverImage {
						large
					}
				}
				user {
					id
					name
					avatar {
						large
					}
				}
			}
		}
		externalLinks {
			site
			url
		}
		streamingEpisodes {
			site
			title
			thumbnail
			url
		}
		trailer {
			id
			site
		}
		rankings {
			id
			rank
			type
			format
			year
			season
			allTime
			context
		}
		tags {
			id
			name
			description
			rank
			isMediaSpoiler
			isGeneralSpoiler
		}
		mediaListEntry {
			id
			status
			score
		}
		stats {
			statusDistribution {
				status
				amount
			}
			scoreDistribution {
				score
				amount
			}
		}
	}
}


`
)

const (
	Crunchyroll Site     = "Crunchyroll"
	Japanese    Language = "Japanese"
)
