package oem750x_test

import (
	"fmt"
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

	var channel uint = 1
	if err := parker.SetNormalMode(channel); err != nil {
		panic(err)
	}
	if err := parker.SetDisableSwitch(channel, protocol.DisableBoth); err != nil {
		panic(err)
	}
	if err := parker.SetTargetVelocity(channel, 0.5); err != nil {
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

func TestReadings(t *testing.T) {
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

	var channel uint = 1
	partNumber, err := parker.GetPartNumber(channel)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part number: %s\n", partNumber)
	indexerStatus, err := parker.GetIndexerStatus(channel)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexer status: %s\n", indexerStatus)
	closedLoopStatus, err := parker.GetClosedLoopStatus(channel)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Closed loop status: %s\n", closedLoopStatus)
	limitStatus, err := parker.GetLimitsStatus(channel)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Limit status: %s\n", limitStatus)
	position, err := parker.GetAbsolutePosition(channel)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Absolute position: %d\n", position)
	position, err = parker.GetRelativePosition(channel)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Relative position: %d\n", position)
}

func TestHoming(t *testing.T) {
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
	if err := parker.SetDisableSwitch(channel, protocol.EnableBoth); err != nil {
		panic(err)
	}
	if err := parker.SetEndLimitsState(channel, protocol.NormallyOpen); err != nil {
		panic(err)
	}
	if err := parker.SetResolution(channel, 50000); err != nil {
		panic(err)
	}
	if err := parker.GoHomeHard(channel, protocol.Backward, 1); err != nil {
		panic(err)
	}
	fmt.Println("Homing finished")
}
