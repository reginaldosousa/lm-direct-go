package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"lm-direct-go/messages"
	"log"
	"time"
)

func main() {
	h := "830546750157890101010233ea600f1dd7600f1dd70b909a79c4e2ef0e000361a400000139002b11000014ffc40f073f002c2c1000000037d00000105600000000001da3a2000056b9000000000000000001ce635a01ce635a00000000000000000000000000000000000000510000000000000000"
	decodedBytes, err := hex.DecodeString(h)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%+v\n", decodedBytes)
	buf := bytes.NewBuffer(decodedBytes)

	// optionsHeader, err := messages.ParseOptionsHeader(buf)
	// if err != nil {
	// 	log.Fatalf("fail to parse options header: %s", err)
	// }
	// fmt.Printf("options header %+v\n", optionsHeader)
	// fmt.Println()

	// messageHeader, err := messages.ParseMessageHeader(buf)
	// if err != nil {
	// 	log.Fatalf("fail to parse message header: %s", err)
	// }
	// fmt.Printf("message header %+v\n", messageHeader)
	// fmt.Println()

	message, err := messages.Parse(buf)
	if err != nil {
		log.Fatalf("fail to parse message: %s", err)
	}
	formatted, err := json.MarshalIndent(&message, "", "  ")
	fmt.Printf("%s", formatted)
	return

	// header := Message{}
	// if err := binary.Read(message, binary.BigEndian, &header); err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%+v\n", header)

	options, _ := buf.ReadByte()
	fmt.Printf("options: %08b\n", options)

	for i := 0; i < 8; i++ {
		fmt.Printf("options %+v %+v\n", i, (options>>i)&1)
	}

	mobileIdLength, _ := buf.ReadByte()
	fmt.Printf("mobileIDLength: %+v\n", mobileIdLength)

	b := make([]byte, mobileIdLength)
	buf.Read(b)
	decoded := hex.EncodeToString(b)
	fmt.Printf("mobileID: %+v\n", decoded)

	mobileIdTypeLength, _ := buf.ReadByte()
	fmt.Printf("mobileIdTypeLength: %+v\n", mobileIdTypeLength)

	b = make([]byte, mobileIdTypeLength)
	buf.Read(b)
	decoded = hex.EncodeToString(b)
	fmt.Printf("mobileIdType: %+v\n", decoded)

	serviceType, _ := buf.ReadByte()
	fmt.Printf("serviceType: %+v\n", serviceType)

	messageType, _ := buf.ReadByte()
	fmt.Printf("messageType: %+v\n", messageType)

	sequenceNumber1, _ := buf.ReadByte()
	fmt.Printf("sequenceNumber1: %+v\n", sequenceNumber1)

	sequenceNumber2, _ := buf.ReadByte()
	fmt.Printf("sequenceNumber2: %+v\n", sequenceNumber2)

	b = make([]byte, 4)
	buf.Read(b)
	updateTime := binary.BigEndian.Uint32(b)
	fmt.Printf("UpdateTime: %+v\n", time.Unix(int64(updateTime), 0))

	b = make([]byte, 4)
	buf.Read(b)
	timeOfFix := int64(binary.BigEndian.Uint32(b))
	fmt.Printf("timeOfFix: %+v\n", time.Unix(timeOfFix, 0))

	b = make([]byte, 4)
	buf.Read(b)
	latitude := int32(binary.BigEndian.Uint32(b))
	fmt.Printf("latitude: %+v\n", float32(latitude)*1e-7)

	b = make([]byte, 4)
	buf.Read(b)
	longitude := int32(binary.BigEndian.Uint32(b))
	fmt.Printf("longitude: %+v\n", float32(longitude)*1e-7)

	b = make([]byte, 4)
	buf.Read(b)
	altitude := binary.BigEndian.Uint32(b)
	fmt.Printf("longitude: %+v\n", altitude)

	b = make([]byte, 4)
	buf.Read(b)
	speed := float32(binary.BigEndian.Uint32(b)) * 0.036
	fmt.Printf("speed: %+v\n", speed)

	b = make([]byte, 2)
	buf.Read(b)
	heading := binary.BigEndian.Uint16(b)
	fmt.Printf("heading: %+v\n", heading)

	satellites, _ := buf.ReadByte()
	fmt.Printf("satellites: %+v\n", satellites)

	fixStatus, _ := buf.ReadByte()
	fmt.Printf("fixStatus: %+v\n", fixStatus)

	b = make([]byte, 2)
	buf.Read(b)
	Carrier := binary.BigEndian.Uint16(b)
	fmt.Printf("Carrier: %+v\n", Carrier)

	b = make([]byte, 2)
	buf.Read(b)
	RSSI := binary.BigEndian.Uint16(b)
	fmt.Printf("RSSI: %+v\n", int16(RSSI))

	commState, _ := buf.ReadByte()
	fmt.Printf("commState: %08b\n", commState)

	fmt.Printf("commState Available %+v\n", (commState>>0)&1)
	fmt.Printf("commState Network Service %+v\n", (commState>>1)&1)
	fmt.Printf("commState Data Service %+v\n", (commState>>2)&1)
	fmt.Printf("commState Connected %+v\n", (commState>>3)&1)
	fmt.Printf("commState Voice Call is Active %+v\n", (commState>>4)&1)
	fmt.Printf("commState Roaming %+v\n", (commState>>5)&1)
	fmt.Printf("commState Network Technology %+v %+v\n", (commState>>6)&1, (commState>>7)&1)

	hdop, _ := buf.ReadByte()
	fmt.Printf("hdop: %+v\n", hdop)

	inputs, _ := buf.ReadByte()
	fmt.Printf("inputs: %08b\n", inputs)

	fmt.Printf("inputs ignition %+v\n", (inputs>>0)&1)
	fmt.Printf("inputs 1 %+v\n", (inputs>>1)&1)
	fmt.Printf("inputs 2 %+v\n", (inputs>>2)&1)
	fmt.Printf("inputs 3 %+v\n", (inputs>>3)&1)
	fmt.Printf("inputs 4 %+v\n", (inputs>>4)&1)
	fmt.Printf("inputs 5 %+v\n", (inputs>>5)&1)
	fmt.Printf("inputs 6 %+v\n", (inputs>>6)&1)

	unitStatus, _ := buf.ReadByte()
	fmt.Printf("unitStatus: %08b\n", unitStatus)

	fmt.Printf("unitStatus LMU32: HTTP OTA Update Status %+v\n", (unitStatus>>0)&1)
	fmt.Printf("unitStatus GPS Antenna Status %+v\n", (unitStatus>>1)&1)
	fmt.Printf("unitStatus GPS Receiver Self-Test %+v\n", (unitStatus>>2)&1)
	fmt.Printf("unitStatus GPS Receiver Tracking %+v\n", (unitStatus>>3)&1)

	eventIndex, _ := buf.ReadByte()
	fmt.Printf("eventIndex: %+v\n", eventIndex)

	eventCode, _ := buf.ReadByte()
	fmt.Printf("eventCode: %+v\n", eventCode)

	accums, _ := buf.ReadByte()
	fmt.Printf("Accums: %+v\n", accums)

	append, _ := buf.ReadByte()
	fmt.Printf("append: %+v\n", append)

	accumCount := accums & 0x3F

	for i := 0; i < int(accumCount); i++ {
		b = make([]byte, 4)
		buf.Read(b)
		accum := binary.BigEndian.Uint32(b)
		fmt.Printf("accum: %+v %+v\n", i, accum)
	}

}
