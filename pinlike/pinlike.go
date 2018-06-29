package pinlike

import (
	"log"

	"github.com/stianeikeland/go-rpio"
)

type PinLike interface {
	High()
	Low()
}

type GpioPinLike struct {
	pin rpio.Pin
}

type StubPinLike struct {
}

func NewGpioPinLike(pinNumber uint8) *GpioPinLike {
	pin := rpio.Pin(pinNumber)

	// Set pin to output mode
	pin.Output()
	pin.Pull(rpio.PullDown)

	return &GpioPinLike{pin}
}

func (p *GpioPinLike) High() {
	p.pin.High()
}

func (p *GpioPinLike) Low() {
	p.pin.Low()
}
func (p *StubPinLike) High() {
	log.Println("High() stub")
}

func (p *StubPinLike) Low() {
	log.Println("Low() stub")
}
