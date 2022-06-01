package messages_test

import (
	"bytes"
	"encoding/hex"
	"lm-direct/messages"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var eventReportTCs = []struct {
	Message string
	Err     error
	Body    messages.EventReportMessageBody
}{
	{
		Message: "830546750157890101010233ea600f1dd7600f1dd70b909a79c4e2ef0e000361a400000139002b11000014ffc40f073f002c2c1000000037d00000105600000000001da3a2000056b9000000000000000001ce635a01ce635a00000000000000000000000000000000000000510000000000000000",
		Err:     nil,
		Body: messages.EventReportMessageBody{
			UpdateTime: time.Date(2021, time.January, 25, 16, 36, 55, 0, time.Local),
			TimeOfFix:  time.Date(2021, time.January, 25, 16, 36, 55, 0, time.Local),
			Latitude:   19.40261,
			Longitude:  -99.17606,
			Altitude:   221604,
			Speed:      11.268,
			Heading:    43,
			Satellites: 17,
			Carrier:    20,
			RSSI:       -60,
			FixStatus:  0,
			EventIndex: 44,
			EventCode:  44,
			Accums:     16,
			HDOP:       7,
			CommState: messages.CommState{
				Available:         true,
				NetworkService:    true,
				DataService:       true,
				Connected:         true,
				VoiceCallIsActive: false,
				Roaming:           false,
				NetworkTechnology: "",
			},
			Inputs: [8]bool{
				true,
				true,
				true,
				true,
				true,
				true,
				false,
			},
			AccumList: []uint32{
				14288,
				4182,
				0,
				1942434,
				22201,
				0,
				0,
				30303066,
				30303066,
				0,
				0,
				0,
				0,
				81,
				0,
				0,
			},
		},
	},
	{
		Message: "8305467502757801010102215661043a7061043a70f1c2cd6fe41bd8520001314e000003f8007f12200005ff930f073f002e2c1000000036a00000106000000000001277d300003b4a0000000000000000012e0947012e094700000000000000000000000000000000000000510000000000000000",
		Err:     nil,
		Body: messages.EventReportMessageBody{
			UpdateTime: time.Date(2021, time.July, 30, 14, 44, 16, 0, time.Local),
			TimeOfFix:  time.Date(2021, time.July, 30, 14, 44, 16, 0, time.Local),
			Latitude:   -23.889166,
			Longitude:  -46.79372,
			Altitude:   78158,
			Speed:      36.576,
			Heading:    127,
			Satellites: 18,
			Carrier:    5,
			RSSI:       -109,
			FixStatus:  32,
			EventIndex: 46,
			EventCode:  44,
			Accums:     16,
			HDOP:       7,
			CommState: messages.CommState{
				Available:         true,
				NetworkService:    true,
				DataService:       true,
				Connected:         true,
				VoiceCallIsActive: false,
				Roaming:           false,
				NetworkTechnology: "",
			},
			Inputs: [8]bool{
				true,
				true,
				true,
				true,
				true,
				true,
				false,
			},
			AccumList: []uint32{
				13984,
				4192,
				0,
				1210323,
				15178,
				0,
				0,
				19794247,
				19794247,
				0,
				0,
				0,
				0,
				81,
				0,
				0,
			},
		},
	},
}

func TestEventReport(t *testing.T) {

	for _, tc := range eventReportTCs {
		h, err := hex.DecodeString(tc.Message)
		assert.Nil(t, err)

		buf := bytes.NewBuffer(h)
		parsedMessage, err := messages.Parse(buf)
		assert.Equal(t, tc.Err, err)
		if err != nil {
			continue
		}

		assert.Equal(t, messages.MessageTypeEventReport, parsedMessage.MessageType)

		message := parsedMessage.Body.(messages.EventReportMessageBody)

		assert.EqualValues(t, tc.Body, message)
	}
}
