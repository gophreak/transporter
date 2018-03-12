package filters

import (
	"path/filepath"
	"plugin"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basePath   = filepath.Dir(b)
	libPath    = basePath + "/plugin"
)

// ProfanityCheck interface for checking for profanities
type ProfanityCheck interface {
	Check(input string) (bool, error)
}

// Profanity structure which holds the underlying gateway to plugin
type Profanity struct {
	checker ProfanityCheck
}

// NewProfanityFilter setups up a new instance of a profanity checker using the configured plugin
func NewProfanityFilter() *Profanity {
	p := new(Profanity)
	plug, err := plugin.Open(libPath + "/profanity.so")
	if err != nil {
		panic("Cannot find plugin: " + err.Error())
	}

	profane, e := plug.Lookup("Profanity")
	if e != nil {
		panic(e)
	}
	p.checker = profane.(ProfanityCheck)

	return p
}

// Check performs a check if any part of the input is profound. Will return true if any words are detected
func (p Profanity) Check(input string) bool {
	profound, fail := p.checker.Check(input)
	if fail != nil {
		panic("Unable to run profanity check")
	}

	return profound
}
