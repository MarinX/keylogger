// model
package keylogger

import (
	"syscall"
	"unsafe"
)

const (
	INPUTS        = "/sys/class/input/event%d/device/uevent"
	DEVICE_FILE   = "/dev/input/event%d"
	MAX_FILES     = 255
	MAX_NAME_SIZE = 256
)

//event types
const (
	EV_SYN       = 0x00
	EV_KEY       = 0x01
	EV_REL       = 0x02
	EV_ABS       = 0x03
	EV_MSC       = 0x04
	EV_SW        = 0x05
	EV_LED       = 0x11
	EV_SND       = 0x12
	EV_REP       = 0x14
	EV_FF        = 0x15
	EV_PWR       = 0x16
	EV_FF_STATUS = 0x17
	EV_MAX       = 0x1f
)

var eventsize = int(unsafe.Sizeof(InputEvent{}))

type KeyLogger struct {
	dev *InputDevice
}

type InputDevice struct {
	Id   int
	Name string
}

type InputEvent struct {
	Time  syscall.Timeval
	Type  uint16
	Code  uint16
	Value int32
}
