package messages

import (
	"bytes"
	"fmt"
)

type Message[T any] struct {
	OptionsHeader
	MessageHeader
	Body T
}

type MessageParser interface {
	Parse(*bytes.Buffer) error
}

func Parse(buf *bytes.Buffer) (message *Message[any], err error) {
	message = &Message[any]{}
	optionsHeader, err := ParseOptionsHeader(buf)
	if err != nil {
		return
	}
	message.OptionsHeader = optionsHeader

	messageHeader, err := ParseMessageHeader(buf)
	if err != nil {
		return
	}
	message.MessageHeader = messageHeader

	switch messageHeader.MessageType {
	case MessageTypeNull:
		message.Body, err = ParseNullMessageBody(buf)
	case MessageTypeAck:
		message.Body, err = ParseAckMessageBody(buf)
	case MessageTypeEventReport:
		message.Body, err = ParseEventReportMessageBody(buf)
	case MessageTypeIdReport:
	case MessageTypeUserData:
	case MessageTypeAppData:
	case MessageTypeConfig:
	case MessageTypeUnitRequest:
	case MessageTypeLocateReport:
	case MessageTypeUserDataAcc:
	case MessageTypeMiniEventReport:
	case MessageTypeMiniUserData:
		return nil, fmt.Errorf("message type not implemented: %v", messageHeader.MessageType)
	default:
		return nil, fmt.Errorf("invalid message type: %v", messageHeader.MessageType)
	}
	return
}
