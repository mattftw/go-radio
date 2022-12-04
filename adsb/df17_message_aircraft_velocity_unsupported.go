package adsb

type DF17MessageAircraftVelocityUnsupported struct {
	Payload []byte
}

func (msg *DF17MessageAircraftVelocityUnsupported) UnmarshalBinary(d []byte) error {
	msg.Payload = d
	return nil
}

func (msg *DF17MessageAircraftVelocityUnsupported) MarshalBinary() ([]byte, error) {
	return msg.Payload, nil
}

func (msg *DF17MessageAircraftVelocityUnsupported) SubTypeCodes() []byte {
	return []byte{}
}
