package transporter

import "fmt"

// Notification type of request
type Notification interface {
	// GetType returns the type of transport being configured
	GetType() string
	// GetAdaptorName returns the name of the adaptor for the type of transport
	GetAdaptorName() string
	// Setup configures the adaptor to be able to run the transport
	Setup(config []byte) error
	// Push a notification to the configured adaptor for the type of transport. Returns error if something goes wrong
	Push(string, string, []byte, Table) error
}

// Table of additional map information for the Push
type Table map[string][]string

// list of adaptors configured for the factory
var adaptors = make(map[string]map[string]Notification)

// Register allows adaptors to register themselves at runtime
func Register(mechanism, adaptorName string, adaptor Notification) error {
	if _, ok := adaptors[mechanism][adaptorName]; ok {
		return fmt.Errorf("adaptor of type %s with name %s is already registered", mechanism, adaptorName)
	}
	if _, ok := adaptors[mechanism]; !ok {
		adaptors[mechanism] = make(map[string]Notification)
	}
	adaptors[mechanism][adaptorName] = adaptor

	return nil
}

// Unregister an adaptor from the registry
func Unregister(notificationType, adaptorName string) {
	if _, ok := adaptors[notificationType]; !ok {
		return
	}

	delete(adaptors[notificationType], adaptorName)
}

// GetAdaptor from the registry
func GetAdaptor(mechanism, adaptorName string) (Notification, error) {
	if _, ok := adaptors[mechanism][adaptorName]; !ok {
		return nil, fmt.Errorf("adaptor of type %s with name %s is not registered", mechanism, adaptorName)
	}

	return adaptors[mechanism][adaptorName], nil
}
