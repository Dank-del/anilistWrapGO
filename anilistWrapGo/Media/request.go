package Media

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type M map[string]interface{}

func DoRequest(query string) (res *AnilistMedia, err error) {

	gq, err := json.Marshal(M{"query": MediaGraphql, "variables": M{"search": query}})
	if err != nil {
		return nil, err
	}
	resp, err := http.Post("https://graphql.anilist.co", "application/json", bytes.NewBuffer(gq))

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	response := new(AnilistMedia)
	err = json.Unmarshal(b, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
