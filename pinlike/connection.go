package pinlike

import (
	"log"

	"github.com/stianeikeland/go-rpio"
)

type Connection interface {
	Open() error
	Close() error
}

type GpioConnection struct {
}

type StubConnection struct {
}

func (p *GpioConnection) Open() error {
	return rpio.Open()
}

func (p *GpioConnection) Close() error {
	return rpio.Close()
}

func (p *StubConnection) Open() error {
	log.Println("Open() stub")
	return nil
}

func (p *StubConnection) Close() error {
	log.Println("Close() stub")
	return nil
}
