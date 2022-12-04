package adsb

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDF17MessageMarshaling(t *testing.T) {
	msg := DF17Message{
		Capability:  5,
		ICAOAddress: 0x4840D6,
		Payload: &DF17MessageAircraftIdentification{
			TypeCode:        4,
			EmitterCategory: 0,
			String:          "KLM1023_",
		},
	}
	msgBytes := []byte{0x8D, 0x48, 0x40, 0xD6, 0x20, 0x2C, 0xC3, 0x71, 0xC3, 0x2C, 0xE0, 0x57, 0x60, 0x98}

	t.Run("marshal", func(t *testing.T) {
		d, err := msg.MarshalBinary()
		require.Nil(t, err)
		require.Equal(t, msgBytes, d)
	})

	t.Run("unmarshal", func(t *testing.T) {
		n := DF17Message{}
		err := n.UnmarshalBinary(msgBytes)
		require.Nil(t, err)
		require.Equal(t, msg, n)
	})
}
