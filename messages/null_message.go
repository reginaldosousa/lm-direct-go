package messages

import "bytes"

type NullMessageBody struct{}

func ParseNullMessageBody(buffer *bytes.Buffer) (NullMessageBody, error) {
	return NullMessageBody{}, nil
}
