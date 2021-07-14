/*
 * This file is part of anilistWrapGo (https://github.com/Dank-del/anilistWrapGo).
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

package media

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Dank-del/anilistWrapGo/anilistWrapGo"
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
}
