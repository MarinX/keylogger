# Keylogger

Capture global keyboard events on Linux

[![Build Status](https://travis-ci.org/MarinX/keylogger.svg?branch=master)](https://travis-ci.org/MarinX/keylogger)
[![GoDoc](https://godoc.org/github.com/MarinX/keylogger?status.svg)](https://godoc.org/github.com/MarinX/keylogger)
[![License MIT](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat)](LICENSE)

## Notes
* Only Linux based
* Need root privilages

## Installation
```sh
go get github.com/MarinX/keylogger
```

## Getting started

### Finding keyboard device
There is a helper on finding the keyboard.
```go
 keyboard := keylogger.FindKeyboardDevice()
```
Which goes through each file device name to find keyword "keyboard"
```sh
/sys/class/input/event[0-255]/device/name
```
and returns the file event path if found
```sh
/dev/input/event2
```
If the function returns empty string, you will need to cat each device name and get the event number.
If you know already, you can easily pass it to constructor
```sh
keylogger.New("/dev/input/event2")
```

### Getting keypress
Once the keylogger returns channel event, you can switch by event code as described in [input_event.go](https://github.com/MarinX/keylogger/blob/master/input_event.go)
For start, you can listen on keyboard state change 
```go
keylogger.EvKey
```
Once you get desire event, there is a helper to parse code into human readable key.
```go
event.KeyString()
```
**NOTE**

If you listen on keyboard state change, it will return __double__ results.
This is because pressing and releasing the key are 2 different state change.
There is a helper function which you can call to see which type of state change happend
```go
// returns true if key on keyboard is pressed
event.KeyPress()

// returns true if key on keyboard is released
event.KeyRelease()
```

### Example
You can find a example script in ```example/main.go```

### Running tests
No magic, just run
```sh
go test -v
```

## Creating key sniffer (needs update)
* [sniffing global keyboard events in go](https://medium.com/@marin.basic02/sniffing-global-keyboard-events-in-go-e5497e618192/)


## License
This library is under the MIT License
