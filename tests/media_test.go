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

package tests

import (
	"github.com/Dank-del/anilistWrapGO"
	"log"
	"testing"
)

func TestMedia(t *testing.T) {
	searchData, err := anilistWrapGO.SearchMedia("higehiro")
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println(searchData.Data.Media.Title.Native)

	getResult, err := anilistWrapGO.GetMediaByID(117343)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println(getResult.Data.Media.Title.Native)
}
