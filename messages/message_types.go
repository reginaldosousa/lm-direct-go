package messages

type MessageType uint8

const (
	MessageTypeNull            MessageType = 0
	MessageTypeAck             MessageType = 1
	MessageTypeEventReport     MessageType = 2
	MessageTypeIdReport        MessageType = 3
	MessageTypeUserData        MessageType = 4
	MessageTypeAppData         MessageType = 5
	MessageTypeConfig          MessageType = 6
	MessageTypeUnitRequest     MessageType = 7
	MessageTypeLocateReport    MessageType = 8
	MessageTypeUserDataAcc     MessageType = 9
	MessageTypeMiniEventReport MessageType = 10
	MessageTypeMiniUserData    MessageType = 11
)
