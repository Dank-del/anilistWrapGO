# anilistWrapGO

[![Go Reference](https://pkg.go.dev/badge/github.com/Dank-del/anilistWrapGO.svg)](https://pkg.go.dev/github.com/Dank-del/anilistWrapGO)

An unofficial Go wrapper for the [AniList GraphQL API](https://anilist.gitbook.io/anilist-apiv2-docs).

## Install

```sh
go get github.com/Dank-del/anilistWrapGO
```

## Usage

Each helper returns a typed result struct and an error. Full examples live in the `test` directory.

### Look up a user by name

```go
res, err := anilistWrapGO.GetUserByUsername("mimiee")
if err != nil {
    log.Fatal(err)
}
log.Println(*res.Data.User.Name)
```

### Search for an anime

```go
res, err := anilistWrapGO.SearchAnime("kanojo mo kanojo")
if err != nil {
    log.Fatal(err)
}
log.Println(res.Data.Media.Title)
```

### Search for any media

```go
res, err := anilistWrapGO.SearchMedia("higehiro")
if err != nil {
    log.Fatal(err)
}
log.Println(res.Data.Media.Title.Native)
```

### Search for a character

```go
res, err := anilistWrapGO.SearchCharacter("Rin Tohsaka")
if err != nil {
    log.Fatal(err)
}
log.Println(res.Data.Character.Name.Native)
```

## Status

This library is a hobby project and is not affiliated with AniList. Pull requests welcome.
