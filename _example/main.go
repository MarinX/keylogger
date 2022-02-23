package main

import (
	"fmt"
	"time"

	"github.com/Acetolyne/keylogger"
)

//@todo update this file to give better examples, these examples came from the forked repository

func main() {

	// find keyboard device, does not require a root permission
	keyboard := keylogger.FindKeyboardDevice()

	// check if we found a path to keyboard
	if len(keyboard) <= 0 {
		fmt.Println("No keyboard found...you will need to provide manual input path")
		return
	}

	fmt.Println("Found a keyboard at", keyboard)
	// init keylogger with keyboard
	k, err := keylogger.New(keyboard)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer k.Close()

	// write to keyboard example:
	go func() {
		time.Sleep(5 * time.Second)
		// open text editor and focus on it, it should say "marin" and new line will be printed
		keys := []string{"m", "a", "r", "i", "n", "ENTER"}
		for _, key := range keys {
			// write once will simulate keyboard press/release, for long press or release, lookup at Write
			k.WriteOnce(key)
		}
	}()

	events := k.Read()

	// range of events
	for e := range events {
		switch e.Type {
		// EvKey is used to describe state changes of keyboards, buttons, or other key-like devices.
		// check the input_event.go for more events
		case keylogger.EvKey:

			// if the state of key is pressed
			if e.KeyPress() {
				fmt.Println("[event] press key ", e.KeyString())
			}

			// if the state of key is released
			if e.KeyRelease() {
				fmt.Println("[event] release key ", e.KeyString())
			}
		}
	}
}
