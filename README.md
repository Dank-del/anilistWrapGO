# anilistWrapGo

> Unofficial Anilist.co GraphQL API wrapper for GoLang.


## Example

```go
package main

import (
	"github.com/Dank-del/anilistWrapGO/anilistWrapGo/anime"
	characters "github.com/Dank-del/anilistWrapGO/anilistWrapGo/character"
	"github.com/Dank-del/anilistWrapGO/anilistWrapGo/media"
	"github.com/Dank-del/anilistWrapGO/anilistWrapGo/user"
	"log"
)

func main() {
	// Search Anime
	animeres, err := anime.AnimeRequest("Youkoso Jitsuryoku Shijou Shugi no Kyoushitsu e")
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println(animeres.Data.Media.Title)
	// 2021/08/09 15:11:19 {Youkoso Jitsuryoku Shijou Shugi no Kyoushitsu e Youkoso Jitsuryoku Shijou Shugi no Kyoushitsu e Classroom of the Elite ようこそ実力至上主義の教室へ}

	// Search Character
	characterRes, err := characters.CharacterRequest("Rin Tohsaka")
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println(characterRes.Data.Character.Name.Native)
	// 2021/08/09 15:11:20 遠坂凛

	// Search Media
	mediaRes, err := media.MediaRequest("higehiro")
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println(mediaRes.Data.Media.Title.Native)
	// 2021/08/09 15:11:20 ひげを剃る。そして女子高生を拾う。

	// Search user
	userres, err := user.UserRequest("mimiee") // this person is a thot
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println(*userres.Data.User.Name)
	// 2021/08/11 16:03:38 Mimiee
}
```