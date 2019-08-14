package iso8583_test

import (
	"strconv"
	"testing"

	"github.com/JoinVerse/cardinal/pkg/iso8583"
	"github.com/stretchr/testify/assert"
)

func TestValidMessages(t *testing.T) {

	for _, m := range validMessages {
		// Test message initialization
		msg, err := iso8583.NewMessage(m.message)
		assert.NoError(t, err)

		// Test Decode method
		decoded_msg, err := iso8583.Decode(&msg)
		assert.NoError(t, err)
		assert.Equal(t, decoded_msg, m.expectedResult)

		// Test ToJSON method
		j, err := msg.ToJSON()
		assert.NoError(t, err)
		assert.Equal(t, j, m.json)
	}
}

func TestWrongMTI(t *testing.T) {
	msg, err := iso8583.NewMessage(wrongMti)
	assert.NoError(t, err)
	_, err = iso8583.Decode(&msg)
	assert.EqualError(t, err, "\"1234\" is not a valid message type")
}

func TestMessageTooShort(t *testing.T) {
	for _, m := range shortMessages {
		_, err := iso8583.NewMessage(m)
		errorString := "Message length too small. Espected at least 23 characters, got " + strconv.Itoa(len(m))
		assert.EqualError(t, err, errorString)
	}
}

func TestFieldNotFound(t *testing.T) {
	var err error
	msg, err := iso8583.NewMessage(invalidField)
	assert.NoError(t, err)

	_, err = iso8583.Decode(&msg)
	assert.EqualError(t, err, "Field not found: 2\n")
}

func TestNonHexadecimalBitmap(t *testing.T) {
	_, err := iso8583.NewMessage(nonHexBitmap)
	assert.EqualError(t, err, "A non-hexadecimal digit present in the input string: H")
}

func TestDataNotPresent(t *testing.T) {
	msg, err := iso8583.NewMessage(dataNotPresent)
	assert.NoError(t, err)

	_, err = iso8583.Decode(&msg)
	assert.EqualError(t, err, "Trying to acess position 16 of string with length 3\n")
}
