package adsb

import (
	"encoding/binary"
	"errors"
	"regexp"
	"strings"
)

var aircraftInformationCharset = "#ABCDEFGHIJKLMNOPQRSTUVWXYZ#####_###############0123456789######"
var regexMatcher = regexp.MustCompile(`[A-Za-z_0-9]`)

type DF17MessageAircraftIdentification struct {
	TypeCode        byte `adsbFieldLength:"5" adsbFieldOffset:"0"`
	EmitterCategory byte `adsbFieldLength:"3" adsbFieldOffset:"5"`
	String          string
}

func (msg *DF17MessageAircraftIdentification) UnmarshalBinary(d []byte) error {
	msg.TypeCode = (d[0] & 0b11111000) >> 3
	msg.EmitterCategory = d[0] & 0b00000111

	// unpack the string
	msg.String = ""
	v := binary.BigEndian.Uint64([]byte{
		d[1],
		d[2],
		d[3],
		d[4],
		d[5],
		d[6],
		0x00,
		0x00,
	})
	for i := 58; i >= 16; i -= 6 {
		c := (v >> i) & 0b00111111
		msg.String += string(aircraftInformationCharset[c])
	}

	return nil
}

func (msg *DF17MessageAircraftIdentification) MarshalBinary() ([]byte, error) {
	if err := msg.Validate(); err != nil {
		return nil, err
	}

	ret := make([]byte, 7)

	// set the typecode
	ret[0] = msg.TypeCode << 3

	// set the emmitter category
	ret[0] = ret[0] | msg.EmitterCategory

	// convert the string to the special characterset
	stringFixed := strings.ToUpper(msg.String)
	var v uint64
	strlen := len(msg.String)
	for i := 0; i < strlen; i++ {
		position := strings.Index(aircraftInformationCharset, string(stringFixed[i]))
		nv := uint64(position) << (6 * (7 - i))
		v = v | nv
	}

	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)
	ret[1] = b[2]
	ret[2] = b[3]
	ret[3] = b[4]
	ret[4] = b[5]
	ret[5] = b[6]
	ret[6] = b[7]

	return ret, nil
}

func (msg *DF17MessageAircraftIdentification) Validate() error {

	if msg.TypeCode < 1 || msg.TypeCode > 4 {
		return errors.New("type code must be > 0 and < 5")
	}

	if msg.EmitterCategory > 0b00000111 {
		return errors.New("EmitterCategory must be <= 7")
	}

	if len(msg.String) > 8 {
		return errors.New("string length must be <=6")
	}

	if !regexMatcher.MatchString(msg.String) {
		return errors.New("string must match A-Z, 0-9, _")
	}

	return nil
}

func (msg *DF17MessageAircraftIdentification) TypeCodes() []byte {
	return []byte{1, 2, 3, 4}
}
