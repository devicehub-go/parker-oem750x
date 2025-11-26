package oem750x

import (
	"github.com/devicehub-go/parker-oem750x/protocol"
	"github.com/devicehub-go/unicomm"
)

/*
Creates a new instance of Parker Drive OEM750X that
allows to communicate and control the connected motors

Paramaters:
  - options: communication options to connect with the
    device including the protocol and the respective options

Note: delimiter is not required
*/
func New(options unicomm.Options) *protocol.OEM750x {
	options.Delimiter = protocol.CR
	oem750 := &protocol.OEM750x{
		Communication: unicomm.New(options),
	}
	return oem750
}
