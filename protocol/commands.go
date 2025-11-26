package protocol

import "fmt"

/*
Sets motor to operate in the mode where the it
will move the steps defined by the user when a go
command is sent
*/
func (o *OEM750x) SetNormalMode(channel uint) error {
	msg := fmt.Sprintf("%dMN", channel)
	return o.Write(msg)
}

/*
Sets motor to operate in the mode where the it
will move continuously uintil a stop command is sent
*/
func (o *OEM750x) SetContinuosMode(channel uint) error {
	msg := fmt.Sprintf("%dMC", channel)
	return o.Write(msg)
}

/*
Sets the positioning mode to absolute
*/
func (o *OEM750x) SetAbsoluteMode(channel uint) error {
	msg := fmt.Sprintf("%dMPA", channel)
	return o.Write(msg)
}

/*
Sets the positioning mode to incremental
*/
func (o *OEM750x) SetIncrementalMode(channel uint) error {
	msg := fmt.Sprintf("%dMPI", channel)
	return o.Write(msg)
}

/*
Sets the absolute position counter to zero
*/
func (o *OEM750x) SetZeroPosition(channel uint) error {
	msg := fmt.Sprintf("%dPZ", channel)
	return o.Write(msg)
}

/*
Moves the motor with the current settings for
steps or target position movement
*/
func (o *OEM750x) Go(channel uint) error {
	msg := fmt.Sprintf("%dG", channel)
	return o.Write(msg)
}

/*
Moves all available motors with the current settings
for steps or target position movement
*/
func (o *OEM750x) GoAll() error {
	return o.Write("G")
}

/*
Stops the motor
*/
func (o *OEM750x) Stop(channel uint) error {
	msg := fmt.Sprintf("%dS", channel)
	return o.Write(msg)
}

/*
Stops all available motors
*/
func (o *OEM750x) StopAll() error {
	return o.Write("S")
}

/*
Ceases the indexer immediately
*/
func (o *OEM750x) Kill(channel uint) error {
	msg := fmt.Sprintf("%dK", channel)
	return o.Write(msg)
}

/*
Returns all internal settings to their power-up values
*/
func (o *OEM750x) Reset(channel uint) error {
	msg := fmt.Sprintf("%dZ", channel)
	return o.Write(msg)
}

/*
Re-establishes the communication and identify the
cause of the communication error
*/
func (o *OEM750x) ResetCommunication(channel uint) (string, error) {
	msg := fmt.Sprintf("%d%%", channel)
	return o.RequestString(msg, false)
}
