package media

const (
	MediaGraphql = `
query ($search: String) {
    Media (search: $search) {
      id
      title {
        romaji
        english
        native
      }
      type
      format
      status
      description
      episodes
      bannerImage
    externalLinks{
      url
    }
      duration
      chapters
      volumes
      genres
      synonyms
      averageScore
      airingSchedule(notYetAired: true) {
        nodes {
          airingAt
          timeUntilAiring
          episode
        }
      }
      siteUrl
    }
  }

`
)

const (
	contentType    = "application/json"
	variablesValue = "variables"
	queryKey       = "query"
	searchKey      = "search"
)
