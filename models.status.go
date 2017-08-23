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

type Status struct {
	HasCredentials bool  `json:"has_credentials"`
	Connecting     bool  `json:"connecting"`
	Connected      bool  `json:"connected"`
	Error          error `json:"error"`
}

var application = Status{
	HasCredentials: false,
	Connecting:     false,
	Connected:      false,
	Error:          nil,
}

func getApplicationStatus() Status {
	Debug.Println("getApplicationStatus: Getting!", application)
	return application
}

func setStatusHasCredentials(creds bool) {
	application.HasCredentials = creds
}

func setStatusConnecting(connecting bool) {
	application.Connecting = connecting
}

func setStatusConnected(connection bool) {
	application.Connected = connection
}

func setStatusError(err error) {
	application.Error = err
}

func clearApplicationStatus() Status {
	Debug.Println("getApplicationStatus: Clearing!")
	setStatusHasCredentials(false)
	setStatusConnecting(false)
	setStatusConnected(false)
	setStatusError(nil)
	return application
}
