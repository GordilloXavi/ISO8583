package iso8583

import (
	"fmt"
	"strconv"

	"github.com/JoinVerse/cardinal/pkg/hex"
)

// Represents a message split by its main fieds
type message struct {
	// The whole unedited message
	rawData string

	//length of the message (bytes 1-4)
	length int

	//message type indicator
	messageType mti

	//First bitmap
	bitmap1 string

	//Secondary bitmap (optional)
	bitmap2 string

	//data fields for the message
	data []rune

	//Int array representing the data fields present in the message
	messageFields []messageField

	//Points to a map with the field_number : value of all fields in the message
	FieldMap *map[string]string
}

type mti string

const (
	AUTH_REQUEST      mti = "0100"
	AUTH_RESPONSE     mti = "0110"
	REVERSAL_REQUEST  mti = "0420"
	REVERSAL_RESPONSE mti = "0430"
	ADVICE_REQUEST    mti = "0120"
	ADVICE_RESPONSE   mti = "0130"
)

//Returns a message object
func NewMessage(rawData string) (message, error) {
	length, messageType, bitmap1, bitmap2, data, err := splitMessage(rawData)
	if err != nil {
		return message{}, err
	}

	messageFields, err := getMessageFields(bitmap1, bitmap2)
	if err != nil {
		return message{}, err
	}

	return message{
		rawData:       rawData,
		length:        length,
		messageType:   messageType,
		bitmap1:       bitmap1,
		bitmap2:       bitmap2,
		data:          data,
		messageFields: messageFields,
	}, nil
}

//
func (m *message) AddField(p int, fieldString string) error {

	_, found := fields[p]
	if !found {
		return fmt.Errorf("Field not found: %v\n", p)
	}
	field := *fields[p] // FIXME: Find a better way to do this
	field.processFieldValue(fieldString)

	m.messageFields = append(m.messageFields, field) // FIXME:May have problems if we pass pointers like this
	return nil
}

func (m *message) getMessageType() (string, error) {

	switch m.messageType {
	case AUTH_REQUEST:
		return "Auth request", nil
	case AUTH_RESPONSE:
		return "Auth response", nil
	case REVERSAL_REQUEST:
		return "Reversal request", nil
	case REVERSAL_RESPONSE:
		return "Reversal response", nil
	case ADVICE_REQUEST:
		return "Advice request", nil
	case ADVICE_RESPONSE:
		return "Advice response", nil
	}
	return "", fmt.Errorf("\"%v\" is not a valid message type", m.messageType)
}

func (msg *message) processField(pos *int, field *messageField) (string, error) {

	if field.hasVariableSize {
		field.size, _ = strconv.Atoi(string(msg.data[*pos : *pos+field.varSizeIndicator]))
		*pos += field.varSizeIndicator
	}

	if *pos+field.size > len(msg.data) {
		return "", fmt.Errorf("Trying to acess position %v of string with length %v\n", *pos+field.size, len(msg.data))
	}
	s := string(msg.data[*pos : *pos+field.size])
	*pos += field.size

	return s, nil
}

func splitMessage(data string) (int, mti, string, string, []rune, error) {
	if len(data) <= 22 {
		return 0, "", "", "", nil, fmt.Errorf("Message length too small. Espected at least 23 characters, got %v", len(data))
	}

	//Parses a string representing a 32 bit integer in base 16
	length, _ := strconv.ParseInt(data[:2], 16, 32)
	messageType := mti(data[2:6])
	bitmap1, secondaryBitmap := getBitmap(data, true)
	dataStart := 22
	var bitmap2 string
	if secondaryBitmap {
		bitmap2, _ = getBitmap(data, false)
		_ = bitmap2
		dataStart += 16
	}
	ret_data := []rune(data[dataStart:])

	return int(length), messageType, bitmap1, bitmap2, ret_data, nil

}

func getBitmap(data string, is_primary bool) (string, bool) { //return err??

	//if(len(data) < 39):
	//return "", false, error

	if is_primary {
		var secondary_is_present bool
		switch string(data[6]) { //useless casting??
		case "8", "9", "a", "A", "b", "B", "c", "C", "d", "D", "e", "E", "f", "F":
			secondary_is_present = true
		}
		return data[6:22], secondary_is_present
	}
	return data[22:38], false
}

//Returns a list of int8 representing the data fields present in the message
func getMessageFields(bitmap1 string, bitmap2 string) ([]messageField, error) {

	bits, err := hex.HexToBin(bitmap1)
	if err != nil {
		return []messageField{}, err
	}

	if bitmap2 != "" {
		b, err := hex.HexToBin(bitmap2)
		if err != nil {
			return []messageField{}, err
		}
		bits = append(bits, b...)
	}

	// Slice of integers with initial size = 0 and cap = 15
	messageFields := make([]messageField, 0, 15)

	for i := 1; i < len(bits); i++ {
		//We start at 1 because the first bit is useless once we process the second bitmap
		if bits[i] {
			f, found := fields[i+1]
			if !found {
				//TODO: return error
			}
			field := *f
			field.pos = i + 1
			//field := messageField{pos: i + 1}
			messageFields = append(messageFields, field)
		}
	}
	return messageFields, nil
}
