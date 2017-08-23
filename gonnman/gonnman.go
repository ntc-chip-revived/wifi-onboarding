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
