package nasType_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/enable-intelligent-containerized-5g/nas"
	"github.com/enable-intelligent-containerized-5g/nas/nasType"
)

func TestNasTypeNewPDUSESSIONAUTHENTICATIONCOMMANDMessageIdentity(t *testing.T) {
	a := nasType.NewPDUSESSIONAUTHENTICATIONCOMMANDMessageIdentity()
	assert.NotNil(t, a)
}

type nasTypePDUSESSIONAUTHENTICATIONCOMMANDMessageIdentityMessageType struct {
	in  uint8
	out uint8
}

var nasTypePDUSESSIONAUTHENTICATIONCOMMANDMessageIdentityMessageTypeTable = []nasTypePDUSESSIONAUTHENTICATIONCOMMANDMessageIdentityMessageType{
	{nas.MsgTypePDUSessionAuthenticationCommand, nas.MsgTypePDUSessionAuthenticationCommand},
}

func TestNasTypeGetSetPDUSESSIONAUTHENTICATIONCOMMANDMessageIdentityMessageType(t *testing.T) {
	a := nasType.NewPDUSESSIONAUTHENTICATIONCOMMANDMessageIdentity()
	for _, table := range nasTypePDUSESSIONAUTHENTICATIONCOMMANDMessageIdentityMessageTypeTable {
		a.SetMessageType(table.in)
		assert.Equal(t, table.out, a.GetMessageType())
	}
}
