package modegenericinterface

import (
	"io"

	modehttp "github.com/johnrichardrinehart/mode-http"
)

// Application provides a common interface for sending commands using the MODE network APIs. Conventionally, we refer to all senders of commands as `applications` - this is the origin of the name. Currently, the MODE network APIs are provided over either HTTP(s)/WebSocket or MQTT.
type Application interface {
	SendCommand(string, int, int, *DeviceCommand, io.Writer) error
}

// DeviceCommand defines the nature of the structure sent by Applications and received by Devices
type DeviceCommand struct {
	Action     string
	Parameters map[string]interface{}
}

// Application is the structure imitating an application connecting over HTTP(S)/WebSocket
type httpApplication struct {
	Key string
	*modehttp.Application
}

func (a *httpApplication) SendCommand(host string, port int, deviceID int, cmd *DeviceCommand, w io.Writer) error {
	httpcmd := &modehttp.DeviceCommand{}
	httpcmd.Action = cmd.Action
	httpcmd.Parameters = cmd.Parameters
	return a.Application.SendCommand(host, port, deviceID, httpcmd, w)
}

// NewApplication returns a fully configured Commander type ready to send commands to devices.
func NewApplication(protocol string, key string, w io.Writer) Application {
	switch protocol {
	case `http`:
		return &httpApplication{Key: key, Application: &modehttp.Application{Key: key, C: nil}}
	case `mqtt`:
		// TODO: Wrap up MQTT just like HTTP
		return nil
	}
	return nil
}
