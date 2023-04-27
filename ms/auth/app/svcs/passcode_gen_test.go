package svcs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPasscodeGenerator_Generate(t *testing.T) {
	pg := NewPasscodeGenerator()
	passcode, err := pg.Generate()

	assert.NoError(t, err)
	assert.NotEmpty(t, passcode)
	assert.Equal(t, 6, len(passcode))
}
