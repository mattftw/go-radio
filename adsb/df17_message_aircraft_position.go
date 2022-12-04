package adsb

import (
	"encoding/binary"
	"errors"
)

const (
	DF17MessageAircraftPositionFrameTypeEven = 0
	DF17MessageAircraftPositionFrameTypeOdd  = 1
)

type DF17MessageAircraftPosition struct {
	TypeCode           byte   `adsbFieldLength:"5" adsbFieldOffset:"0"`
	SurveillanceStatus byte   `adsbFieldLength:"2" adsbFieldOffset:"5"`
	NICsb              bool   `adsbFieldLength:"1" adsbFieldOffset:"7"`
	Altitude           uint16 `adsbFieldLength:"12" adsbFieldOffset:"8"`
	Time               bool   `adsbFieldLength:"1" adsbFieldOffset:"20"`
	FrameType          byte   `adsbFieldLength:"1" adsbFieldOffset:"21"`
	LatitudeCPR        uint32 `adsbFieldLength:"17" adsbFieldOffset:"22"`
	LongitudeCPR       uint32 `adsbFieldLength:"17" adsbFieldOffset:"39"`
}

func (msg *DF17MessageAircraftPosition) UnmarshalBinary(d []byte) error {
	payload := binary.BigEndian.Uint64(append(d, 0x00))
	applyPayloadFieldToStruct(&payload, msg, "TypeCode")
	applyPayloadFieldToStruct(&payload, msg, "SurveillanceStatus")
	applyPayloadFieldToStruct(&payload, msg, "NICsb")
	applyPayloadFieldToStruct(&payload, msg, "Altitude")
	applyPayloadFieldToStruct(&payload, msg, "Time")
	applyPayloadFieldToStruct(&payload, msg, "FrameType")
	applyPayloadFieldToStruct(&payload, msg, "LatitudeCPR")
	applyPayloadFieldToStruct(&payload, msg, "LongitudeCPR")
	return nil
}

func (msg *DF17MessageAircraftPosition) MarshalBinary() ([]byte, error) {
	if err := msg.Validate(); err != nil {
		return nil, err
	}

	dataInt := uint64(0)
	applyFieldToPayload(&dataInt, msg, "TypeCode")
	applyFieldToPayload(&dataInt, msg, "SurveillanceStatus")
	applyFieldToPayload(&dataInt, msg, "NICsb")
	applyFieldToPayload(&dataInt, msg, "Altitude")
	applyFieldToPayload(&dataInt, msg, "Time")
	applyFieldToPayload(&dataInt, msg, "FrameType")
	applyFieldToPayload(&dataInt, msg, "LatitudeCPR")
	applyFieldToPayload(&dataInt, msg, "LongitudeCPR")

	data := make([]byte, 8)
	binary.BigEndian.PutUint64(data, dataInt)

	return data[0:7], nil
}

func (msg *DF17MessageAircraftPosition) Validate() error {
	if msg.TypeCode < 9 || msg.TypeCode > 18 {
		return errors.New("type code must be > 8 and < 19")
	}

	if msg.SurveillanceStatus > 0b11 {
		return errors.New("surveillance status must be >= 0  and <= 3")
	}

	return nil
}

func (msg *DF17MessageAircraftPosition) TypeCodes() []byte {
	return []byte{9, 10, 11, 12, 13, 14, 15, 16, 17, 18}
}
