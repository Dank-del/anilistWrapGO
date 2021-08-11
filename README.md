# anilistWrapGo

> Unofficial Anilist.co GraphQL API wrapper for GoLang.


## Examples

### All examples are present as tests in `test` directory.
Below are a few snippets from there :3

#### User 
```go
   usernameResult, err := anilistWrapGO.GetUserByUsername("mimiee") // this person is a thot
   if err != nil {
   log.Fatal(err.Error())
   }
   log.Println(*usernameResult.Data.User.Name)
```

#### Anime 
```go
   searchResult, err := anilistWrapGO.SearchAnime("kanojo mo kanojo") // stupid anime
   if err != nil {
   log.Fatal(err.Error())
   }
   log.Println(searchResult.Data.Media.Title)
```

#### Media
```go
   searchData, err := anilistWrapGO.SearchMedia("higehiro") // thot anime
   if err != nil {
   log.Fatal(err.Error())
   }
   log.Println(searchData.Data.Media.Title.Native)
```

#### Character
```go
   searchResult, err := anilistWrapGO.SearchCharacter("Rin Tohsaka") // goddess
   if err != nil {
   log.Fatal(err.Error())
   }

   log.Println(searchResult.Data.Character.Name.Native)
```