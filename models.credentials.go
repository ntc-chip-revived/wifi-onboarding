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
