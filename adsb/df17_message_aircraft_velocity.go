package adsb

import "encoding/binary"

type DF17MessageAircraftVelocityPayload interface {
	MarshalBinary() ([]byte, error)
	UnmarshalBinary([]byte) error
	SubTypeCodes() []byte
}

type DF17MessageAircraftVelocity struct {
	TypeCode byte `adsbFieldLength:"5" adsbFieldOffset:"0"` // TC=17
	Payload  DF17MessageAircraftVelocityPayload
}

var df17MessageAircraftVeloicityPayloadTypes = []DF17MessageAircraftVelocityPayload{
	&DF17MessageAircraftVelocityGroundSubsonic{},
	&DF17MessageAircraftVelocityAirSubsonic{},
}

func (msg *DF17MessageAircraftVelocity) UnmarshalBinary(d []byte) error {

	payload := binary.BigEndian.Uint64(append(d, 0x00))
	applyPayloadFieldToStruct(&payload, msg, "TypeCode")

	for _, payloadType := range df17MessageAircraftVeloicityPayloadTypes {
		if byteInSlice(d[0]&0b00000111, payloadType.SubTypeCodes()) {
			if err := payloadType.UnmarshalBinary(d); err != nil {
				return err
			}
			msg.Payload = payloadType
			break
		}
	}

	if msg.Payload == nil {
		msg.Payload = &DF17MessageAircraftVelocityUnsupported{}
		if err := msg.Payload.UnmarshalBinary(d); err != nil {
			return err
		}
	}

	return nil
}

func (msg *DF17MessageAircraftVelocity) MarshalBinary() ([]byte, error) {

	payload, err := msg.Payload.MarshalBinary()
	if err != nil {
		return nil, err
	}

	payload[0] = payload[0] | (msg.TypeCode << 3)
	return payload, nil
}

func (msg *DF17MessageAircraftVelocity) TypeCodes() []byte {
	return []byte{19}
}
