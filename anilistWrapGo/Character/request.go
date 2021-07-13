package Character

import (
	"anilistWrapGo/anilistWrapGo"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type M map[string]interface{}

func DoRequest(query string) (res *AnilistCharacter, err error) {

	gq, err := json.Marshal(M{"query": CharacterQuery, "variables": M{"query": query}})
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(anilistWrapGo.BaseUrl, "application/json", bytes.NewBuffer(gq))

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	response := new(AnilistCharacter)
	err = json.Unmarshal(b, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
