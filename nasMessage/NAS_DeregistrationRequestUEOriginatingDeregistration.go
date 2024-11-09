// Code generated by generate.sh, DO NOT EDIT.

package nasMessage

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/enable-intelligent-containerized-5g/nas/nasType"
)

type DeregistrationRequestUEOriginatingDeregistration struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.SpareHalfOctetAndSecurityHeaderType
	nasType.DeregistrationRequestMessageIdentity
	nasType.NgksiAndDeregistrationType
	nasType.MobileIdentity5GS
}

func NewDeregistrationRequestUEOriginatingDeregistration(iei uint8) (deregistrationRequestUEOriginatingDeregistration *DeregistrationRequestUEOriginatingDeregistration) {
	deregistrationRequestUEOriginatingDeregistration = &DeregistrationRequestUEOriginatingDeregistration{}
	return deregistrationRequestUEOriginatingDeregistration
}

func (a *DeregistrationRequestUEOriginatingDeregistration) EncodeDeregistrationRequestUEOriginatingDeregistration(buffer *bytes.Buffer) error {
	if err := binary.Write(buffer, binary.BigEndian, a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return fmt.Errorf("NAS encode error (DeregistrationRequestUEOriginatingDeregistration/ExtendedProtocolDiscriminator): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return fmt.Errorf("NAS encode error (DeregistrationRequestUEOriginatingDeregistration/SpareHalfOctetAndSecurityHeaderType): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.DeregistrationRequestMessageIdentity.Octet); err != nil {
		return fmt.Errorf("NAS encode error (DeregistrationRequestUEOriginatingDeregistration/DeregistrationRequestMessageIdentity): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.NgksiAndDeregistrationType.Octet); err != nil {
		return fmt.Errorf("NAS encode error (DeregistrationRequestUEOriginatingDeregistration/NgksiAndDeregistrationType): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.MobileIdentity5GS.GetLen()); err != nil {
		return fmt.Errorf("NAS encode error (DeregistrationRequestUEOriginatingDeregistration/MobileIdentity5GS): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.MobileIdentity5GS.Buffer); err != nil {
		return fmt.Errorf("NAS encode error (DeregistrationRequestUEOriginatingDeregistration/MobileIdentity5GS): %w", err)
	}
	return nil
}

func (a *DeregistrationRequestUEOriginatingDeregistration) DecodeDeregistrationRequestUEOriginatingDeregistration(byteArray *[]byte) error {
	buffer := bytes.NewBuffer(*byteArray)
	if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return fmt.Errorf("NAS decode error (DeregistrationRequestUEOriginatingDeregistration/ExtendedProtocolDiscriminator): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return fmt.Errorf("NAS decode error (DeregistrationRequestUEOriginatingDeregistration/SpareHalfOctetAndSecurityHeaderType): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.DeregistrationRequestMessageIdentity.Octet); err != nil {
		return fmt.Errorf("NAS decode error (DeregistrationRequestUEOriginatingDeregistration/DeregistrationRequestMessageIdentity): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.NgksiAndDeregistrationType.Octet); err != nil {
		return fmt.Errorf("NAS decode error (DeregistrationRequestUEOriginatingDeregistration/NgksiAndDeregistrationType): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.MobileIdentity5GS.Len); err != nil {
		return fmt.Errorf("NAS decode error (DeregistrationRequestUEOriginatingDeregistration/MobileIdentity5GS): %w", err)
	}
	if a.MobileIdentity5GS.Len < 4 {
		return fmt.Errorf("invalid ie length (DeregistrationRequestUEOriginatingDeregistration/MobileIdentity5GS): %d", a.MobileIdentity5GS.Len)
	}
	a.MobileIdentity5GS.SetLen(a.MobileIdentity5GS.GetLen())
	if err := binary.Read(buffer, binary.BigEndian, a.MobileIdentity5GS.Buffer); err != nil {
		return fmt.Errorf("NAS decode error (DeregistrationRequestUEOriginatingDeregistration/MobileIdentity5GS): %w", err)
	}
	for buffer.Len() > 0 {
		var ieiN uint8
		var tmpIeiN uint8
		if err := binary.Read(buffer, binary.BigEndian, &ieiN); err != nil {
			return fmt.Errorf("NAS decode error (DeregistrationRequestUEOriginatingDeregistration/iei): %w", err)
		}
		// fmt.Println(ieiN)
		if ieiN >= 0x80 {
			tmpIeiN = (ieiN & 0xf0) >> 4
		} else {
			tmpIeiN = ieiN
		}
		// fmt.Println("type", tmpIeiN)
		switch tmpIeiN {
		default:
		}
	}
	return nil
}
