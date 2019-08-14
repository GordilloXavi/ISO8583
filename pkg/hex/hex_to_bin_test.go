package hex_test

import (
	"testing"

	"github.com/GordilloXavi/ISO8583/pkg/hex"
	"github.com/stretchr/testify/assert"
)

// I == 1, O == 0
const (
	I = true
	O = false
)

var hexNumbers = []struct {
	hex            string
	expectedResult []bool
	err            bool
	expectedError  string
}{
	{
		hex:            "ACAB",
		expectedResult: []bool{I, O, I, O, I, I, O, O, I, O, I, O, I, O, I, I},
	},
	{
		hex:            "0123456789ABCDEF",
		expectedResult: []bool{O, O, O, O, O, O, O, I, O, O, I, O, O, O, I, I, O, I, O, O, O, I, O, I, O, I, I, O, O, I, I, I, I, O, O, O, I, O, O, I, I, O, I, O, I, O, I, I, I, I, O, O, I, I, O, I, I, I, I, O, I, I, I, I},
	},
	{
		hex:            "0000",
		expectedResult: []bool{O, O, O, O, O, O, O, O, O, O, O, O, O, O, O, O},
	},
	{
		hex:            "",
		expectedResult: []bool{},
	},
	{
		"THIS IS NOT A HEXADECIMAL NUMBER!! :D",
		[]bool{},
		true,
		"A non-hexadecimal digit present in the input string: T",
	},
}

func TestHex(t *testing.T) {
	for _, h := range hexNumbers {
		bin, err := hex.HexToBin(h.hex)
		if h.err {
			assert.EqualError(t, err, h.expectedError)
		} else {
			assert.NoError(t, err)
		}
		assert.Equal(t, bin, h.expectedResult)
	}
}
