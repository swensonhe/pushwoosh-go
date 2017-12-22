# Pushwoosh-Go

Pushwoosh-Go is a [Pushwoosh](https://www.pushwoosh.com/) API client written in Go. Currently it can:

- Register devices
- Unregister devices
- Create messages

## Installation

```bash
$ go get github.com/swensonhe/pushwoosh-go
```

## Examples

### Register device

```go
package main

import (
  "github.com/swensonhe/pushwoosh-go"
  "fmt"
)

func main() {
	client := pushwoosh.NewClient("PUSHWOOSH_APPLICATION_CODE", "PUSHWOOSH_APPLICATION_TOKEN")
	device := &pushwoosh.Device{
        PushToken: "TOKEN",
        Hwid: "HWID",
        DeviceType: pushwoosh.DeviceTypeiOS,
    }
	err := client.RegisterDevice(device)
	if err != nil {
	    fmt.Println(err)
	}
}

```

### Unregister device

```go
package main

import (
  "github.com/swensonhe/pushwoosh-go"
  "fmt"
)

func main() {
	client := pushwoosh.NewClient("PUSHWOOSH_APPLICATION_CODE", "PUSHWOOSH_APPLICATION_TOKEN")
	device := &pushwoosh.Device{
        Hwid: "HWID",
    }
	err := client.UnregisterDevice(device)
	if err != nil {
	    fmt.Println(err)
	}
}

```

### Create message

```go
package main

import (
  "github.com/swensonhe/pushwoosh-go"
  "fmt"
)

func main() {
    client := pushwoosh.NewClient("PUSHWOOSH_APPLICATION_CODE", "PUSHWOOSH_APPLICATION_TOKEN")
    message := &pushwoosh.Message{
        Notifications: []*pushwoosh.Notification{
		    &pushwoosh.Notification{
				SendDate: "now",
				Devices: []string{"TOKEN"},
				Content: "MESSAGE",
				IosRootParams: pushwoosh.IosRootParams{
				    Aps: pushwoosh.Aps{
				        ContentAvailable: pushwoosh.NoContentAvailable
				    },
				    Data: n.Data,
				},
				IosSound: "default",
				IosSilent: pushwoosh.SilentSettingOff,
				Platforms: []int{pushwoosh.DeviceTypeiOS},
			},
		},
	}

	err := client.CreateMessage(message)
	if err != nil {
		fmt.Println(err)
	}
}
```
