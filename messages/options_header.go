package messages

import (
	"bytes"
	"encoding/hex"
)

type OptionsHeader struct {
	Options            Options
	MobileIDLength     int
	MobileID           string
	MobileIDTypeLength int
	MobileIDType       string
}

type MobileIDType uint8

const (
	MobileIDTypeOff         MobileIDType = 0
	MobileIDTypeEsn         MobileIDType = 1
	MobileIDTypeEquipment   MobileIDType = 2
	MobileIDTypeSubscriber  MobileIDType = 3
	MobileIDTypeDefined     MobileIDType = 4
	MobileIDTypePhoneNumber MobileIDType = 5
	MobileIDTypeIpAddress   MobileIDType = 6
	MobileIDTypeCdma        MobileIDType = 7
)

func parseOptions(options *Options, optionsByte byte) {
	options.IsMobileID = (optionsByte>>0)&1 == 1
	options.IsMobileIDType = (optionsByte>>1)&1 == 1
	options.IsAuthenticationWord = (optionsByte>>2)&1 == 1
	options.IsRouting = (optionsByte>>3)&1 == 1
	options.IsForwarding = (optionsByte>>4)&1 == 1
	options.IsResponseRedirection = (optionsByte>>5)&1 == 1
	options.IsOptionsExtension = (optionsByte>>6)&1 == 1
	options.IsAlwaysSet = (optionsByte>>7)&1 == 1
}

type Options struct {
	IsMobileID            bool
	IsMobileIDType        bool
	IsAuthenticationWord  bool
	IsRouting             bool
	IsForwarding          bool
	IsResponseRedirection bool
	IsOptionsExtension    bool
	IsAlwaysSet           bool
}

func ParseOptionsHeader(buf *bytes.Buffer) (optionsHeader OptionsHeader, err error) {
	options, err := buf.ReadByte()
	if err != nil {
		return
	}
	parseOptions(&optionsHeader.Options, options)

	mobileIdLength, err := buf.ReadByte()
	if err != nil {
		return
	}
	optionsHeader.MobileIDLength = int(mobileIdLength)

	b := make([]byte, mobileIdLength)
	buf.Read(b)
	optionsHeader.MobileID = hex.EncodeToString(b)

	mobileIdTypeLength, err := buf.ReadByte()
	optionsHeader.MobileIDTypeLength = int(mobileIdTypeLength)

	b = make([]byte, mobileIdTypeLength)
	buf.Read(b)
	optionsHeader.MobileIDType = hex.EncodeToString(b)

	return
}
