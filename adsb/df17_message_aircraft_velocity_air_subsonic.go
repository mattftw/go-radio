package adsb

import "encoding/binary"

type DF17MessageAircraftVelocityAirSubsonic struct {
	SubType             byte   `adsbFieldLength:"3" adsbFieldOffset:"5"`   // ST
	IntentChange        bool   `adsbFieldLength:"1" adsbFieldOffset:"8"`   // IC
	Reserved1           bool   `adsbFieldLength:"1" adsbFieldOffset:"9"`   // RESV_A
	VelocityUncertainty byte   `adsbFieldLength:"3" adsbFieldOffset:"10"`  // NAC
	HeadingStatus       bool   `adsbFieldLength:"1" adsbFieldOffset:"13"`  // H-s
	Heading             uint16 `adsbFieldLength:"10" adsbFieldOffset:"14"` // Hdg
	AirspeedType        bool   `adsbFieldLength:"1" adsbFieldOffset:"24"`  // AS-T
	AirSpeed            uint16 `adsbFieldLength:"10" adsbFieldOffset:"25"` // AS
	VerticalRateSource  bool   `adsbFieldLength:"1" adsbFieldOffset:"35"`  // VrSrc
	VerticalRateSign    bool   `adsbFieldLength:"1" adsbFieldOffset:"36"`  // S-Vr
	VerticalRate        uint16 `adsbFieldLength:"9" adsbFieldOffset:"37"`  // Vr
	Reserved2           byte   `adsbFieldLength:"2" adsbFieldOffset:"46"`  // RESV_B
	DiffFromBaroAltSign bool   `adsbFieldLength:"1" adsbFieldOffset:"48"`  // S-Dif
	DiffFromBaroAlt     byte   `adsbFieldLength:"7" adsbFieldOffset:"49"`  // Dif
}

func (msg *DF17MessageAircraftVelocityAirSubsonic) UnmarshalBinary(d []byte) error {
	payload := binary.BigEndian.Uint64(append(d, 0x00))
	applyPayloadFieldToStruct(&payload, msg, "SubType")
	applyPayloadFieldToStruct(&payload, msg, "IntentChange")
	applyPayloadFieldToStruct(&payload, msg, "Reserved1")
	applyPayloadFieldToStruct(&payload, msg, "VelocityUncertainty")
	applyPayloadFieldToStruct(&payload, msg, "HeadingStatus")
	applyPayloadFieldToStruct(&payload, msg, "Heading")
	applyPayloadFieldToStruct(&payload, msg, "AirspeedType")
	applyPayloadFieldToStruct(&payload, msg, "AirSpeed")
	applyPayloadFieldToStruct(&payload, msg, "VerticalRateSource")
	applyPayloadFieldToStruct(&payload, msg, "VerticalRateSign")
	applyPayloadFieldToStruct(&payload, msg, "VerticalRate")
	applyPayloadFieldToStruct(&payload, msg, "Reserved2")
	applyPayloadFieldToStruct(&payload, msg, "DiffFromBaroAltSign")
	applyPayloadFieldToStruct(&payload, msg, "DiffFromBaroAlt")
	return nil
}

func (msg *DF17MessageAircraftVelocityAirSubsonic) MarshalBinary() ([]byte, error) {
	dataInt := uint64(0)
	applyFieldToPayload(&dataInt, msg, "SubType")
	applyFieldToPayload(&dataInt, msg, "IntentChange")
	applyFieldToPayload(&dataInt, msg, "Reserved1")
	applyFieldToPayload(&dataInt, msg, "VelocityUncertainty")
	applyFieldToPayload(&dataInt, msg, "HeadingStatus")
	applyFieldToPayload(&dataInt, msg, "Heading")
	applyFieldToPayload(&dataInt, msg, "AirspeedType")
	applyFieldToPayload(&dataInt, msg, "AirSpeed")
	applyFieldToPayload(&dataInt, msg, "VerticalRateSource")
	applyFieldToPayload(&dataInt, msg, "VerticalRateSign")
	applyFieldToPayload(&dataInt, msg, "VerticalRate")
	applyFieldToPayload(&dataInt, msg, "Reserved2")
	applyFieldToPayload(&dataInt, msg, "DiffFromBaroAltSign")
	applyFieldToPayload(&dataInt, msg, "DiffFromBaroAlt")

	data := make([]byte, 8)
	binary.BigEndian.PutUint64(data, dataInt)

	return data[0:7], nil
}

func (msg *DF17MessageAircraftVelocityAirSubsonic) SubTypeCodes() []byte {
	return []byte{3}
}
