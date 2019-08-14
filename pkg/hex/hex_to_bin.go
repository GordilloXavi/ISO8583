package hex

import "fmt"

// Auxiliary function needed by iso8583 package to extract hexadecimal bitmaps
//Converts a hexadecimal string to an array of bools representing that string in binary
func HexToBin(hex string) ([]bool, error) {
	bits := make([]bool, 0, len(hex)*4)

	for i := range hex {
		x, err := charToBin(hex[i])
		if err != nil {
			return []bool{}, err
		}
		bits = append(bits, x...)
	}
	return bits, nil
}

// Converts a single hexadecimal ASCII char to its binary representation
func charToBin(char byte) ([]bool, error) {
	res := make([]bool, 4, 4)
	switch string(char) {
	case "0":
		return res, nil
	case "1":
		res[3] = true
	case "2":
		res[2] = true
	case "3":
		res[2] = true
		res[3] = true
	case "4":
		res[1] = true
	case "5":
		res[1] = true
		res[3] = true
	case "6":
		res[1] = true
		res[2] = true
	case "7":
		res = []bool{false, true, true, true}
	case "8":
		res[0] = true
	case "9":
		res[0] = true
		res[3] = true
	case "A", "a":
		res[0] = true
		res[2] = true
	case "B", "b":
		res = []bool{true, false, true, true}
	case "C", "c":
		res[0] = true
		res[1] = true
	case "D", "d":
		res = []bool{true, true, false, true}
	case "E", "e":
		res = []bool{true, true, true, false}
	case "F", "f":
		res = []bool{true, true, true, true}
	default:
		return res, fmt.Errorf("A non-hexadecimal digit present in the input string: %v", string(char))
	}
	return res, nil
}
