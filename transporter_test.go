package transporter_test

import (
	"testing"

	"github.com/gophreak/transporter"
	"github.com/stretchr/testify/assert"
)

const (
	testTransportMechanism1 = "testingReg"

	testTransportMechanism2 = "testingUnreg"

	testTransportMechanism3 = "testingAdaptor"

	testTransportAdaptor1 = "inMemory"

	testTransportAdaptor = "inMemory2"
)

func TestRegister(t *testing.T) {
	// initial registration
	noErr := transporter.Register(testTransportMechanism1, testTransportAdaptor1, TestNotification{})
	assert.NoError(t, noErr)

	// register same mechanism, different adaptor
	noErr2 := transporter.Register(testTransportMechanism1, testTransportAdaptor, TestNotification{})
	assert.NoError(t, noErr2)

	err := transporter.Register(testTransportMechanism1, testTransportAdaptor1, TestNotification{})
	assert.Error(t, err)
}

func TestUnRegister(t *testing.T) {
	// initial registration
	noErr := transporter.Register(testTransportMechanism2, testTransportAdaptor1, TestNotification{})
	assert.NoError(t, noErr)

	// unregister so we can reregister
	transporter.Unregister(testTransportMechanism2, testTransportAdaptor1)

	noErr = transporter.Register(testTransportMechanism2, testTransportAdaptor1, TestNotification{})
	assert.NoError(t, noErr)

	// ensure no errors if something is not registered
	transporter.Unregister(testTransportMechanism2, "someRandomString0123")

	// ensure no errors if something is not registered
	transporter.Unregister("someRandomString0123", testTransportAdaptor1)
}

func TestGetAdaptor(t *testing.T) {
	adaptor, err := transporter.GetAdaptor(testTransportMechanism3, testTransportAdaptor1)
	assert.Nil(t, adaptor)
	assert.Error(t, err)

	// initial registration
	noErr := transporter.Register(testTransportMechanism3, testTransportAdaptor1, TestNotification{})
	assert.NoError(t, noErr)

	adaptor, err = transporter.GetAdaptor(testTransportMechanism3, testTransportAdaptor1)
	assert.NotNil(t, adaptor)
	assert.NoError(t, err)
}

type TestNotification struct{}

// GetType returns the type of transport being configured
func (TestNotification) GetType() string { return "" }

// GetAdaptorName returns the name of the adaptor for the type of transport
func (TestNotification) GetAdaptorName() string { return "" }

// Setup configures the adaptor to be able to run the transport
func (TestNotification) Setup(config []byte) error { return nil }

// Push a notification to the configured adaptor for the type of transport. Returns error if something goes wrong
func (TestNotification) Push(string, string, []byte, transporter.Table) error { return nil }
