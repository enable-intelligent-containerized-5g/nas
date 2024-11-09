package nasType_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/enable-intelligent-containerized-5g/nas"
	"github.com/enable-intelligent-containerized-5g/nas/nasType"
)

type nasTypeDeregistrationAcceptMessageIdentityData struct {
	in  uint8
	out uint8
}

var nasTypeDeregistrationAcceptMessageIdentityTable = []nasTypeDeregistrationAcceptMessageIdentityData{
	{nas.MsgTypeDeregistrationAcceptUETerminatedDeregistration, nas.MsgTypeDeregistrationAcceptUETerminatedDeregistration},
}

func TestNasTypeNewDeregistrationAcceptMessageIdentity(t *testing.T) {
	a := nasType.NewDeregistrationAcceptMessageIdentity()
	assert.NotNil(t, a)
}

func TestNasTypeGetSetDeregistrationAcceptMessageIdentity(t *testing.T) {
	a := nasType.NewDeregistrationAcceptMessageIdentity()
	for _, table := range nasTypeDeregistrationAcceptMessageIdentityTable {
		a.SetMessageType(table.in)
		assert.Equal(t, table.out, a.GetMessageType())
	}
}
