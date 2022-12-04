package adsb

import "encoding/binary"

type DF17MessageAircraftVelocityGroundSubsonic struct {
	SubType                byte   `adsbFieldLength:"3" adsbFieldOffset:"5"`   // ST
	IntentChange           bool   `adsbFieldLength:"1" adsbFieldOffset:"8"`   // IC
	Reserved1              bool   `adsbFieldLength:"1" adsbFieldOffset:"9"`   // RESV_A
	VelocityUncertainty    byte   `adsbFieldLength:"3" adsbFieldOffset:"10"`  // NAC
	EastWestVelocitySign   bool   `adsbFieldLength:"1" adsbFieldOffset:"13"`  // S-EW
	EastWestVelocity       uint16 `adsbFieldLength:"10" adsbFieldOffset:"14"` // V-EW
	NorthSouthVelocitySign bool   `adsbFieldLength:"1" adsbFieldOffset:"24"`  // S-NS
	NorthSouthVelocity     uint16 `adsbFieldLength:"10" adsbFieldOffset:"25"` // V-NS
	VerticalRateSource     bool   `adsbFieldLength:"1" adsbFieldOffset:"35"`  // VrSrc
	VerticalRateSign       bool   `adsbFieldLength:"1" adsbFieldOffset:"36"`  // S-Vr
	VerticalRate           uint16 `adsbFieldLength:"9" adsbFieldOffset:"37"`  // Vr
	Reserved2              byte   `adsbFieldLength:"2" adsbFieldOffset:"46"`  // RESV_B
	DiffFromBaroAltSign    bool   `adsbFieldLength:"1" adsbFieldOffset:"48"`  // S-Dif
	DiffFromBaroAlt        byte   `adsbFieldLength:"7" adsbFieldOffset:"49"`  // Dif
}

func (msg *DF17MessageAircraftVelocityGroundSubsonic) UnmarshalBinary(d []byte) error {
	payload := binary.BigEndian.Uint64(append(d, 0x00))
	applyPayloadFieldToStruct(&payload, msg, "SubType")
	applyPayloadFieldToStruct(&payload, msg, "IntentChange")
	applyPayloadFieldToStruct(&payload, msg, "Reserved1")
	applyPayloadFieldToStruct(&payload, msg, "VelocityUncertainty")
	applyPayloadFieldToStruct(&payload, msg, "EastWestVelocitySign")
	applyPayloadFieldToStruct(&payload, msg, "EastWestVelocity")
	applyPayloadFieldToStruct(&payload, msg, "NorthSouthVelocitySign")
	applyPayloadFieldToStruct(&payload, msg, "NorthSouthVelocity")
	applyPayloadFieldToStruct(&payload, msg, "VerticalRateSource")
	applyPayloadFieldToStruct(&payload, msg, "VerticalRateSign")
	applyPayloadFieldToStruct(&payload, msg, "VerticalRate")
	applyPayloadFieldToStruct(&payload, msg, "Reserved2")
	applyPayloadFieldToStruct(&payload, msg, "DiffFromBaroAltSign")
	applyPayloadFieldToStruct(&payload, msg, "DiffFromBaroAlt")
	return nil
}

func (msg *DF17MessageAircraftVelocityGroundSubsonic) MarshalBinary() ([]byte, error) {
	dataInt := uint64(0)
	applyFieldToPayload(&dataInt, msg, "SubType")
	applyFieldToPayload(&dataInt, msg, "IntentChange")
	applyFieldToPayload(&dataInt, msg, "Reserved1")
	applyFieldToPayload(&dataInt, msg, "VelocityUncertainty")
	applyFieldToPayload(&dataInt, msg, "EastWestVelocitySign")
	applyFieldToPayload(&dataInt, msg, "EastWestVelocity")
	applyFieldToPayload(&dataInt, msg, "NorthSouthVelocitySign")
	applyFieldToPayload(&dataInt, msg, "NorthSouthVelocity")
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

func (msg *DF17MessageAircraftVelocityGroundSubsonic) SubTypeCodes() []byte {
	return []byte{1}
}
