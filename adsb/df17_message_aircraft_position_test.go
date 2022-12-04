package adsb

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDF17MessageAircraftPositionMarshalBinary(t *testing.T) {
	msg := DF17MessageAircraftPosition{
		TypeCode:           11,
		SurveillanceStatus: 0,
		NICsb:              false,
		Altitude:           3128,
		Time:               false,
		FrameType:          DF17MessageAircraftPositionFrameTypeEven,
		LatitudeCPR:        93000,
		LongitudeCPR:       51372,
	}

	msgBytes := []byte{0x58, 0xC3, 0x82, 0xD6, 0x90, 0xC8, 0xAC}

	t.Run("marshal", func(t *testing.T) {
		d, err := msg.MarshalBinary()
		require.Nil(t, err)
		require.Equal(t, msgBytes, d)
	})

	t.Run("unmarshal", func(t *testing.T) {
		unmarshalInto := DF17MessageAircraftPosition{}
		err := unmarshalInto.UnmarshalBinary(msgBytes)
		require.Nil(t, err)
		require.Equal(t, msg, unmarshalInto)
	})

}
