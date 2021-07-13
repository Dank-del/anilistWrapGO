package tests

import (
	"anilistWrapGo/anilistWrapGo/Character"
	"log"
	"testing"
)

func TestCharacter(t *testing.T) {
	res, err := Character.DoRequest("Rin Tohsaka")
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println(res.Data.Character.Name.Native)
}
