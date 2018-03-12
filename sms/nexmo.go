package sms

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"

	"github.com/gophreak/transporter"
	"github.com/gophreak/transporter/filters"
	nexmolib "github.com/mikeyscode/nexmo"
)

const (
	// NexmoAdaptor is the name of this adaptor for SMS
	NexmoAdaptor = "nexmo"
)

// Nexmo holds the underlying library adaptor information for dispatching function calls
type Nexmo struct {
	lib nexmolib.APIInterface
}

// NewNexmoAdaptor returns a new instance of nexmo with the underlying library configured
func NewNexmoAdaptor(l nexmolib.APIInterface) *Nexmo {
	n := new(Nexmo)
	n.lib = l

	return n
}

// GetType returns the type of notification being configured
func (nexmo Nexmo) GetType() string {
	return AdaptorType
}

// GetAdaptorName returns the name of the adaptor for the type of notification
func (nexmo Nexmo) GetAdaptorName() string {
	return NexmoAdaptor
}

// Setup configures the adaptor to be able to run the notification
func (nexmo *Nexmo) Setup(raw []byte) error {
	var c nexmolib.Auth
	json.Unmarshal(raw, &c)

	return nexmo.lib.Setup(c)
}

// Push a notification to the configured adaptor for the type of Notification. Returns error if something goes wrong
func (nexmo Nexmo) Push(from string, to string, payload []byte, _ transporter.Table) error {
	if !validatePhoneNumber(to) {
		return fmt.Errorf("invalid number to send to %s", to)
	}
	message := string(payload)

	if !validateText(message) {
		return errors.New("your message contains invalid text")
	}

	r, e := nexmo.lib.SendSMS(from, to, nexmolib.SMSOptions{
		Text: message,
	})

	fmt.Println(r)

	return e
}

func validatePhoneNumber(n string) bool {
	const pattern = `^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`
	re := regexp.MustCompile(pattern)

	return re.MatchString(n)
}

func validateText(text string) bool {
	filter := filters.NewProfanityFilter()

	return !filter.Check(text)
}
