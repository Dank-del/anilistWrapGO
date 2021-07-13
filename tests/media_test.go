package tests

import (
	"anilistWrapGo/anilistWrapGo/Media"
	"log"
	"testing"
)

func TestMedia(t *testing.T) {
	res, err := Media.DoRequest("higehiro")
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println(res)
}
