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

func TestUser(t *testing.T) {
	usernameResult, err := anilistWrapGO.GetUserByUsername("mimiee") // this person is a thot
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println(*usernameResult.Data.User.Name)

	idResult, err := anilistWrapGO.GetUserByID(724274) // its meee owo
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println(*idResult.Data.User.Name)
}
