package adsb

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDF17MessageAircraftVelocityUnsupportedMarshalBinary(t *testing.T) {
	msgBytes := []byte{0x9F, 0x44, 0x09, 0x94, 0x08, 0x38, 0x17}

	msg := DF17MessageAircraftVelocity{
		TypeCode: 19,
		Payload: &DF17MessageAircraftVelocityUnsupported{
			Payload: msgBytes,
		},
	}

	t.Run("marshal", func(t *testing.T) {
		d, err := msg.MarshalBinary()
		require.Nil(t, err)
		require.Equal(t, msgBytes, d)
	})

	t.Run("unmarshal", func(t *testing.T) {
		unmarshalInto := DF17MessageAircraftVelocity{}
		err := unmarshalInto.UnmarshalBinary(msgBytes)
		require.Nil(t, err)
		require.Equal(t, msg, unmarshalInto)
	})

}
