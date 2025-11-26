package protocol

import "fmt"

type MovementMode uint
type SwitchState uint
type IndexerMode uint
type Polarity uint
type DisableSwitch uint
type Direction string

const (
	Incremental    MovementMode  = 0
	Absolute       MovementMode  = 1
	NormallyClosed SwitchState   = 0
	NormallyOpen   SwitchState   = 1
	MotorSteps     IndexerMode   = 0
	EncoderSteps   IndexerMode   = 1
	Normal         Polarity      = 0
	Inverted       Polarity      = 1
	EnableBoth     DisableSwitch = 0
	DisableCW      DisableSwitch = 1
	DisableCCW     DisableSwitch = 2
	DisableBoth    DisableSwitch = 3
	Forward        Direction     = "+"
	Backward       Direction     = "-"
	Toggle         Direction     = ""
)

/*
Sets the indexer to operate with incremental or absolute
when in steps mode
*/
func (o *OEM750x) SetIndexerMovementMode(channel uint, mode MovementMode) error {
	if mode != Incremental && mode != Absolute {
		return fmt.Errorf("invalid movement mode: %d", mode)
	}
	msg := fmt.Sprintf("%dFSA%d", channel, mode)
	return o.Write(msg)
}

/*
Sets the active state of clockwise (CW) and counter-clockwise (CCW)
end-of-travel limit switches
*/
func (o *OEM750x) SetEndLimitsState(channel uint, mode SwitchState) error {
	if mode != NormallyClosed && mode != NormallyOpen {
		return fmt.Errorf("invalid switch state: %d", mode)
	}
	msg := fmt.Sprintf("%dOSA%d", channel, mode)
	return o.Write(msg)
}

/*
Sets the indexer to perfom moves in motor steps or encoder steps
*/
func (o *OEM750x) SetIndexerMode(channel uint, mode IndexerMode) error {
	if mode != MotorSteps && mode != EncoderSteps {
		return fmt.Errorf("invalid indexer mode: %d", mode)
	}
	msg := fmt.Sprintf("%dFSB%d", channel, mode)
	return o.Write(msg)
}

/*
Sets the direction polarity of the motor
*/
func (o *OEM750x) SetPolarity(channel uint, polarity Polarity) error {
	if polarity != Normal && polarity != Inverted {
		return fmt.Errorf("invalid polarity: %d", polarity)
	}
	msg := fmt.Sprintf("%dCMDDIR%d", channel, polarity)
	return o.Write(msg)
}

/*
Sets the resolution of the motor in steps per revolution
*/
func (o *OEM750x) SetResolution(channel uint, value uint) error {
	if value < 200 || value > 50800 {
		return fmt.Errorf("motor resolution must be between 200 and 50800")
	}
	msg := fmt.Sprintf("%dMR%d", channel, value)
	return o.Write(msg)
}

/*
Gets current motor resolution in steps per revolution
*/
func (o *OEM750x) GetResolution(channel uint) (int, error) {
	msg := fmt.Sprintf("%dMR", channel)
	return o.RequestInt(msg)
}

/*
Sets status of communication error checking
*/
func (o *OEM750x) SetErrorChecking(channel uint, enable bool) error {
	var value uint
	if enable {
		value = 1
	} else {
		value = 0
	}
	msg := fmt.Sprintf("%dSSE%d", channel, value)
	return o.Write(msg)
}

/*
Sets the shutdown status of the motor, that rapidly decreases
the motor current to zero and the system will ignore move
commands
*/
func (o *OEM750x) SetShutdown(channel uint, enable bool) error {
	var value uint
	if enable {
		value = 1
	} else {
		value = 0
	}
	msg := fmt.Sprintf("%dST%d", channel, value)
	return o.Write(msg)
}

/*
Sets disable status of end-of-travel limit switches
*/
func (o *OEM750x) SetDisableSwitch(channel uint, mode DisableSwitch) error {
	if mode != EnableBoth && mode != DisableCW &&
		mode != DisableCCW && mode != DisableBoth {
		return fmt.Errorf("invalid disable switch mode: %d", mode)
	}
	msg := fmt.Sprintf("%dLD%d", channel, mode)
	return o.Write(msg)
}

/*
Sets movement direction
*/
func (o *OEM750x) SetDirection(channel uint, direction Direction) error {
	if direction != Forward && direction != Backward && direction != Toggle {
		return fmt.Errorf("invalid direction: %s", direction)
	}
	msg := fmt.Sprintf("%dH%s", channel, direction)
	return o.Write(msg)
}
