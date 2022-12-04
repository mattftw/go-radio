package adsb

type DF17MessageUnsupportedPayload struct {
	Payload []byte
}

func (msg *DF17MessageUnsupportedPayload) UnmarshalBinary(d []byte) error {
	if len(d) != 7 {
		return ErrInvalidPayloadLength
	}

	msg.Payload = d
	return nil
}

func (msg *DF17MessageUnsupportedPayload) MarshalBinary() ([]byte, error) {
	if len(msg.Payload) != 7 {
		return nil, ErrInvalidPayloadLength
	}

	return msg.Payload, nil
}

func (msg *DF17MessageUnsupportedPayload) TypeCodes() []byte {
	return []byte{}
}
