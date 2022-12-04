package adsb

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDF17MessageAircraftVelocityAirSubsonicMarshalBinary(t *testing.T) {
	msg := DF17MessageAircraftVelocity{
		TypeCode: 19,
		Payload: &DF17MessageAircraftVelocityAirSubsonic{
			SubType:             3,
			IntentChange:        false,
			Reserved1:           false,
			VelocityUncertainty: 0,
			HeadingStatus:       true,
			Heading:             694,
			AirspeedType:        true,
			AirSpeed:            376,
			VerticalRateSource:  true,
			VerticalRateSign:    true,
			VerticalRate:        37,
			Reserved2:           0x00,
			DiffFromBaroAltSign: false,
			DiffFromBaroAlt:     0,
		},
	}

	msgBytes := []byte{0x9B, 0x06, 0xB6, 0xAF, 0x18, 0x94, 0x00}

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
