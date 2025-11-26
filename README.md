# Parker OEM750x Driver

A Go driver for communicating with Parker OEM750x series motor controllers. This package provides a simple and intuitive interface for controlling stepper motors through serial communication, including position control, velocity management, and status monitoring.

## Features

- **Simple Initialization**: Easy device creation with flexible communication options
- **Serial Communication**: Built-in support for RS-232/RS-485 serial protocols
- **Connection Management**: Automatic connection state tracking and management
- **Motion Control**: Support for absolute and incremental positioning modes
- **Velocity & Acceleration**: Configure motor speed and acceleration parameters
- **Status Monitoring**: Real-time indexer status and position tracking
- **Limit Detection**: End-of-travel limit status retrieval
- **Multiple Operation Modes**: Normal (stepped) and continuous motion modes
- **Multi-Motor Support**: Control individual motors or all motors simultaneously
- **Type-Safe API**: Strongly-typed methods with proper error handling

## Installation
```bash
go get github.com/devicehub-go/parker-oem750x
go get github.com/devicehub-go/unicomm
```

## Quick Start

### Basic Usage
```go
package main

import (
    "log"
    oem750x "github.com/devicehub-go/parker-oem750x"
    "github.com/devicehub-go/unicomm"
    "github.com/devicehub-go/unicomm/protocol/unicommserial"
)

func main() {
    // Create a new Parker OEM750x instance
    parker := oem750x.New(unicomm.Options{
        Protocol: unicomm.Serial,
        Serial: unicommserial.SerialOptions{
            PortName: "/dev/ttyUSB0",
            BaudRate: 9600,
            DataBits: 8,
            StopBits: 1,
            Parity:   unicommserial.NoParity,
        },
    })
    
    // Connect to the device
    if err := parker.Connect(); err != nil {
        log.Fatal(err)
    }
    defer parker.Disconnect()
    
    // Configure motor on channel 1
    channel := uint(1)
    
    // Set normal mode (move specified steps)
    parker.SetNormalMode(channel)
    
    // Set velocity to 5 revolutions per second
    parker.SetTargetVelocity(channel, 5.0)
    
    // Set acceleration to 10 rps²
    parker.SetTargetAcceleration(channel, 10.0)
    
    // Move 1000 steps
    parker.SetTargetDistance(channel, 1000)
    
    // Execute the move
    parker.Go(channel)
}
```

## API Reference

### Interface

#### Package Function

**`New(options unicomm.Options) *protocol.OEM750x`**

Creates a new instance of Parker OEM750x driver.

**Parameters:**
- `options`: Communication options including protocol type and protocol-specific settings
  - The delimiter is automatically set to `\r` (carriage return)

**Returns:** Pointer to an initialized `OEM750x` instance

**Example:**
```go
parker := oem750x.New(unicomm.Options{
    Protocol: unicomm.Serial,
    Serial: unicommserial.SerialOptions{
        PortName: "/dev/ttyUSB0",
        BaudRate: 9600,
        DataBits: 8,
        StopBits: 1,
        Parity:   unicommserial.NoParity,
    },
})
```

#### Methods

**Connection Management**
- `Connect() error` - Establishes connection with the device
- `Disconnect() error` - Closes the connection with the device
- `IsConnected() bool` - Returns true if device is connected

**Motion Control**
- `SetNormalMode(channel uint) error` - Sets motor to move a defined number of steps
- `SetContinuosMode(channel uint) error` - Sets motor to move continuously until stopped
- `SetAbsoluteMode(channel uint) error` - Sets positioning mode to absolute
- `SetIncrementalMode(channel uint) error` - Sets positioning mode to incremental
- `SetZeroPosition(channel uint) error` - Sets the absolute position counter to zero
- `Go(channel uint) error` - Moves the motor with current settings
- `GoAll() error` - Moves all available motors
- `Stop(channel uint) error` - Stops the specified motor
- `StopAll() error` - Stops all available motors
- `Kill(channel uint) error` - Ceases the indexer immediately

**Configuration**
- `SetTargetVelocity(channel uint, value float64) error` - Sets velocity in rps (0.001-50.0)
- `GetTargetVelocity(channel uint) (float64, error)` - Gets velocity in rps
- `SetTargetAcceleration(channel uint, value float64) error` - Sets acceleration in rps² (0.01-999.0)
- `GetTargetAcceleration(channel uint) (float64, error)` - Gets acceleration in rps²
- `SetTargetDistance(channel uint, value int) error` - Sets distance in steps (±2,147,483,648)
- `GetTargetDistance(channel uint) (int, error)` - Gets distance in steps

**Status & Monitoring**
- `GetPartNumber(channel uint) (string, error)` - Gets software part number and revision
- `GetIndexerStatus(channel uint) (IndexerStatus, error)` - Gets indexer status (Ready, Busy, Attention)
- `GetLimitsStatus(channel uint) (string, error)` - Gets end-of-travel limit status (4-character string)
- `GetAbsolutePosition(channel uint) (int, error)` - Gets absolute position in steps
- `GetRelativePosition(channel uint) (int, error)` - Gets position relative to current move start

**System**
- `Reset(channel uint) error` - Returns settings to power-up values
- `ResetCommunication(channel uint) (string, error)` - Re-establishes communication

### Configuration

**IndexerStatus Constants**
```go
const (
    IndexerReady          IndexerStatus = "R"  // Ready to accept commands
    IndexerReadyAttention IndexerStatus = "S"  // Ready but attention required
    IndexerBusy           IndexerStatus = "B"  // Executing a command
    IndexerBusyAttention  IndexerStatus = "C"  // Busy and attention required
)
```

**Serial Communication Options**
```go
unicommserial.SerialOptions{
    PortName: "/dev/ttyUSB0",  // Serial port path
    BaudRate: 9600,            // Communication speed
    DataBits: 8,               // Data bits (typically 8)
    StopBits: 1,               // Stop bits (typically 1)
    Parity:   unicommserial.NoParity,  // Parity setting
}
```

**Parameter Limits**
- Velocity: 0.001 - 50.0 rps
- Acceleration: 0.01 - 999.0 rps²
- Distance: ±2,147,483,648 steps

## Examples

### Complete Motor Rotation
```go
parker := oem750x.New(unicomm.Options{
    Protocol: unicomm.Serial,
    Serial: unicommserial.SerialOptions{
        PortName: "/dev/ttyUSB0",
        BaudRate: 9600,
        DataBits: 8,
        StopBits: 1,
        Parity:   unicommserial.NoParity,
    },
})

if err := parker.Connect(); err != nil {
    panic(err)
}
defer parker.Disconnect()

channel := uint(1)

// Configure for one complete rotation
parker.SetNormalMode(channel)
parker.SetTargetVelocity(channel, 1.0)
parker.SetTargetAcceleration(channel, 0.5)
parker.SetResolution(channel, 50000)
parker.SetTargetDistance(channel, 50000)  // Full rotation

// Execute
parker.Go(channel)
```

### Incremental Movement
```go
parker.SetIncrementalMode(1)
parker.SetTargetVelocity(1, 2.5)
parker.SetTargetDistance(1, 500)
parker.Go(1)
```

### Continuous Motion
```go
parker.SetContinuosMode(1)
parker.SetTargetVelocity(1, 10.0)
parker.Go(1)

// Motor moves continuously
time.Sleep(5 * time.Second)

parker.Stop(1)
```

### Status Monitoring
```go
status, err := parker.GetIndexerStatus(1)
if err != nil {
    log.Fatal(err)
}

if status == protocol.IndexerBusy {
    fmt.Println("Motor is moving...")
}

limits, _ := parker.GetLimitsStatus(1)
fmt.Printf("Limit status: %s\n", limits)

position, _ := parker.GetAbsolutePosition(1)
fmt.Printf("Current position: %d steps\n", position)
```

### Multi-Motor Control
```go
// Configure multiple motors
for channel := uint(1); channel <= 3; channel++ {
    parker.SetAbsoluteMode(channel)
    parker.SetTargetVelocity(channel, 5.0)
    parker.SetTargetDistance(channel, 1000)
}

// Move all motors simultaneously
parker.GoAll()
```

### Error Handling Best Practices
```go
if err := parker.Connect(); err != nil {
    log.Fatalf("Failed to connect: %v", err)
}
defer func() {
    if err := parker.Disconnect(); err != nil {
        log.Printf("Error disconnecting: %v", err)
    }
}()

// Check connection before operations
if !parker.IsConnected() {
    log.Fatal("Device not connected")
}

// Handle command errors
if err := parker.SetTargetVelocity(1, 5.0); err != nil {
    log.Printf("Failed to set velocity: %v", err)
    return
}
```

## License

This project is authored by Leonardo Rossi Leao and was created on November 26th, 2025.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.