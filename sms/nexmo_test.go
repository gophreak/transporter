package sms_test

import (
	"testing"

	"github.com/mikeyscode/nexmo"
	"github.com/stretchr/testify/assert"

	"github.com/gophreak/transporter"
	"github.com/gophreak/transporter/sms"
)

const (
	nexmoSMSTypeName = "sms"

	nexmoSMSAdaptorName = "nexmo"

	nexmoValidMobile = "+447902040506"

	nexmoInvalidMobile = "+a47902040506"
)

func TestNewNexmoAdaptor(t *testing.T) {
	nexmo := sms.NewNexmoAdaptor(MockNexmo{})

	assert.IsType(t, &sms.Nexmo{}, nexmo)
}

func TestNexmo_GetType(t *testing.T) {
	nexmo := sms.NewNexmoAdaptor(MockNexmo{})

	assert.Equal(t, nexmoSMSTypeName, nexmo.GetType())
}

func TestNexmo_GetAdaptorName(t *testing.T) {
	nexmo := sms.NewNexmoAdaptor(MockNexmo{})

	assert.Equal(t, nexmoSMSAdaptorName, nexmo.GetAdaptorName())
}

func TestNexmo_Setup(t *testing.T) {
	nexmo := sms.NewNexmoAdaptor(MockNexmo{})

	var config []byte

	e := nexmo.Setup(config)

	assert.NoError(t, e)
}

func TestNexmo_Push(t *testing.T) {
	nexmo := sms.NewNexmoAdaptor(MockNexmo{})

	e := nexmo.Push("0", nexmoValidMobile, []byte("Hello, World!"), transporter.Table{})

	assert.NoError(t, e)
}

func TestNexmo_PushInvalidNumber(t *testing.T) {
	nexmo := sms.NewNexmoAdaptor(MockNexmo{})

	e := nexmo.Push("0", nexmoInvalidMobile, []byte("Hello, World!"), transporter.Table{})

	assert.Error(t, e)
}

func TestNexmo_PushInvalidText(t *testing.T) {
	nexmo := sms.NewNexmoAdaptor(MockNexmo{})

	e := nexmo.Push("0", "1", []byte("Hello, anus!"), transporter.Table{})

	assert.Error(t, e)
}

// MockNexmoSMS mocks the NexmoSMS interface for overriding the library
type MockNexmo struct{}

// Setup mocks the SMS setup function
func (MockNexmo) Setup(config nexmo.Auth) error { return nil }

// SendSMS mocks the SMS sending functionality
func (MockNexmo) SendSMS(from, to string, options nexmo.SMSOptions) (nexmo.SMSResponseInterface, error) {
	return nil, nil
}

// DispatchTextCall mocks the DispatchTextCall functionality
func (MockNexmo) DispatchTextCall(from, to string, options nexmo.TextCallOptions) (nexmo.TextCallResponseInterface, error) {
	return nil, nil
}
