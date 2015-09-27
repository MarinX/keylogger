# KeyLogger

# Description
Capture global keyboard events on Linux

# Installation
    go get github.com/MarinX/keylogger
# Notes
* Only Linux based
* Need root privilages


# Example
    package main

    import (
	    "fmt"
	    "github.com/MarinX/keylogger"
    )

    func main() {
	    devs, err := keylogger.NewDevices()
	    if err != nil {
		    fmt.Println(err)
		    return
	    }

	    for _, val := range devs {
		    fmt.Println("Id->", val.Id, "Device->", val.Name)
	    }

	    //keyboard device file, on your system it will be diffrent!
	    rd := keylogger.NewKeyLogger(devs[3])

	    in, err := rd.Read()
	    if err != nil {
		    fmt.Println(err)
		    return
	    }

	    for i := range in {

		    //we only need keypress
		    if i.Type == keylogger.EV_KEY {
			    fmt.Println(i.KeyString())
		    }
	    }
    }

# Creating key sniffer
* [sniffing global keyboard eventy in go](https://medium.com/@marin.basic02/sniffing-global-keyboard-events-in-go-e5497e618192/)


# License
This library is under the MIT License
# Author
Marin Basic 
