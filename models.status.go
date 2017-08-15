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
