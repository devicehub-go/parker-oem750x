package protocol

import (
	"fmt"
	"time"
)

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
Executes the homing procedure with the current settings
*/
func (o *OEM750x) GoHome(channel uint, direction Direction, speed float64) error {
	if speed < 0.01 || 50 < speed {
		return fmt.Errorf("speed must be between 0.01 and 50.00, got %.2f", speed)
	}
	if direction != Forward && direction != Backward {
		return fmt.Errorf("direction must be '+' (forward) or '-' (backward), got %s", direction)
	}
	msg := fmt.Sprintf("%dGH%s%.2f", channel, direction, speed)
	return o.Write(msg)
}

/*
Executes the homing procedure for all motors
*/
func (o *OEM750x) GoHomeAll(direction Direction, speed float64) error {
	if speed < 0.01 || 50 < speed {
		return fmt.Errorf("speed must be between 0.01 and 50.00, got %.2f", speed)
	}
	if direction != Forward && direction != Backward {
		return fmt.Errorf("direction must be '+' (forward) or '-' (backward), got %s", direction)
	}
	msg := fmt.Sprintf("GH%s%.2f", direction, speed)
	return o.Write(msg)
}

/*
Executes the homing procedure when just exists CW or CCW limits switches.
This function blocks until the procedure is finished
*/
func (o *OEM750x) GoHomeHard(channel uint, direction Direction, velocity float64, limitSwitch string) error {
	var limitReached bool = false
	var limitSwitchReleased bool = false

	if err := o.Stop(channel); err != nil {
		return err
	} else if err := o.SetTargetVelocity(channel, velocity); err != nil {
		return err
	} else if err := o.SetDirection(channel, direction); err != nil {
		return err
	} else if err := o.SetContinuosMode(channel); err != nil {
		return err
	} else if err := o.Go(channel); err != nil {
		return err
	}

	for !limitReached {
		status, err := o.GetLimitsStatus(channel)
		if err != nil {
			return err
		}
		if limitSwitch == "CW" {
			limitReached = status[2] == '1'
		} else {
			limitReached = status[3] == '1'
		}
		time.Sleep(100 * time.Millisecond)
	}

	if err := o.SetTargetVelocity(channel, 0.01); err != nil {
		return err
	} else if err := o.SetDirection(channel, Toggle); err != nil {
		return err
	} else if err := o.Go(channel); err != nil {
		return err
	}

	for !limitSwitchReleased {
		status, err := o.GetLimitsStatus(channel)
		if err != nil {
			return err
		}
		if direction == Forward {
			limitSwitchReleased = status[2] == '0'
		} else {
			limitSwitchReleased = status[3] == '0'
		}
	}

	if err := o.Stop(channel); err != nil {
		return err
	} else if err := o.SetAbsoluteMode(channel); err != nil {
		return err
	} else if err := o.SetNormalMode(channel); err != nil {
		return err
	}
	return o.SetZeroPosition(channel)
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
