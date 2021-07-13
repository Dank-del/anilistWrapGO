package Character

const (
	CharacterQuery = ` query ($query: String) {
        Character (search: $query) {
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
