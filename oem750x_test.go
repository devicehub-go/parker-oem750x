package oem750x_test

import (
	"testing"

	oem750x "github.com/devicehub-go/parker-oem750x"
	"github.com/devicehub-go/parker-oem750x/protocol"
	"github.com/devicehub-go/unicomm"
	"github.com/devicehub-go/unicomm/protocol/unicommserial"
)

func TestMoveOneRound(t *testing.T) {
	parker := oem750x.New(unicomm.Options{
		Protocol: unicomm.Serial,
		Serial: unicommserial.SerialOptions{
			PortName: "/dev/ttyUSB0",
			BaudRate: 9600,
			DataBits: 8,
			StopBits: unicommserial.OneStopBit,
			Parity:   unicommserial.NoParity,
		},
	})
	if err := parker.Connect(); err != nil {
		panic("Connection" + err.Error())
	}
	defer parker.Disconnect()

	var channel uint = 4
	if err := parker.SetNormalMode(channel); err != nil {
		panic(err)
	}
	if err := parker.SetDisableSwitch(channel, protocol.DisableBoth); err != nil {
		panic(err)
	}
	if err := parker.SetTargetVelocity(channel, 1); err != nil {
		panic(err)
	}
	if err := parker.SetTargetAcceleration(channel, 0.5); err != nil {
		panic(err)
	}
	if err := parker.SetResolution(channel, 50000); err != nil {
		panic(err)
	}
	if err := parker.SetTargetDistance(channel, 50000); err != nil {
		panic(err)
	}
	if err := parker.Go(channel); err != nil {
		panic(err)
	}
}
