package messages_test

import (
	"bytes"
	"encoding/hex"
	"errors"
	"lm-direct/messages"
	"testing"

	"github.com/stretchr/testify/assert"
)

type test struct {
	Message string
	Err     error
	Body    messages.AckMessageBody
}

var tests = []test{
	{
		Message: "83050102030405010102010001000000363067",
		Err:     nil,
		Body: messages.AckMessageBody{
			Success:     true,
			MessageType: messages.MessageTypeNull,
			AppVersion:  "60g",
			Reason:      "operation successful",
		},
	},

	{
		Message: "83050102030405010102010001001100363067",
		Err:     errors.New("invalid ack value: 17"),
		Body:    messages.AckMessageBody{},
	},

	{
		Message: "83050102030405010102010001000100363067",
		Err:     nil,
		Body: messages.AckMessageBody{
			Success:     false,
			MessageType: messages.MessageTypeNull,
			AppVersion:  "60g",
			Reason:      "operation failed, no reason",
		},
	},
	{
		Message: "83050102030405010102010001000200363067",
		Err:     nil,
		Body: messages.AckMessageBody{
			Success:     false,
			MessageType: messages.MessageTypeNull,
			AppVersion:  "60g",
			Reason:      "operation failed, not a supported message type",
		},
	},
	{
		Message: "83050102030405010102010001000300363067",
		Err:     nil,
		Body: messages.AckMessageBody{
			Success:     false,
			MessageType: messages.MessageTypeNull,
			AppVersion:  "60g",
			Reason:      "operation failed, not a supported operation",
		},
	},
	{
		Message: "83050102030405010102010001000300363067",
		Err:     nil,
		Body: messages.AckMessageBody{
			Success:     false,
			MessageType: messages.MessageTypeNull,
			AppVersion:  "60g",
			Reason:      "operation failed, not a supported operation",
		},
	},
	{
		Message: "83050102030405010102010001000400363067",
		Err:     nil,
		Body: messages.AckMessageBody{
			Success:     false,
			MessageType: messages.MessageTypeNull,
			AppVersion:  "60g",
			Reason:      "operation failed, unable to pass to serial port",
		},
	},
	{
		Message: "83050102030405010102010001000500363067",
		Err:     nil,
		Body: messages.AckMessageBody{
			Success:     false,
			MessageType: messages.MessageTypeNull,
			AppVersion:  "60g",
			Reason:      "operation failed, authentication failure",
		},
	},
	{
		Message: "83050102030405010102010001000600363067",
		Err:     nil,
		Body: messages.AckMessageBody{
			Success:     false,
			MessageType: messages.MessageTypeNull,
			AppVersion:  "60g",
			Reason:      "operation failed, Mobile ID look-up failure",
		},
	},
	{
		Message: "83050102030405010102010001000700363067",
		Err:     nil,
		Body: messages.AckMessageBody{
			Success:     false,
			MessageType: messages.MessageTypeNull,
			AppVersion:  "60g",
			Reason:      "operation failed, non-zero sequence number same as last received message",
		},
	},
	{
		Message: "83050102030405010102010001000800363067",
		Err:     nil,
		Body: messages.AckMessageBody{
			Success:     false,
			MessageType: messages.MessageTypeNull,
			AppVersion:  "60g",
			Reason:      "operation failed, message authentication failure",
		},
	},
	{
		Message: "83050102030405010102010001000900363067",
		Err:     nil,
		Body: messages.AckMessageBody{
			Success:     false,
			MessageType: messages.MessageTypeNull,
			AppVersion:  "60g",
			Reason:      "operation failed, message format failure",
		},
	},
	{
		Message: "83050102030405010102010001000A00363067",
		Err:     nil,
		Body: messages.AckMessageBody{
			Success:     false,
			MessageType: messages.MessageTypeNull,
			AppVersion:  "60g",
			Reason:      "operation failed, parameter update failure",
		},
	},
}

func TestAckMessage(t *testing.T) {

	for _, tc := range tests {
		h, err := hex.DecodeString(tc.Message)
		assert.Nil(t, err)

		buf := bytes.NewBuffer(h)
		message, err := messages.Parse(buf)
		assert.Equal(t, tc.Err, err)
		if err != nil {
			continue
		}

		assert.Equal(t, messages.MessageTypeAck, message.MessageType)

		ack := message.Body.(messages.AckMessageBody)

		assert.Equal(t, tc.Body.Success, ack.Success)

		assert.Equal(t, tc.Body.MessageType, ack.MessageType)

		assert.Equal(t, tc.Body.AppVersion, ack.AppVersion)

		assert.Equal(t, tc.Body.Reason, ack.Reason)
	}
}
