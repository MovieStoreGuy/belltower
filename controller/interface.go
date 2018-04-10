package controller

import "os"

// Handler is the abstraction that allows for nice interacts
// with the underlying implementation
type Handler interface {
	// Register will add the given operation to the signal handler processs
	// and once called will process call functions given
	Register(sig os.Signal, op interface{}) error

	// Cancel will remove the given function from the desired signal
	Cancel(sig os.Signal, op interface{}) error

	// OnError allows the callee to defined how the Handler should
	// handle in errors in its processing
	OnError(handler func(error))

	// Run will start listening on all the signal handler in the background
	Run() chan bool
}
