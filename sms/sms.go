package sms

import (
	"github.com/gophreak/transporter"
	nexmolib "github.com/mikeyscode/nexmo"
)

const (
	// AdaptorType holds the type of adaptor being implemented.
	AdaptorType = "sms"
)

// initialise the setup to register nexmo as an SMS adaptor
func init() {
	transporter.Register(AdaptorType, NexmoAdaptor, NewNexmoAdaptor(&nexmolib.Nexmo{}))
}
