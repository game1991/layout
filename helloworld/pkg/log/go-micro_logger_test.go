package log

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMicroLogger(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			assert.Nil(t, err)
		}
	}()
	NewMicroLogger()
}
