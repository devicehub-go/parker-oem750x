package protocol

import (
	"fmt"
)

/*
Sets thet target velocity of the motor in revolutions
per second (rps)
*/
func (o *OEM750x) SetTargetVelocity(channel uint, value float64) error {
	if value < 0.001 || value > 50.0 {
		return fmt.Errorf("motor velocity must be between 0.01 and 50.0 rps")
	}
	msg := fmt.Sprintf("%dV%.2f", channel, value)
	return o.Write(msg)
}

/*
Gets the target velocity of the motor in rps
*/
func (o *OEM750x) GetTargetVelocity(channel uint) (float64, error) {
	msg := fmt.Sprintf("%dV", channel)
	return o.RequestFloat(msg)
}

/*
Sets the target acceleration of the motor in revolutions
per second squared (rps²)
*/
func (o *OEM750x) SetTargetAcceleration(channel uint, value float64) error {
	if value < 0.01 || value > 999.0 {
		return fmt.Errorf("motor acceleration must be between 0.01 and 100.0 rps²")
	}
	msg := fmt.Sprintf("%dA%.2f", channel, value)
	return o.Write(msg)
}

/*
Gets the target acceleration of the motor in rps²
*/
func (o *OEM750x) GetTargetAcceleration(channel uint) (float64, error) {
	msg := fmt.Sprintf("%dA", channel)
	return o.RequestFloat(msg)
}

/*
Sets the target distance of the motor in steps
*/
func (o *OEM750x) SetTargetDistance(channel uint, value int) error {
	if value > 2147483648 || value < -2147483648 {
		return fmt.Errorf("distance must be between ±2147483648")
	}
	msg := fmt.Sprintf("%dD%d", channel, value)
	return o.Write(msg)
}

/*
Gets the target distance of the motor in steps
*/
func (o *OEM750x) GetTargetDistance(channel uint) (int, error) {
	msg := fmt.Sprintf("%dD", channel)
	return o.RequestInt(msg)
}
