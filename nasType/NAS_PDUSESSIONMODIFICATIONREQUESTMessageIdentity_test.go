package nasType_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/enable-intelligent-containerized-5g/nas"
	"github.com/enable-intelligent-containerized-5g/nas/nasType"
)

func TestNasTypeNewPDUSESSIONMODIFICATIONREQUESTMessageIdentity(t *testing.T) {
	a := nasType.NewPDUSESSIONMODIFICATIONREQUESTMessageIdentity()
	assert.NotNil(t, a)
}

type nasTypePDUSESSIONMODIFICATIONREQUESTMessageIdentityMessageType struct {
	in  uint8
	out uint8
}

var nasTypePDUSESSIONMODIFICATIONREQUESTMessageIdentityMessageTypeTable = []nasTypePDUSESSIONMODIFICATIONREQUESTMessageIdentityMessageType{
	{nas.MsgTypePDUSessionModificationRequest, nas.MsgTypePDUSessionModificationRequest},
}

func TestNasTypeGetSetPDUSESSIONMODIFICATIONREQUESTMessageIdentityMessageType(t *testing.T) {
	a := nasType.NewPDUSESSIONMODIFICATIONREQUESTMessageIdentity()
	for _, table := range nasTypePDUSESSIONMODIFICATIONREQUESTMessageIdentityMessageTypeTable {
		a.SetMessageType(table.in)
		assert.Equal(t, table.out, a.GetMessageType())
	}
}
