package modegenericinterface

import (
	"io"

	modehttp "github.com/johnrichardrinehart/mode-http"
)

// Device provides a common interface for simulating a device connected to the MODE network. Conventionally, we refer to all receivers of commands and emitters of events as `devices` - this is the origin of the name. Currently, the MODE network APIs are provided over either HTTP(s)/WebSocket or MQTT.
type Device interface {
	Listen(string, int, chan<- *DeviceCommand)
}

type httpDevice struct {
	*modehttp.Device
}

func (dh *httpDevice) Listen(host string, port int, cmdchan chan<- *DeviceCommand) {
	rchan := make(chan *modehttp.DeviceCommand)
	dh.Device.Listen(host, port, rchan)
	cmd := <-rchan
	cmdchan <- &DeviceCommand{Action: cmd.Action, Parameters: cmd.Parameters}
}

// NewDevice returns a fully configured Device type ready to receive commands from the API.
func NewDevice(protocol string, host string, port int, id int, key string, w io.Writer) Device {
	switch protocol {
	case `http`:
		return &httpDevice{&modehttp.Device{ID: id, Key: key, Wc: nil, Hc: nil}}
	case `mqtt`:
		// TODO: Wrap up MQTT just like HTTP
	}
	return nil
}
