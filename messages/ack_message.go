package messages

import (
	"bytes"
	"fmt"
)

type AckMessageBody struct {
	MessageType MessageType
	Success     bool
	Reason      string
	AppVersion  string
}

func ParseAckMessageBody(buffer *bytes.Buffer) (body AckMessageBody, err error) {
	messageType, err := buffer.ReadByte()
	if err != nil {
		err = fmt.Errorf("fail to read message type: %w", err)
		return
	}
	body.MessageType = MessageType(messageType)

	ack, err := buffer.ReadByte()
	if err != nil {
		err = fmt.Errorf("fail to read ack value: %w", err)
		return
	}

	if ack > 11 {
		err = fmt.Errorf("invalid ack value: %d", ack)
		return
	}

	body.Success = ack == 0

	body.Reason = ackReasons[ack]

	// Spare
	_, err = buffer.ReadByte()

	if err != nil {
		return
	}

	body.AppVersion = buffer.String()

	return
}

var ackReasons = [11]string{
	"operation successful",
	"operation failed, no reason",
	"operation failed, not a supported message type",
	"operation failed, not a supported operation",
	"operation failed, unable to pass to serial port",
	"operation failed, authentication failure",
	"operation failed, Mobile ID look-up failure",
	"operation failed, non-zero sequence number same as last received message",
	"operation failed, message authentication failure",
	"operation failed, message format failure",
	"operation failed, parameter update failure",
}
