package messages

import (
	"bytes"
	"encoding/binary"
	"time"
)

type EventReportMessageBody struct {
	UpdateTime time.Time
	TimeOfFix  time.Time
	Latitude   float32
	Longitude  float32
	Speed      float32
	Satellites uint8
	FixStatus  uint8
	EventIndex uint8
	EventCode  uint8
	Accums     uint8
	HDOP       uint8
	Carrier    int16
	RSSI       int16
	Heading    int16
	Altitude   int32
	Inputs     [8]bool
	AccumList  []uint32
	CommState  CommState
	UnitStatus UnitStatus
}

type CommState struct {
	Available         bool
	NetworkService    bool
	DataService       bool
	Connected         bool
	VoiceCallIsActive bool
	Roaming           bool
	NetworkTechnology NetworkTechnology
}

type NetworkTechnology string

const (
	NetworkTechnology2G       NetworkTechnology = "2G Network (CDMA or GSM)"
	NetworkTechnology3G       NetworkTechnology = "3G Network (UMTS)"
	NetworkTechnology4G       NetworkTechnology = "4G Network (LTE)"
	NetworkTechnologyReserved NetworkTechnology = "Reserved"
)

type UnitStatus struct {
	HTTPOTAUpdate       bool
	GPSAntenna          bool
	GPSReceiverSelfTest bool
	GPSReceiverTracking bool
}

func ParseEventReportMessageBody(buf *bytes.Buffer) (body EventReportMessageBody, err error) {
	var updateTime uint32
	if err = binary.Read(buf, binary.BigEndian, &updateTime); err != nil {
		return
	}
	body.UpdateTime = time.Unix(int64(updateTime), 0)

	var timeOfFix uint32
	if err = binary.Read(buf, binary.BigEndian, &timeOfFix); err != nil {
		return
	}
	body.TimeOfFix = time.Unix(int64(updateTime), 0)

	var latitude int32
	if err = binary.Read(buf, binary.BigEndian, &latitude); err != nil {
		return
	}
	body.Latitude = float32(latitude) * 1e-7

	var longitude int32
	if err = binary.Read(buf, binary.BigEndian, &longitude); err != nil {
		return
	}
	body.Longitude = float32(longitude) * 1e-7

	if err = binary.Read(buf, binary.BigEndian, &body.Altitude); err != nil {
		return
	}

	var speed uint32
	if err = binary.Read(buf, binary.BigEndian, &speed); err != nil {
		return
	}
	body.Speed = float32(speed) * 0.036

	if err = binary.Read(buf, binary.BigEndian, &body.Heading); err != nil {
		return
	}

	body.Satellites, err = buf.ReadByte()
	if err != nil {
		return
	}

	body.FixStatus, err = buf.ReadByte()
	if err != nil {
		return
	}

	if err = binary.Read(buf, binary.BigEndian, &body.Carrier); err != nil {
		return
	}

	if err = binary.Read(buf, binary.BigEndian, &body.RSSI); err != nil {
		return
	}

	commState, err := buf.ReadByte()
	if err != nil {
		return
	}

	body.CommState.Available = (commState>>0)&1 == 1
	body.CommState.NetworkService = (commState>>1)&1 == 1
	body.CommState.DataService = (commState>>2)&1 == 1
	body.CommState.Connected = (commState>>3)&1 == 1
	body.CommState.VoiceCallIsActive = (commState>>4)&1 == 1
	body.CommState.Roaming = (commState>>5)&1 == 1
	// msg.CommState.NetworkTechnology = (commState>>6)&1, (commState>>7)&1)

	body.HDOP, err = buf.ReadByte()
	if err != nil {
		return
	}

	inputs, err := buf.ReadByte()
	if err != nil {
		return
	}
	for i := 0; i < 8; i++ {
		body.Inputs[i] = (inputs>>i)&1 == 1
	}

	unitStatus, err := buf.ReadByte()
	if err != nil {
		return
	}

	body.UnitStatus.HTTPOTAUpdate = (unitStatus>>0)&1 == 1
	body.UnitStatus.GPSAntenna = (unitStatus>>1)&1 == 1
	body.UnitStatus.GPSReceiverSelfTest = (unitStatus>>2)&1 == 1
	body.UnitStatus.GPSReceiverTracking = (unitStatus>>3)&1 == 1

	eventIndex, _ := buf.ReadByte()
	if err != nil {
		return
	}
	body.EventIndex = eventIndex

	eventCode, err := buf.ReadByte()
	if err != nil {
		return
	}
	body.EventCode = eventCode

	accums, _ := buf.ReadByte()
	if err != nil {
		return
	}

	// Append
	_, err = buf.ReadByte()
	if err != nil {
		return
	}

	accumCount := accums & 0x3F

	body.Accums = accumCount

	body.AccumList = make([]uint32, accumCount)

	for i := 0; i < int(accumCount); i++ {
		accBuf := make([]byte, 4)
		buf.Read(accBuf)
		body.AccumList[i] = binary.BigEndian.Uint32(accBuf)
	}

	return
}
