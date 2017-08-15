package gonnman

import (
  "fmt"
  "github.com/nextthingco/gonnman"
)

func GetWifiTechnology() (*connman.Technology, error) {
  var err error
  var technologies []*connman.Technology
  var wifi *connman.Technology
  wifi = nil

  // fmt.Println("getWifiTechnology: scanning for connection technologies")
  if technologies, err = connman.GetTechnologies(); err != nil {
    fmt.Println("getWifiTechnology Error:", err)
    return nil, err
  }

  for _, tech := range technologies {
    if tech.Type == "wifi" {
      wifi = tech
    }
  }
  return wifi, nil
}

func GetServices() ([]*connman.Service, error) {
  if services, err := connman.GetServices(); err != nil {
    fmt.Println("getServices Error:", err)
    return nil, err
  } else {
    return services, nil
  }
}
