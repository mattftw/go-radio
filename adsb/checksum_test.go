package adsb

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestModeSChecksum(t *testing.T) {
	type tableItem struct {
		input          []byte
		expectedOutput uint32
	}

	table := []tableItem{
		{
			input:          []byte{0x8F, 0x4D, 0x20, 0x23, 0x58, 0x7F, 0x34, 0x5E, 0x35, 0x83, 0x7E, 0x22, 0x18, 0xB2}, // DF17
			expectedOutput: 0x2218B2,
		},
	}

	for _, tt := range table {
		actualOutput, err := modeSChecksum(tt.input)
		require.Nil(t, err)
		require.Equal(t, actualOutput, tt.expectedOutput)
	}
}
