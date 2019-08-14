package iso8583

import (
	"errors"
	"strconv"
)

//Represents a data element aka a field of the message
type messageField struct {
	// Position of the field in the bitmap
	pos int

	//The name of the field
	name string

	//Brief description of the field
	desc string

	//Indicates wether the size of the field is variable
	hasVariableSize bool

	//Indicates, in case of variable size, how many bytes is the size of the field
	varSizeIndicator int

	//The size of the field in bytes
	size int

	// The value of that field
	value string
}

// Returns an error if val does not fit into field
func (f *messageField) processFieldValue(val string) error {
	var err error

	if f.hasVariableSize {
		varSize, err := toFieldString(strconv.Itoa(len(val)), f.varSizeIndicator)
		f.size = len(val)
		f.value = varSize + val
		return err
	}
	v, err := toFieldString(val, f.size)
	f.value = v

	return err
}

// Takes a string of arbitrary length and adds 0's to its left until its len == size
func toFieldString(txt string, size int) (string, error) {
	if len(txt) > size {
		return "", errors.New("Variable size indicator is too big")
	}
	for i := size - len(txt); i > 0; i-- {
		txt = "0" + txt
	}
	return txt, nil
}

// TODO: add the rest of the message fields
var fields = map[int]*messageField{
	1: &messageField{
		name: "secondary bitmap",
		size: 16,
	},
	3: &messageField{
		name: "Processing Code",
		desc: "",
		size: 6,
	},
	4: &messageField{
		name: "amount1",
		desc: "Amount of transaction in original currency",
		size: 12,
	},
	6: &messageField{
		name: "amount2",
		desc: "Amount of transaction in cardholder billing currency",
		size: 12,
	},
	7: &messageField{
		name: "dateTime",
		desc: "Date and time of the transaction",
		size: 20,
	},
	10: &messageField{
		name: "convRate",
		desc: "cardholder billing conversion rate",
		size: 8,
	},
	11: &messageField{
		name: "traceNumber",
		desc: "Trace number (STAN)",
		size: 6,
	},
	12: &messageField{
		name: "transactionTime",
		desc: "Transaction time",
		size: 6,
	},
	13: &messageField{
		name: "transactionDate",
		desc: "Transaction date",
		size: 4,
	},
	15: &messageField{
		name: "settlementDate",
		desc: "Settlement Date",
		size: 8,
	},
	17: &messageField{
		name: "captureDate",
		desc: "Capture date",
		size: 4,
	},
	18: &messageField{
		name: "merchantType",
		desc: "Merchant Type",
		size: 4,
	},
	22: &messageField{
		name: "POS",
		desc: "Point of Service entry mode",
		size: 3,
	},
	32: &messageField{
		name:             "AcqInstID",
		desc:             "Acquiring institution Id code",
		hasVariableSize:  true,
		varSizeIndicator: 2,
	},
	35: &messageField{
		name:             "cardNum",
		desc:             "Card number",
		hasVariableSize:  true,
		varSizeIndicator: 2,
	},
	37: &messageField{
		name: "RNN",
		desc: "Retrieval Reference Number",
		size: 12,
	},
	38: &messageField{
		name: "authCode",
		desc: "Auth Code - CBS assigned authorization code",
		size: 6,
	},
	39: &messageField{
		name: "authResponseCode",
		desc: "Auth Processing response code: code representing result of authorization",
		size: 2,
	},
	41: &messageField{
		name: "TerminalID",
		desc: "Terminal Id",
		size: 16,
	},
	44: &messageField{
		name: "additionalResponse",
		desc: "Additional Response Data",
		size: 27,
	},
	42: &messageField{
		name: "terminalCode",
		desc: "Terminal code",
		size: 15,
	},
	43: &messageField{
		name: "merchantName",
		desc: "Merchant Name",
		size: 40,
	},
	48: &messageField{
		name:             "additionalData",
		desc:             "",
		hasVariableSize:  true,
		varSizeIndicator: 3,
	},
	49: &messageField{
		name: "ogCurrCode",
		desc: "Original Transaction Currency Code",
		size: 3,
	},
	51: &messageField{
		name: "CHCurrCode",
		desc: "Cardholder Billing Currency Code",
		size: 3,
	},
	61: &messageField{
		name:             "POSData",
		desc:             "Point of Service Data",
		hasVariableSize:  true,
		varSizeIndicator: 3,
	},
	63: &messageField{
		name: "networkData",
		desc: "Network Data",
		size: 12,
	},
	95: &messageField{
		name: "replacementAmounts",
		desc: "Replacement Amounts",
		size: 42,
	},
	108: &messageField{
		name:             "msrd",
		desc:             "Money Send Reference Data",
		hasVariableSize:  true,
		varSizeIndicator: 3,
	},
}
