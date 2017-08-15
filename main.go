package main

import (
	"io/ioutil"
	"os"
	"time"
	"github.com/nextthingco/gonnman"
	"./gonnman"
	"./hostapd"
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
