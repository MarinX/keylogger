package keylogger

import (
	"testing"
)

func TestFileDescriptor(t *testing.T) {
	k := &KeyLogger{}

	err := k.Close()
	if err != nil {
		t.Error("Closing empty file descriptor should not yield error", err)
		return
	}
}

func TestBufferParser(t *testing.T) {
	k := &KeyLogger{}

	// keyboard
	input, err := k.eventFromBuffer([]byte{138, 180, 84, 92, 0, 0, 0, 0, 62, 75, 8, 0, 0, 0, 0, 0, 4, 0, 4, 0, 30, 0, 0, 0})
	if err != nil {
		t.Error(err)
		return
	}
	if input == nil {
		t.Error("Event is empty, expected parsed event")
		return
	}

	if input.KeyString() != "3" {
		t.Errorf("wrong input key. got %v, expected %v", input.KeyString(), "3")
		return
	}

	if input.Type != EvMsc {
		t.Errorf("wrong event type. expected key press but got %v", input.Type)
		return
	}
}
