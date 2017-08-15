package main

import (
	"errors"
	"github.com/nextthingco/gonnman"
)

func getAccessPointByPath(path string) (*connman.Service, error) {
	for _, service := range connectionServices {
		if string(service.Path) == path {
			return service, nil
		}
	}
	return nil, errors.New("Access point not found for path: " + path)
}

func clearAccessPointList() {
	Debug.Println("clearAccessPointList: Clearing...")
	connectionServices = []*connman.Service{}
}
