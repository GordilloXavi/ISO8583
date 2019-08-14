package iso8583

import (
	"encoding/json"
	"fmt"
	"strconv"
)

//Returns a map {fieldNumber : fieldValue} representing the message fields
func Decode(msg *message) (map[string]string, error) {
	retMap := make(map[string]string)

	mti, err := msg.getMessageType()
	if err != nil {
		return retMap, err
	}
	retMap["messageType"] = mti

	var pos int
	for _, f := range msg.messageFields {
		field, found := fields[f.pos]
		if !found {
			return retMap, fmt.Errorf("Field not found: %v\n", f.pos)
		}

		val, err := msg.processField(&pos, field)
		if err != nil {
			return retMap, err
		}
		//retMap[field.name] = val // TODO: Use names as keys for more readable output
		retMap[strconv.Itoa(f.pos)] = val
	}
	msg.FieldMap = &retMap
	return retMap, nil
}

// Converts a message to its json representation
func (m *message) ToJSON() (string, error) {
	if m.FieldMap == nil {
		_, err := Decode(m)
		if err != nil {
			return "", err
		}
	}
	j, err := json.Marshal(*m.FieldMap)
	if err != nil {
		return "", err
	}
	return string(j), nil
}
