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

const (
	UserGraphQL = `
query($id: Int, $name: String) {
	User(id: $id, name: $name) {
		id
		name
		previousNames {
			name
			updatedAt
		}
		avatar {
			large
		}
		bannerImage
		about (asHtml: true)
		isFollowing
		isFollower
		donatorTier
		donatorBadge
		createdAt
		moderatorRoles
		isBlocked
		bans
		options {
			profileColor
		}
		mediaListOptions {
			scoreFormat
		}
		statistics {
			anime {
				count
				meanScore
				standardDeviation
				minutesWatched
				episodesWatched
				genrePreview: genres(limit: 10, sort: COUNT_DESC) {
					genre
					count
				}
			}
			manga {
				count
				meanScore
				standardDeviation
				chaptersRead
				volumesRead
				genrePreview: genres(limit: 10, sort: COUNT_DESC) {
					genre
					count
				}
			}
		}
		stats {
			activityHistory {
				date
				amount
				level
			}
		}
		favourites {
			anime {
				edges {
					favouriteOrder
					node {
						id
						type
						status(version: 2)
						format
						isAdult
						bannerImage
						title {
							userPreferred
						}
						coverImage {
							large
						}
						startDate {
							year
						}
					}
				}
			}
			manga {
				edges {
					favouriteOrder
					node {
						id
						type
						status(version: 2)
						format
						isAdult
						bannerImage
						title {
							userPreferred
						}
						coverImage {
							large
						}
						startDate {
							year
						}
					}
				}
			}
			characters {
				edges {
					favouriteOrder
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
			staff {
				edges {
					favouriteOrder
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
			studios {
				edges {
					favouriteOrder
					node {
						id
						name
					}
				}
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
	searchKey      = "name"
	idKey          = "id"
)
