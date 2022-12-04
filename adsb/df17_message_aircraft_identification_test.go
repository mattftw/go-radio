package adsb

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDF17MessageAircraftIdentificationMarshalBinary(t *testing.T) {
	msg := DF17MessageAircraftIdentification{
		TypeCode:        4,
		EmitterCategory: 0,
		String:          "KLM1023_",
	}

	msgBytes := []byte{0x20, 0x2C, 0xC3, 0x71, 0xC3, 0x2C, 0xE0}

	t.Run("marshal", func(t *testing.T) {
		d, err := msg.MarshalBinary()
		require.Nil(t, err)
		require.Equal(t, msgBytes, d)
	})

	t.Run("unmarshal", func(t *testing.T) {
		unmarshalInto := DF17MessageAircraftIdentification{}
		err := unmarshalInto.UnmarshalBinary(msgBytes)
		require.Nil(t, err)
		require.Equal(t, msg, unmarshalInto)
	})

}
