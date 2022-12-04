package adsb

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDF17MessageAircraftVelocityGroundSubsonicMarshalBinary(t *testing.T) {
	msg := DF17MessageAircraftVelocity{
		TypeCode: 19,
		Payload: &DF17MessageAircraftVelocityGroundSubsonic{
			SubType:                1,
			IntentChange:           false,
			Reserved1:              true,
			VelocityUncertainty:    0,
			EastWestVelocitySign:   true,
			EastWestVelocity:       9,
			NorthSouthVelocitySign: true,
			NorthSouthVelocity:     160,
			VerticalRateSource:     false,
			VerticalRateSign:       true,
			VerticalRate:           14,
			Reserved2:              0x00,
			DiffFromBaroAltSign:    false,
			DiffFromBaroAlt:        23,
		},
	}

	msgBytes := []byte{0x99, 0x44, 0x09, 0x94, 0x08, 0x38, 0x17}

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
