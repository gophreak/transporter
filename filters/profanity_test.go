package filters_test

import (
	"testing"

	"github.com/gophreak/transporter/filters"
	"github.com/stretchr/testify/assert"
)

func TestNewProfanityFilter(t *testing.T) {
	pf := filters.NewProfanityFilter()

	assert.IsType(t, &filters.Profanity{}, pf)
}

func TestProfanity_Check(t *testing.T) {
	pf := filters.NewProfanityFilter()

	assert.IsType(t, &filters.Profanity{}, pf)

	assert.False(t, pf.Check("Hello, world!"))
}

func TestProfanity_CheckInvalid(t *testing.T) {
	pf := filters.NewProfanityFilter()

	assert.IsType(t, &filters.Profanity{}, pf)

	assert.True(t, pf.Check("Hello, anus!"))
}
