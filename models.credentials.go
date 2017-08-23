// Copyright (C) 2017 Next Thing Co. <software@nextthing.co>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 2 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>

package main

type Credentials struct {
	SSID string `json:"ssid"`
	PSK  string `json:"psk"`
	Path string `json:"path"`
}

var credentials = Credentials{
	SSID: "",
	PSK:  "",
	Path: "",
}

func getCredentials() Credentials {
	Debug.Println("getCredentials: Getting!")
	return credentials
}

func setCredentialSSID(name string) {
	credentials.SSID = name
}

func setCredentialPSK(password string) {
	credentials.PSK = password
}

func setCredentialPath(path string) {
	credentials.Path = path
}

func clearCredentials() Credentials {
	Debug.Println("clearCredentials: Clearing!")
	credentials.SSID = ""
	credentials.PSK = ""
	credentials.Path = ""
	return credentials
}
