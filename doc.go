// Package keylogger is a go library for linux to capture keyboard events. About:
//
// * No C deps
//
// * Events ported from uapi/linux/input.h
//
// * Needs root privilages
//
// * Transfer keyboard code into human readable key
//
// * Capture state changes
//
// See README at https://github.com/MarinX/keylogger for more info.
package keylogger
