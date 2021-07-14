package media

import (
	"anilistWrapGo/anilistWrapGo"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type queryMap map[string]interface{}

func DoRequest(query string) (*AnilistMedia, error) {
	gq, err := json.Marshal(getQuery(query))
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(anilistWrapGo.BaseUrl, contentType, bytes.NewBuffer(gq))
	if err != nil {
		return nil, err
	}

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

func getQuery(query string) queryMap {
	return queryMap{
		queryKey: MediaGraphql,
		variablesValue: queryMap{
			searchKey: query,
		},
	}

	// m{"query": MediaGraphql, "variables": m{"search": query}}
}
