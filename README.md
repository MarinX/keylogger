# Keylogger

Capture global keyboard events on Linux

Forked from github.com/MarinX/keylogger
Modified to return correct case of character and other characters accessed by holding down modifier keys such as SHIFT

## Notes
* Only Linux based
* Need root privilages

## Installation
```sh
go get github.com/Acetolyne/keylogger
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
Once the keylogger returns channel event, you can switch by event code as described in [input_event.go](https://github.com/Acetolyne/keylogger/blob/master/input_event.go)
For start, you can listen on keyboard state change 
```go
keylogger.EvKey
```
Once you get desire event, there is a helper to parse code into human readable key.
```go
event.KeyString()
```

### Writing keypress
Best way is to open an text editor and see how keyboard will react
There are 2 methods:
```go
func (k *KeyLogger) WriteOnce(key string) error
```
and
```go
func (k *KeyLogger) Write(direction KeyEvent, key string) error 
```
`WriteOnce` method simulates single key press, eg: press and release letter M

`Write` writes to keyboard and sync the event. 
This will keep the key pressed or released until you call another write with other direction
eg, if the key is "A" and direction is press, on UI, you will see "AAAAA..." until you stop with release

Probably you want to use `WriteOnce` method

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


## License
This library is under the MIT License
