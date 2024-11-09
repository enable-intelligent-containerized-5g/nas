// Code generated by generate.sh, DO NOT EDIT.

package nas

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/enable-intelligent-containerized-5g/nas/nasMessage"
)

type GmmMessage struct {
	GmmHeader
	*nasMessage.AuthenticationRequest                            // 8.2.1
	*nasMessage.AuthenticationResponse                           // 8.2.2
	*nasMessage.AuthenticationResult                             // 8.2.3
	*nasMessage.AuthenticationFailure                            // 8.2.4
	*nasMessage.AuthenticationReject                             // 8.2.5
	*nasMessage.RegistrationRequest                              // 8.2.6
	*nasMessage.RegistrationAccept                               // 8.2.7
	*nasMessage.RegistrationComplete                             // 8.2.8
	*nasMessage.RegistrationReject                               // 8.2.9
	*nasMessage.ULNASTransport                                   // 8.2.10
	*nasMessage.DLNASTransport                                   // 8.2.11
	*nasMessage.DeregistrationRequestUEOriginatingDeregistration // 8.2.12
	*nasMessage.DeregistrationAcceptUEOriginatingDeregistration  // 8.2.13
	*nasMessage.DeregistrationRequestUETerminatedDeregistration  // 8.2.14
	*nasMessage.DeregistrationAcceptUETerminatedDeregistration   // 8.2.15
	*nasMessage.ServiceRequest                                   // 8.2.16
	*nasMessage.ServiceAccept                                    // 8.2.17
	*nasMessage.ServiceReject                                    // 8.2.18
	*nasMessage.ConfigurationUpdateCommand                       // 8.2.19
	*nasMessage.ConfigurationUpdateComplete                      // 8.2.20
	*nasMessage.IdentityRequest                                  // 8.2.21
	*nasMessage.IdentityResponse                                 // 8.2.22
	*nasMessage.Notification                                     // 8.2.23
	*nasMessage.NotificationResponse                             // 8.2.24
	*nasMessage.SecurityModeCommand                              // 8.2.25
	*nasMessage.SecurityModeComplete                             // 8.2.26
	*nasMessage.SecurityModeReject                               // 8.2.27
	*nasMessage.SecurityProtected5GSNASMessage                   // 8.2.28
	*nasMessage.Status5GMM                                       // 8.2.29
}

const (
	MsgTypeRegistrationRequest                              uint8 = 65
	MsgTypeRegistrationAccept                               uint8 = 66
	MsgTypeRegistrationComplete                             uint8 = 67
	MsgTypeRegistrationReject                               uint8 = 68
	MsgTypeDeregistrationRequestUEOriginatingDeregistration uint8 = 69
	MsgTypeDeregistrationAcceptUEOriginatingDeregistration  uint8 = 70
	MsgTypeDeregistrationRequestUETerminatedDeregistration  uint8 = 71
	MsgTypeDeregistrationAcceptUETerminatedDeregistration   uint8 = 72
	MsgTypeServiceRequest                                   uint8 = 76
	MsgTypeServiceReject                                    uint8 = 77
	MsgTypeServiceAccept                                    uint8 = 78
	MsgTypeConfigurationUpdateCommand                       uint8 = 84
	MsgTypeConfigurationUpdateComplete                      uint8 = 85
	MsgTypeAuthenticationRequest                            uint8 = 86
	MsgTypeAuthenticationResponse                           uint8 = 87
	MsgTypeAuthenticationReject                             uint8 = 88
	MsgTypeAuthenticationFailure                            uint8 = 89
	MsgTypeAuthenticationResult                             uint8 = 90
	MsgTypeIdentityRequest                                  uint8 = 91
	MsgTypeIdentityResponse                                 uint8 = 92
	MsgTypeSecurityModeCommand                              uint8 = 93
	MsgTypeSecurityModeComplete                             uint8 = 94
	MsgTypeSecurityModeReject                               uint8 = 95
	MsgTypeStatus5GMM                                       uint8 = 100
	MsgTypeNotification                                     uint8 = 101
	MsgTypeNotificationResponse                             uint8 = 102
	MsgTypeULNASTransport                                   uint8 = 103
	MsgTypeDLNASTransport                                   uint8 = 104
)

func (a *Message) GmmMessageDecode(byteArray *[]byte) error {
	buffer := bytes.NewBuffer(*byteArray)
	a.GmmMessage = NewGmmMessage()
	if err := binary.Read(buffer, binary.BigEndian, &a.GmmMessage.GmmHeader); err != nil {
		return fmt.Errorf("GMM NAS decode Fail: read fail - %+v", err)
	}
	switch a.GmmMessage.GmmHeader.GetMessageType() {
	case MsgTypeRegistrationRequest:
		a.GmmMessage.RegistrationRequest = nasMessage.NewRegistrationRequest(MsgTypeRegistrationRequest)
		return a.GmmMessage.DecodeRegistrationRequest(byteArray)
	case MsgTypeRegistrationAccept:
		a.GmmMessage.RegistrationAccept = nasMessage.NewRegistrationAccept(MsgTypeRegistrationAccept)
		return a.GmmMessage.DecodeRegistrationAccept(byteArray)
	case MsgTypeRegistrationComplete:
		a.GmmMessage.RegistrationComplete = nasMessage.NewRegistrationComplete(MsgTypeRegistrationComplete)
		return a.GmmMessage.DecodeRegistrationComplete(byteArray)
	case MsgTypeRegistrationReject:
		a.GmmMessage.RegistrationReject = nasMessage.NewRegistrationReject(MsgTypeRegistrationReject)
		return a.GmmMessage.DecodeRegistrationReject(byteArray)
	case MsgTypeDeregistrationRequestUEOriginatingDeregistration:
		a.GmmMessage.DeregistrationRequestUEOriginatingDeregistration = nasMessage.NewDeregistrationRequestUEOriginatingDeregistration(MsgTypeDeregistrationRequestUEOriginatingDeregistration)
		return a.GmmMessage.DecodeDeregistrationRequestUEOriginatingDeregistration(byteArray)
	case MsgTypeDeregistrationAcceptUEOriginatingDeregistration:
		a.GmmMessage.DeregistrationAcceptUEOriginatingDeregistration = nasMessage.NewDeregistrationAcceptUEOriginatingDeregistration(MsgTypeDeregistrationAcceptUEOriginatingDeregistration)
		return a.GmmMessage.DecodeDeregistrationAcceptUEOriginatingDeregistration(byteArray)
	case MsgTypeDeregistrationRequestUETerminatedDeregistration:
		a.GmmMessage.DeregistrationRequestUETerminatedDeregistration = nasMessage.NewDeregistrationRequestUETerminatedDeregistration(MsgTypeDeregistrationRequestUETerminatedDeregistration)
		return a.GmmMessage.DecodeDeregistrationRequestUETerminatedDeregistration(byteArray)
	case MsgTypeDeregistrationAcceptUETerminatedDeregistration:
		a.GmmMessage.DeregistrationAcceptUETerminatedDeregistration = nasMessage.NewDeregistrationAcceptUETerminatedDeregistration(MsgTypeDeregistrationAcceptUETerminatedDeregistration)
		return a.GmmMessage.DecodeDeregistrationAcceptUETerminatedDeregistration(byteArray)
	case MsgTypeServiceRequest:
		a.GmmMessage.ServiceRequest = nasMessage.NewServiceRequest(MsgTypeServiceRequest)
		return a.GmmMessage.DecodeServiceRequest(byteArray)
	case MsgTypeServiceReject:
		a.GmmMessage.ServiceReject = nasMessage.NewServiceReject(MsgTypeServiceReject)
		return a.GmmMessage.DecodeServiceReject(byteArray)
	case MsgTypeServiceAccept:
		a.GmmMessage.ServiceAccept = nasMessage.NewServiceAccept(MsgTypeServiceAccept)
		return a.GmmMessage.DecodeServiceAccept(byteArray)
	case MsgTypeConfigurationUpdateCommand:
		a.GmmMessage.ConfigurationUpdateCommand = nasMessage.NewConfigurationUpdateCommand(MsgTypeConfigurationUpdateCommand)
		return a.GmmMessage.DecodeConfigurationUpdateCommand(byteArray)
	case MsgTypeConfigurationUpdateComplete:
		a.GmmMessage.ConfigurationUpdateComplete = nasMessage.NewConfigurationUpdateComplete(MsgTypeConfigurationUpdateComplete)
		return a.GmmMessage.DecodeConfigurationUpdateComplete(byteArray)
	case MsgTypeAuthenticationRequest:
		a.GmmMessage.AuthenticationRequest = nasMessage.NewAuthenticationRequest(MsgTypeAuthenticationRequest)
		return a.GmmMessage.DecodeAuthenticationRequest(byteArray)
	case MsgTypeAuthenticationResponse:
		a.GmmMessage.AuthenticationResponse = nasMessage.NewAuthenticationResponse(MsgTypeAuthenticationResponse)
		return a.GmmMessage.DecodeAuthenticationResponse(byteArray)
	case MsgTypeAuthenticationReject:
		a.GmmMessage.AuthenticationReject = nasMessage.NewAuthenticationReject(MsgTypeAuthenticationReject)
		return a.GmmMessage.DecodeAuthenticationReject(byteArray)
	case MsgTypeAuthenticationFailure:
		a.GmmMessage.AuthenticationFailure = nasMessage.NewAuthenticationFailure(MsgTypeAuthenticationFailure)
		return a.GmmMessage.DecodeAuthenticationFailure(byteArray)
	case MsgTypeAuthenticationResult:
		a.GmmMessage.AuthenticationResult = nasMessage.NewAuthenticationResult(MsgTypeAuthenticationResult)
		return a.GmmMessage.DecodeAuthenticationResult(byteArray)
	case MsgTypeIdentityRequest:
		a.GmmMessage.IdentityRequest = nasMessage.NewIdentityRequest(MsgTypeIdentityRequest)
		return a.GmmMessage.DecodeIdentityRequest(byteArray)
	case MsgTypeIdentityResponse:
		a.GmmMessage.IdentityResponse = nasMessage.NewIdentityResponse(MsgTypeIdentityResponse)
		return a.GmmMessage.DecodeIdentityResponse(byteArray)
	case MsgTypeSecurityModeCommand:
		a.GmmMessage.SecurityModeCommand = nasMessage.NewSecurityModeCommand(MsgTypeSecurityModeCommand)
		return a.GmmMessage.DecodeSecurityModeCommand(byteArray)
	case MsgTypeSecurityModeComplete:
		a.GmmMessage.SecurityModeComplete = nasMessage.NewSecurityModeComplete(MsgTypeSecurityModeComplete)
		return a.GmmMessage.DecodeSecurityModeComplete(byteArray)
	case MsgTypeSecurityModeReject:
		a.GmmMessage.SecurityModeReject = nasMessage.NewSecurityModeReject(MsgTypeSecurityModeReject)
		return a.GmmMessage.DecodeSecurityModeReject(byteArray)
	case MsgTypeStatus5GMM:
		a.GmmMessage.Status5GMM = nasMessage.NewStatus5GMM(MsgTypeStatus5GMM)
		return a.GmmMessage.DecodeStatus5GMM(byteArray)
	case MsgTypeNotification:
		a.GmmMessage.Notification = nasMessage.NewNotification(MsgTypeNotification)
		return a.GmmMessage.DecodeNotification(byteArray)
	case MsgTypeNotificationResponse:
		a.GmmMessage.NotificationResponse = nasMessage.NewNotificationResponse(MsgTypeNotificationResponse)
		return a.GmmMessage.DecodeNotificationResponse(byteArray)
	case MsgTypeULNASTransport:
		a.GmmMessage.ULNASTransport = nasMessage.NewULNASTransport(MsgTypeULNASTransport)
		return a.GmmMessage.DecodeULNASTransport(byteArray)
	case MsgTypeDLNASTransport:
		a.GmmMessage.DLNASTransport = nasMessage.NewDLNASTransport(MsgTypeDLNASTransport)
		return a.GmmMessage.DecodeDLNASTransport(byteArray)
	default:
		return fmt.Errorf("NAS decode Fail: MsgType[%d] doesn't exist in GMM Message",
			a.GmmMessage.GmmHeader.GetMessageType())
	}
}

func (a *Message) GmmMessageEncode(buffer *bytes.Buffer) error {
	switch a.GmmMessage.GmmHeader.GetMessageType() {
	case MsgTypeRegistrationRequest:
		return a.GmmMessage.EncodeRegistrationRequest(buffer)
	case MsgTypeRegistrationAccept:
		return a.GmmMessage.EncodeRegistrationAccept(buffer)
	case MsgTypeRegistrationComplete:
		return a.GmmMessage.EncodeRegistrationComplete(buffer)
	case MsgTypeRegistrationReject:
		return a.GmmMessage.EncodeRegistrationReject(buffer)
	case MsgTypeDeregistrationRequestUEOriginatingDeregistration:
		return a.GmmMessage.EncodeDeregistrationRequestUEOriginatingDeregistration(buffer)
	case MsgTypeDeregistrationAcceptUEOriginatingDeregistration:
		return a.GmmMessage.EncodeDeregistrationAcceptUEOriginatingDeregistration(buffer)
	case MsgTypeDeregistrationRequestUETerminatedDeregistration:
		return a.GmmMessage.EncodeDeregistrationRequestUETerminatedDeregistration(buffer)
	case MsgTypeDeregistrationAcceptUETerminatedDeregistration:
		return a.GmmMessage.EncodeDeregistrationAcceptUETerminatedDeregistration(buffer)
	case MsgTypeServiceRequest:
		return a.GmmMessage.EncodeServiceRequest(buffer)
	case MsgTypeServiceReject:
		return a.GmmMessage.EncodeServiceReject(buffer)
	case MsgTypeServiceAccept:
		return a.GmmMessage.EncodeServiceAccept(buffer)
	case MsgTypeConfigurationUpdateCommand:
		return a.GmmMessage.EncodeConfigurationUpdateCommand(buffer)
	case MsgTypeConfigurationUpdateComplete:
		return a.GmmMessage.EncodeConfigurationUpdateComplete(buffer)
	case MsgTypeAuthenticationRequest:
		return a.GmmMessage.EncodeAuthenticationRequest(buffer)
	case MsgTypeAuthenticationResponse:
		return a.GmmMessage.EncodeAuthenticationResponse(buffer)
	case MsgTypeAuthenticationReject:
		return a.GmmMessage.EncodeAuthenticationReject(buffer)
	case MsgTypeAuthenticationFailure:
		return a.GmmMessage.EncodeAuthenticationFailure(buffer)
	case MsgTypeAuthenticationResult:
		return a.GmmMessage.EncodeAuthenticationResult(buffer)
	case MsgTypeIdentityRequest:
		return a.GmmMessage.EncodeIdentityRequest(buffer)
	case MsgTypeIdentityResponse:
		return a.GmmMessage.EncodeIdentityResponse(buffer)
	case MsgTypeSecurityModeCommand:
		return a.GmmMessage.EncodeSecurityModeCommand(buffer)
	case MsgTypeSecurityModeComplete:
		return a.GmmMessage.EncodeSecurityModeComplete(buffer)
	case MsgTypeSecurityModeReject:
		return a.GmmMessage.EncodeSecurityModeReject(buffer)
	case MsgTypeStatus5GMM:
		return a.GmmMessage.EncodeStatus5GMM(buffer)
	case MsgTypeNotification:
		return a.GmmMessage.EncodeNotification(buffer)
	case MsgTypeNotificationResponse:
		return a.GmmMessage.EncodeNotificationResponse(buffer)
	case MsgTypeULNASTransport:
		return a.GmmMessage.EncodeULNASTransport(buffer)
	case MsgTypeDLNASTransport:
		return a.GmmMessage.EncodeDLNASTransport(buffer)
	default:
		return fmt.Errorf("NAS Encode Fail: MsgType[%d] doesn't exist in GMM Message",
			a.GmmMessage.GmmHeader.GetMessageType())
	}
}

type GsmMessage struct {
	GsmHeader
	*nasMessage.PDUSessionEstablishmentRequest      // 8.3.1
	*nasMessage.PDUSessionEstablishmentAccept       // 8.3.2
	*nasMessage.PDUSessionEstablishmentReject       // 8.3.3
	*nasMessage.PDUSessionAuthenticationCommand     // 8.3.4
	*nasMessage.PDUSessionAuthenticationComplete    // 8.3.5
	*nasMessage.PDUSessionAuthenticationResult      // 8.3.6
	*nasMessage.PDUSessionModificationRequest       // 8.3.7
	*nasMessage.PDUSessionModificationReject        // 8.3.8
	*nasMessage.PDUSessionModificationCommand       // 8.3.9
	*nasMessage.PDUSessionModificationComplete      // 8.3.10
	*nasMessage.PDUSessionModificationCommandReject // 8.3.11
	*nasMessage.PDUSessionReleaseRequest            // 8.3.12
	*nasMessage.PDUSessionReleaseReject             // 8.3.13
	*nasMessage.PDUSessionReleaseCommand            // 8.3.14
	*nasMessage.PDUSessionReleaseComplete           // 8.3.15
	*nasMessage.Status5GSM                          // 8.3.16
}

const (
	MsgTypePDUSessionEstablishmentRequest      uint8 = 193
	MsgTypePDUSessionEstablishmentAccept       uint8 = 194
	MsgTypePDUSessionEstablishmentReject       uint8 = 195
	MsgTypePDUSessionAuthenticationCommand     uint8 = 197
	MsgTypePDUSessionAuthenticationComplete    uint8 = 198
	MsgTypePDUSessionAuthenticationResult      uint8 = 199
	MsgTypePDUSessionModificationRequest       uint8 = 201
	MsgTypePDUSessionModificationReject        uint8 = 202
	MsgTypePDUSessionModificationCommand       uint8 = 203
	MsgTypePDUSessionModificationComplete      uint8 = 204
	MsgTypePDUSessionModificationCommandReject uint8 = 205
	MsgTypePDUSessionReleaseRequest            uint8 = 209
	MsgTypePDUSessionReleaseReject             uint8 = 210
	MsgTypePDUSessionReleaseCommand            uint8 = 211
	MsgTypePDUSessionReleaseComplete           uint8 = 212
	MsgTypeStatus5GSM                          uint8 = 214
)

func (a *Message) GsmMessageDecode(byteArray *[]byte) error {
	buffer := bytes.NewBuffer(*byteArray)
	a.GsmMessage = NewGsmMessage()
	if err := binary.Read(buffer, binary.BigEndian, &a.GsmMessage.GsmHeader); err != nil {
		return fmt.Errorf("GSM NAS decode Fail: read fail - %+v", err)
	}
	switch a.GsmMessage.GsmHeader.GetMessageType() {
	case MsgTypePDUSessionEstablishmentRequest:
		a.GsmMessage.PDUSessionEstablishmentRequest = nasMessage.NewPDUSessionEstablishmentRequest(MsgTypePDUSessionEstablishmentRequest)
		return a.GsmMessage.DecodePDUSessionEstablishmentRequest(byteArray)
	case MsgTypePDUSessionEstablishmentAccept:
		a.GsmMessage.PDUSessionEstablishmentAccept = nasMessage.NewPDUSessionEstablishmentAccept(MsgTypePDUSessionEstablishmentAccept)
		return a.GsmMessage.DecodePDUSessionEstablishmentAccept(byteArray)
	case MsgTypePDUSessionEstablishmentReject:
		a.GsmMessage.PDUSessionEstablishmentReject = nasMessage.NewPDUSessionEstablishmentReject(MsgTypePDUSessionEstablishmentReject)
		return a.GsmMessage.DecodePDUSessionEstablishmentReject(byteArray)
	case MsgTypePDUSessionAuthenticationCommand:
		a.GsmMessage.PDUSessionAuthenticationCommand = nasMessage.NewPDUSessionAuthenticationCommand(MsgTypePDUSessionAuthenticationCommand)
		return a.GsmMessage.DecodePDUSessionAuthenticationCommand(byteArray)
	case MsgTypePDUSessionAuthenticationComplete:
		a.GsmMessage.PDUSessionAuthenticationComplete = nasMessage.NewPDUSessionAuthenticationComplete(MsgTypePDUSessionAuthenticationComplete)
		return a.GsmMessage.DecodePDUSessionAuthenticationComplete(byteArray)
	case MsgTypePDUSessionAuthenticationResult:
		a.GsmMessage.PDUSessionAuthenticationResult = nasMessage.NewPDUSessionAuthenticationResult(MsgTypePDUSessionAuthenticationResult)
		return a.GsmMessage.DecodePDUSessionAuthenticationResult(byteArray)
	case MsgTypePDUSessionModificationRequest:
		a.GsmMessage.PDUSessionModificationRequest = nasMessage.NewPDUSessionModificationRequest(MsgTypePDUSessionModificationRequest)
		return a.GsmMessage.DecodePDUSessionModificationRequest(byteArray)
	case MsgTypePDUSessionModificationReject:
		a.GsmMessage.PDUSessionModificationReject = nasMessage.NewPDUSessionModificationReject(MsgTypePDUSessionModificationReject)
		return a.GsmMessage.DecodePDUSessionModificationReject(byteArray)
	case MsgTypePDUSessionModificationCommand:
		a.GsmMessage.PDUSessionModificationCommand = nasMessage.NewPDUSessionModificationCommand(MsgTypePDUSessionModificationCommand)
		return a.GsmMessage.DecodePDUSessionModificationCommand(byteArray)
	case MsgTypePDUSessionModificationComplete:
		a.GsmMessage.PDUSessionModificationComplete = nasMessage.NewPDUSessionModificationComplete(MsgTypePDUSessionModificationComplete)
		return a.GsmMessage.DecodePDUSessionModificationComplete(byteArray)
	case MsgTypePDUSessionModificationCommandReject:
		a.GsmMessage.PDUSessionModificationCommandReject = nasMessage.NewPDUSessionModificationCommandReject(MsgTypePDUSessionModificationCommandReject)
		return a.GsmMessage.DecodePDUSessionModificationCommandReject(byteArray)
	case MsgTypePDUSessionReleaseRequest:
		a.GsmMessage.PDUSessionReleaseRequest = nasMessage.NewPDUSessionReleaseRequest(MsgTypePDUSessionReleaseRequest)
		return a.GsmMessage.DecodePDUSessionReleaseRequest(byteArray)
	case MsgTypePDUSessionReleaseReject:
		a.GsmMessage.PDUSessionReleaseReject = nasMessage.NewPDUSessionReleaseReject(MsgTypePDUSessionReleaseReject)
		return a.GsmMessage.DecodePDUSessionReleaseReject(byteArray)
	case MsgTypePDUSessionReleaseCommand:
		a.GsmMessage.PDUSessionReleaseCommand = nasMessage.NewPDUSessionReleaseCommand(MsgTypePDUSessionReleaseCommand)
		return a.GsmMessage.DecodePDUSessionReleaseCommand(byteArray)
	case MsgTypePDUSessionReleaseComplete:
		a.GsmMessage.PDUSessionReleaseComplete = nasMessage.NewPDUSessionReleaseComplete(MsgTypePDUSessionReleaseComplete)
		return a.GsmMessage.DecodePDUSessionReleaseComplete(byteArray)
	case MsgTypeStatus5GSM:
		a.GsmMessage.Status5GSM = nasMessage.NewStatus5GSM(MsgTypeStatus5GSM)
		return a.GsmMessage.DecodeStatus5GSM(byteArray)
	default:
		return fmt.Errorf("NAS decode Fail: MsgType[%d] doesn't exist in GSM Message",
			a.GsmMessage.GsmHeader.GetMessageType())
	}
}

func (a *Message) GsmMessageEncode(buffer *bytes.Buffer) error {
	switch a.GsmMessage.GsmHeader.GetMessageType() {
	case MsgTypePDUSessionEstablishmentRequest:
		return a.GsmMessage.EncodePDUSessionEstablishmentRequest(buffer)
	case MsgTypePDUSessionEstablishmentAccept:
		return a.GsmMessage.EncodePDUSessionEstablishmentAccept(buffer)
	case MsgTypePDUSessionEstablishmentReject:
		return a.GsmMessage.EncodePDUSessionEstablishmentReject(buffer)
	case MsgTypePDUSessionAuthenticationCommand:
		return a.GsmMessage.EncodePDUSessionAuthenticationCommand(buffer)
	case MsgTypePDUSessionAuthenticationComplete:
		return a.GsmMessage.EncodePDUSessionAuthenticationComplete(buffer)
	case MsgTypePDUSessionAuthenticationResult:
		return a.GsmMessage.EncodePDUSessionAuthenticationResult(buffer)
	case MsgTypePDUSessionModificationRequest:
		return a.GsmMessage.EncodePDUSessionModificationRequest(buffer)
	case MsgTypePDUSessionModificationReject:
		return a.GsmMessage.EncodePDUSessionModificationReject(buffer)
	case MsgTypePDUSessionModificationCommand:
		return a.GsmMessage.EncodePDUSessionModificationCommand(buffer)
	case MsgTypePDUSessionModificationComplete:
		return a.GsmMessage.EncodePDUSessionModificationComplete(buffer)
	case MsgTypePDUSessionModificationCommandReject:
		return a.GsmMessage.EncodePDUSessionModificationCommandReject(buffer)
	case MsgTypePDUSessionReleaseRequest:
		return a.GsmMessage.EncodePDUSessionReleaseRequest(buffer)
	case MsgTypePDUSessionReleaseReject:
		return a.GsmMessage.EncodePDUSessionReleaseReject(buffer)
	case MsgTypePDUSessionReleaseCommand:
		return a.GsmMessage.EncodePDUSessionReleaseCommand(buffer)
	case MsgTypePDUSessionReleaseComplete:
		return a.GsmMessage.EncodePDUSessionReleaseComplete(buffer)
	case MsgTypeStatus5GSM:
		return a.GsmMessage.EncodeStatus5GSM(buffer)
	default:
		return fmt.Errorf("NAS Encode Fail: MsgType[%d] doesn't exist in GSM Message",
			a.GsmMessage.GsmHeader.GetMessageType())
	}
}
