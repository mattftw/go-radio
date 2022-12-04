package adsb

import (
	"encoding/binary"
	"errors"
)

const df17MessageLengthBits = 112

var (
	ErrInvalidMessageLength = errors.New("message length invalid")
	ErrInvalidPayloadLength = errors.New("payload length invalid")
	ErrNotDF17Message       = errors.New("not df17 message")
	ErrCRCInvalid           = errors.New("CRC invalid")
)

var df17MessagePayloadTypes = []DF17MessagePayload{
	&DF17MessageAircraftIdentification{},
	&DF17MessageAircraftPosition{},
	&DF17MessageAircraftVelocity{},
}

type DF17Message struct {
	Capability  byte
	ICAOAddress uint32
	Payload     DF17MessagePayload
}

func (msg *DF17Message) Validate() error {
	if msg.Capability > 0b00000111 {
		return errors.New("capability needs to be between 0b00000000 and 0b00000111")
	}

	if msg.ICAOAddress > 0xFFFFFF {
		return errors.New("highest ICAO address is 0xFFFFFF")
	}

	return nil
}

func (msg *DF17Message) MarshalBinary() ([]byte, error) {

	if err := msg.Validate(); err != nil {
		return nil, err
	}

	d := make([]byte, df17MessageLengthBits/8)

	// set the Downlink format
	d[0] = 17 << 3

	// set the capability
	d[0] = d[0] | msg.Capability

	// set the ICAO address
	icaoBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(icaoBytes, msg.ICAOAddress)
	d[1] = icaoBytes[1]
	d[2] = icaoBytes[2]
	d[3] = icaoBytes[3]

	// get the payload
	payload, err := msg.Payload.MarshalBinary()
	if err != nil {
		return nil, err
	}

	if len(payload) != 7 {
		return nil, ErrInvalidPayloadLength
	}

	for i := 0; i < 7; i++ {
		d[4+i] = payload[i]
	}

	checksumInt, err := modeSChecksum(d)
	if err != nil {
		return nil, err
	}

	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, checksumInt)
	d[11] = b[1]
	d[12] = b[2]
	d[13] = b[3]

	return d, nil
}

func (msg *DF17Message) UnmarshalBinary(b []byte) error {
	// check message format
	if (b[0]&0b11111000)>>3 != 17 {
		return ErrNotDF17Message
	}

	// basic message validation
	actualMsgLength := len(b)
	if actualMsgLength != 14 {
		return ErrInvalidMessageLength
	}

	// ensure checksum is valid
	checksumInt, err := modeSChecksum(b)
	if err != nil {
		return err
	}
	checksum := make([]byte, 4)
	binary.BigEndian.PutUint32(checksum, checksumInt)
	if checksum[1] != b[actualMsgLength-3] || checksum[2] != b[actualMsgLength-2] || checksum[3] != b[actualMsgLength-1] {
		return ErrCRCInvalid
	}

	// set capability
	msg.Capability = b[0] & 0b00000111

	// set icao address
	msg.ICAOAddress = binary.BigEndian.Uint32([]byte{0x00, b[1], b[2], b[3]})

	// unmarshal the payload
	payloadData := b[4 : 4+7]
	typeCode := (payloadData[0] & 0b11111000) >> 3

	for _, payloadType := range df17MessagePayloadTypes {
		if byteInSlice(typeCode, payloadType.TypeCodes()) {
			if err := payloadType.UnmarshalBinary(payloadData); err != nil {
				return err
			}
			msg.Payload = payloadType
			break
		}
	}

	if msg.Payload == nil {
		msg.Payload = &DF17MessageUnsupportedPayload{}
		if err := msg.Payload.UnmarshalBinary(payloadData); err != nil {
			return err
		}
	}

	return nil
}
