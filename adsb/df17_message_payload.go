package adsb

type DF17MessagePayload interface {
	MarshalBinary() ([]byte, error)
	UnmarshalBinary([]byte) error
	TypeCodes() []byte
}
