/*
 * This file is part of anilistWrapGo (https://github.com/Dank-del/anilistWrapGO).
 * Copyright (c) 2021 Sayan Biswas, ALiwoto.
 *
 * This library is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, version 3.
 *
 * This library is distributed in the hope that it will be useful, but
 * WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU
 * General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this library. If not, see <http://www.gnu.org/licenses/>.
 */

package anime

import (
	"bytes"
	"encoding/json"
	"github.com/Dank-del/anilistWrapGO/anilistWrapGo"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type queryMap map[string]interface{}

func SearchRequest(query string) (*AnilistAnime, error) {
	gq, err := json.Marshal(getSearchQuery(query))
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(anilistWrapGo.BaseUrl, contentType, bytes.NewBuffer(gq))
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err.Error())
			return
		}
	}(resp.Body)

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	response := new(AnilistAnime)
	err = json.Unmarshal(b, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func IdRequest(ID int64) (*AnilistAnime, error) {
	gq, err := json.Marshal(getIDQuery(ID))
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(anilistWrapGo.BaseUrl, contentType, bytes.NewBuffer(gq))
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err.Error())
			return
		}
	}(resp.Body)

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	response := new(AnilistAnime)
	err = json.Unmarshal(b, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func getSearchQuery(query string) queryMap {
	return queryMap{
		queryKey: AnimeGraphQL,
		variablesValue: queryMap{
			searchKey: query,
		},
	}
}

func getIDQuery(ID int64) queryMap {
	return queryMap{
		queryKey: AnimeGraphQL,
		variablesValue: queryMap{
			idKey: strconv.FormatInt(ID, 10),
		},
	}
}
