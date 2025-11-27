package protocol

import (
	"fmt"
	"strconv"
	"strings"
)

type IndexerStatus string

const (
	IndexerReady          IndexerStatus = "R"
	IndexerReadyAttention IndexerStatus = "S"
	IndexerBusy           IndexerStatus = "B"
	IndexerBusyAttention  IndexerStatus = "C"
)

/*
Gets the software part number and its revision level
*/
func (o *OEM750x) GetPartNumber(channel uint) (string, error) {
	msg := fmt.Sprintf("%dRV", channel)
	return o.RequestString(msg, false)
}

/*
Gets the general status of the indexer

Busy (B): indicates that the indexer is executing a
command (e.g., moving, waiting trigger, pausing, etc.)

Attention (S or C): indicates a drive fault, go home failed,
limit end reached, unsuccessful sequence or memory
checksum error
*/
func (o *OEM750x) GetIndexerStatus(channel uint) (IndexerStatus, error) {
	msg := fmt.Sprintf("%dR", channel)
	response, err := o.RequestString(msg, true)
	if err != nil {
		return "", err
	}
	return IndexerStatus(response), nil
}

/*
Gets the closed loop status

The response is a 2-character string where:
- The first indicates if at the last move the indexer detected a stall
- The second indicates if the homing procedures occurs sucessfully
*/
func (o *OEM750x) GetClosedLoopStatus(channel uint) (string, error) {
	msg := fmt.Sprintf("%dRA", channel)
	response, err := o.RequestString(msg, true)
	if err != nil {
		return "", err
	}
	switch response {
	case "@":
		return "01", nil
	case "A":
		return "11", nil
	case "B":
		return "00", nil
	case "C":
		return "10", nil
	}
	return "", fmt.Errorf("invalid limits status response: %s", response)
}

/*
Retrieves the status of the end-of-travel limits for the specified channel.

The response is a 4-character string where:
- The first two characters represent the last move terminated by CW and CCW limits.
- The last two characters represent the current condition of CW and CCW limits.
*/
func (o *OEM750x) GetLimitsStatus(channel uint) (string, error) {
	msg := fmt.Sprintf("%dRA", channel)
	response, err := o.RequestString(msg, true)
	if err != nil {
		return "", err
	}
	switch response {
	case "@":
		return "0000", nil
	case "A":
		return "1000", nil
	case "B":
		return "0100", nil
	case "D":
		return "0010", nil
	case "E":
		return "1010", nil
	case "F":
		return "0110", nil
	case "H":
		return "0001", nil
	case "I":
		return "1001", nil
	case "J":
		return "0101", nil
	case "L":
		return "0011", nil
	case "M":
		return "1011", nil
	case "N":
		return "0111", nil
	}
	return "", fmt.Errorf("invalid limits status response: %s", response)
}

/*
Gets the absolute position of the motor in steps
*/
func (o *OEM750x) GetAbsolutePosition(channel uint) (int, error) {
	msg := fmt.Sprintf("%dPR", channel)
	return o.RequestInt(msg)
}

/*
Gets an immediate position relative to start of the current
move in steps
*/
func (o *OEM750x) GetRelativePosition(channel uint) (int, error) {
	msg := fmt.Sprintf("%dW3", channel)
	response, err := o.RequestString(msg, false)
	if err != nil {
		return 0, err
	}
	response = strings.TrimPrefix(response, "*")
	value, err := strconv.ParseUint(response, 16, 64)
	if err != nil {
		return 0, err
	}
	return int(int32(uint32(value))), nil
}
