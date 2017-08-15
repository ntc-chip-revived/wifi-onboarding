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

import (
	"./gonnman"
	"./hostapd"
	"github.com/nextthingco/gonnman"
	"io/ioutil"
	"os"
	"time"
)

var wifi *connman.Technology
var connectionServices []*connman.Service

func main() {
	initLoggers(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)

	retryAttempts := 3
	for {
		Debug.Println("start")

		time.Sleep(time.Second * 3)
		if connectedToWifi() {
			Debug.Println("already connected - idling...")
			continue
		}

		Debug.Println("retrying attempts left", retryAttempts)
		if retryAttempts > 0 {
			Debug.Println("not connected: retrying in 3 seconds")
			retryAttempts -= 1
			continue
		}
		retryAttempts = 3

		// Search for wifi capability
		Debug.Println("exhausted all attempts, time to start up the wifi onboarding server")
		if wifiTechnology, err := gonnman.GetWifiTechnology(); err != nil || wifiTechnology == nil {
			Warning.Println(err)
			continue
		} else {
			// Save wifi reference to memory
			wifi = wifiTechnology
		}

		Debug.Println("enabling and scanning wifi")
		wifi.Enable()
		wifi.Scan()

		// Get connection services (access points)
		if services, err := gonnman.GetServices(); err != nil {
			Debug.Println(err)
			continue
		} else {
			// Save list of services in memory
			connectionServices = services
		}

		wifi.Disable()
		hostapd.Start()

		// Start serving API and web application
		creds, _ := startServer()
		Debug.Printf("received credentials from channel!\n\nssid=%s\npsk=%s\npath=%s\n\n", creds.SSID, creds.PSK, creds.Path)
		stopServer()

		hostapd.Stop()
		wifi.Enable()
		wifi.Scan()

		connectToAccessPoint(creds.Path, creds.PSK)
	}
}

// We want to get a fresh wifi object to make sure data is updated
func connectedToWifi() bool {
	if wifiTechnology, err := gonnman.GetWifiTechnology(); err != nil {
		Error.Println("connectedToWifi:", err)
	} else {
		wifi = wifiTechnology
	}

	wifi.Enable()
	setStatusConnected(wifi.Connected)
	return wifi.Connected
}
