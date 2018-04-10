package controller

import (
	"fmt"
	"os"
)

type driver struct {
	registered map[os.Signal][]interface{}
	errhandler func(error)
}

// New will return a ready to go Signal Handler
func New() Handler {
	return &driver{
		registered: map[os.Signal][]interface{}{},
		errhandler: func(e error) { fmt.Println("An error occured:", e) },
	}
}

func (d *driver) Register(sig os.Signal, op interface{}) error {
	d.registered[sig] = append(d.registered[sig], op)
	return nil
}

func (d *driver) Cancel(sig os.Signal, op interface{}) error {
	for index, stored := range d.registerd[sig] {
		if stored == op {
			d.registered[sig] = append(d.registered[sig][:index], d.registered[sig][index+1:]...)
			break
		}
	}
	return nil
}

func (d *driver) OnError(handler func(error)) {
	d.errhandler = handler
}
