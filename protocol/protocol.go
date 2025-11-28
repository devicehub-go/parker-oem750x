package protocol

import (
	"fmt"
	"regexp"
	"strconv"
	"sync"

	"github.com/devicehub-go/unicomm"
)

const (
	CR string = "\r"
)

type OEM750x struct {
	Communication unicomm.Unicomm
	mutex         sync.Mutex
}

/*
Establishes a connection with the device
*/
func (o *OEM750x) Connect() error {
	return o.Communication.Connect()
}

/*
Closes the connection with the device
*/
func (o *OEM750x) Disconnect() error {
	return o.Communication.Disconnect()
}

/*
Returns true if device is connected
*/
func (o *OEM750x) IsConnected() bool {
	return o.Communication.IsConnected()
}

/*
Writes a message to the device
*/
func (o *OEM750x) Write(message string) error {
	if !o.IsConnected() {
		return fmt.Errorf("device not connected")
	}
	err := o.Communication.Write([]byte(message + CR))
	if err != nil {
		return err
	}
	echo, err := o.Communication.ReadUntil(CR)
	if err != nil {
		return err
	}
	if string(echo) != message+CR {
		return fmt.Errorf("unexpected response: %s", string(echo))
	}
	return nil
}

/*
Cleans the response byte array to prevent characters
as CR and NULL from being present
*/
func cleanResponse(response []byte) []byte {
	cleaned := make([]byte, 0, len(response))
	for _, b := range response {
		if b != '\r' && b != '\x00' {
			cleaned = append(cleaned, b)
		}
	}
	return cleaned
}

/*
Sends a command to the device and returns
the response
*/
func (o *OEM750x) Request(message string) ([]byte, error) {
	if !o.IsConnected() {
		return nil, fmt.Errorf("device not connected")
	}
	o.mutex.Lock()
	err := o.Write(message)
	if err != nil {
		return nil, err
	}
	response, err := o.Communication.ReadUntil(CR)
	if err != nil {
		return nil, err
	}
	o.mutex.Unlock()
	return cleanResponse(response), nil
}

/*
Parses the drive response
*/
func ParseValueResponse(response []byte) (string, error) {
	var expected = regexp.MustCompile(`^\*\d+[A-Z]*(.+)$`)
	responseStr := string(response)
	matches := expected.FindStringSubmatch(responseStr)
	if len(matches) != 2 {
		return "", fmt.Errorf("invalid response format: %s", responseStr)
	}
	return matches[1], nil
}

/*
Request a string value from the device
*/
func (o *OEM750x) RequestString(message string, parse bool) (string, error) {
	response, err := o.Request(message)
	if err != nil {
		return "", err
	}
	if parse {
		return ParseValueResponse(response)
	}
	return string(response), nil
}

/*
Request an integer value from the device
*/
func (o *OEM750x) RequestInt(message string) (int, error) {
	response, err := o.Request(message)
	if err != nil {
		return 0, err
	}
	valueStr, err := ParseValueResponse(response)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(valueStr)
}

/*
Request a float64 value from the device
*/
func (o *OEM750x) RequestFloat(message string) (float64, error) {
	response, err := o.Request(message)
	if err != nil {
		return 0, err
	}
	valueStr, err := ParseValueResponse(response)
	if err != nil {
		return 0, err
	}
	return strconv.ParseFloat(valueStr, 64)
}
