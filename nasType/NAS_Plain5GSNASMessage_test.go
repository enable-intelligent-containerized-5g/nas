package nasType_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/enable-intelligent-containerized-5g/nas/nasType"
)

func TestNasTypeNewPlain5GSNASMessage(t *testing.T) {
	a := nasType.NewPlain5GSNASMessage()
	assert.NotNil(t, a)
}
