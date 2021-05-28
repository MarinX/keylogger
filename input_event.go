package keylogger

import (
	"syscall"
	"unsafe"
)

const (
	// EvSyn is used as markers to separate events. Events may be separated in time or in space, such as with the multitouch protocol.
	EvSyn EventType = 0x00
	// EvKey is used to describe state changes of keyboards, buttons, or other key-like devices.
	EvKey EventType = 0x01
	// EvRel is used to describe relative axis value changes, e.g. moving the mouse 5 units to the left.
	EvRel EventType = 0x02
	// EvAbs is used to describe absolute axis value changes, e.g. describing the coordinates of a touch on a touchscreen.
	EvAbs EventType = 0x03
	// EvMsc is used to describe miscellaneous input data that do not fit into other types.
	EvMsc EventType = 0x04
	// EvSw is used to describe binary state input switches.
	EvSw EventType = 0x05
	// EvLed is used to turn LEDs on devices on and off.
	EvLed EventType = 0x11
	// EvSnd is used to output sound to devices.
	EvSnd EventType = 0x12
	// EvRep is used for autorepeating devices.
	EvRep EventType = 0x14
	// EvFf is used to send force feedback commands to an input device.
	EvFf EventType = 0x15
	// EvPwr is a special type for power button and switch input.
	EvPwr EventType = 0x16
	// EvFfStatus is used to receive force feedback device status.
	EvFfStatus EventType = 0x17
)

// EventType are groupings of codes under a logical input construct.
// Each type has a set of applicable codes to be used in generating events.
// See the Ev section for details on valid codes for each type
type EventType uint16

// eventsize is size of structure of InputEvent
var eventsize = int(unsafe.Sizeof(InputEvent{}))

// InputEvent is the keyboard event structure itself
type InputEvent struct {
	Time  syscall.Timeval
	Type  EventType
	Code  uint16
	Value int32
}

// KeyString returns representation of pressed key as string
// eg enter, space, a, b, c...
func (i *InputEvent) KeyString() string {
	return keyCodeMap[i.Code]
}

// KeyPress is the value when we press the key on keyboard
func (i *InputEvent) KeyPress() bool {
	return i.Value == 1
}

// KeyRelease is the value when we release the key on keyboard
func (i *InputEvent) KeyRelease() bool {
	return i.Value == 0
}

// KeyEvent is the keyboard event for up/down (press/release)
type KeyEvent int32

const (
	KeyPress   KeyEvent = 1
	KeyRelease KeyEvent = 0
)
