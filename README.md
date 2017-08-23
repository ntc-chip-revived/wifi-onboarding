### Wifi Onboarding

Copyright (C) 2017 Next Thing Co. <software@nextthing.co>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 2 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>

---
#### API

Method | Path | Use
--- | --- | ---
`GET` | `/ap/list` | Returns an array of [Service structs](https://github.com/NextThingCo/gonnman/blob/master/service.go#L45).
`GET` | `/ap/status` | Returns a Status struct.
`POST` | `/ap/connect` | Accepts values from a urlencoded form or multipart form. Returns a Status struct.

#### Example Calls
Request access point list:

`curl -H 'Accept: application/json' http://192.168.84.1:8080/ap/list`
<br>Protip: You can use any arbitrary address

Example Response:
```json
[
  {
    "Path": "/net/connman/service/<example path>",
    "Name": "<example SSID>",
    "Type": "wifi",
    "State": "idle",
    "Error": "",
    ...
  }
]
```

Request a connection:

`curl -X POST -H 'Accept: application/json' -F 'accessPointPath=<Path here>' -F 'accessPointPassKey=<Password here>' http://192.168.84.1:8080/ap/connect`

Example Response:

```go
{
  "has_credentials": true,
  "connecting": true,
  "connected": false,
  "error": nil,
}
```

#### Status Struct

```go
  type Status struct {
  HasCredentials bool
  Connecting     bool
  Connected      bool
  Error          error
}
```

#### Flow
![](http://d2rchup4fs07xx.cloudfront.net/images/wifi_onboarding_flow.png)
