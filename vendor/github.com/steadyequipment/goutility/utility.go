package goutility

import (
	"os"

	"errors"
	"fmt"
)

// EmitOnChan emit values through a channel, drain the channel if we're full
func EmitOnChan(c chan string) func(string) error {
	return func(m string) error {
		select {
		case c <- m:

		default: // No room, we have to drain the channel first
			for i := 0; i < len(c); i++ {
				<-c
			}

			c <- m
		}

		return nil
	}
}

// NewError make a new error using SPrintf
func NewError(format string, parameters ...interface{}) error {

	errorMessage := fmt.Sprintf(format, parameters...)
	return errors.New(errorMessage)
}

// CurrentExecutable return the name of the current executable
func CurrentExecutable() string {
	return os.Args[0]
}
