package messages

import (
	"bytes"
	"encoding/binary"
)

type ServiceType uint8

const (
	ServiceTypeUnacknowledged ServiceType = 0
	ServiceTypeAcknowledged   ServiceType = 1
	ServiceTypeResponse       ServiceType = 2
)

type MessageHeader struct {
	ServiceType    ServiceType
	MessageType    MessageType
	SequenceNumber uint16
}

func ParseMessageHeader(buf *bytes.Buffer) (header MessageHeader, err error) {
	serviceType, err := buf.ReadByte()
	if err != nil {
		return
	}
	header.ServiceType = ServiceType(serviceType)

	messageType, err := buf.ReadByte()
	if err != nil {
		return
	}
	header.MessageType = MessageType(messageType)

	err = binary.Read(buf, binary.BigEndian, &header.SequenceNumber)

	return
}
