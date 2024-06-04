package keylogger

import (
	"os"
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

func TestWithPermission(t *testing.T) {
	fd, err := os.CreateTemp("", "*")
	if err != nil {
		t.Fatal(err)
	}
	// try to create new keylogger with file descriptor which has the permission
	k, err := New(fd.Name())
	if err != nil {
		t.Fatal(err)
	}
	k.Close()
	fd.Close()

	// try to create new keylogger with file descriptor which has no permission
	_, err = New("/dev/tty0")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if err.Error() != "permission denied. run with root permission or use a user with access to /dev/tty0" {
		t.Fatalf("unexpected error: %v", err)
	}
}
