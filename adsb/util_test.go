package adsb

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLengthToBitmask(t *testing.T) {
	tt := map[uint]uint64{
		1:  1,
		2:  3,
		3:  7,
		4:  15,
		5:  31,
		6:  63,
		7:  127,
		8:  255,
		9:  511,
		10: 1023,
		11: 2047,
	}

	for k, v := range tt {
		t.Run(fmt.Sprintf("%d to %d", k, v), func(t *testing.T) {
			require.Equal(t, v, lengthToBitmask(k))
		})
	}
}
