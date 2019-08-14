package iso8583_test

// Messages used in message_test.go
var (
	shortMessages = []string{
		"",
		"000100",
		"0001000000000000200000",
	}

	wrongMti       = "001234367AC40128E1A00A00000000000000140000000000140004020915240402091525610000000000601114360402201704020402504581006123456165222222222222115C0930107FSJS59991234········59991234·······TREVICA EComm·········Warsaw··········PL0106100510001985985021102510000610061601864MCC000123456"
	invalidField   = "000100767AC40128E1A00A00000000000000140000000000140004020915240402091525610000000000601114360402201704020402504581006123456165222222222222115C0930107FSJS59991234········59991234·······TREVICA"
	nonHexBitmap   = "000100HOLAC40128E1A00A00000000000000140000000000140004020915240402091525610000000000601114360402201704020402504581006123456165222222222222115C0930107FSJS59991234········59991234·······TREVICA"
	dataNotPresent = "0001000000000000800000XXX"

	validMessages = []struct {
		message        string
		expectedResult map[string]string
		json           string
	}{
		{ // ECommerce Transaction
			"000100367AC40128E1A00A00000000000000140000000000140004020915240402091525610000000000601114360402201704020402504581006123456165222222222222115C0930107FSJS59991234········59991234·······TREVÍCA EComm·········Warsaw··········PL0106100510001985985021102510000610061601864MCC000123456",
			map[string]string{
				"10":          "61000000",
				"11":          "000060",
				"12":          "111436",
				"13":          "0402",
				"15":          "20170402",
				"17":          "0402",
				"18":          "5045",
				"22":          "810",
				"3":           "000000",
				"32":          "123456",
				"35":          "5222222222222115",
				"37":          "C0930107FSJS",
				"4":           "000000001400",
				"41":          "59991234········",
				"42":          "59991234·······",
				"43":          "TREVÍCA EComm·········Warsaw··········PL",
				"48":          "6100510001",
				"49":          "985",
				"51":          "985",
				"6":           "000000001400",
				"61":          "102510000610061601864",
				"63":          "MCC000123456",
				"7":           "04020915240402091525",
				"messageType": "Auth request",
			},
			"{\"10\":\"61000000\",\"11\":\"000060\",\"12\":\"111436\",\"13\":\"0402\",\"15\":\"20170402\",\"17\":\"0402\",\"18\":\"5045\",\"22\":\"810\",\"3\":\"000000\",\"32\":\"123456\",\"35\":\"5222222222222115\",\"37\":\"C0930107FSJS\",\"4\":\"000000001400\",\"41\":\"59991234········\",\"42\":\"59991234·······\",\"43\":\"TREVÍCA EComm·········Warsaw··········PL\",\"48\":\"6100510001\",\"49\":\"985\",\"51\":\"985\",\"6\":\"000000001400\",\"61\":\"102510000610061601864\",\"63\":\"MCC000123456\",\"7\":\"04020915240402091525\",\"messageType\":\"Auth request\"}",
		},
		{ // ECommerce Transaction
			"0001103238C0002ED0800000000000000000140004020915240402091525000060111436040204025045165222222222222115C0930107FSJS2064200059991234········59991234·······0002985C0000000000000568345985",
			map[string]string{
				"11":          "000060",
				"12":          "111436",
				"13":          "0402",
				"17":          "0402",
				"18":          "5045",
				"3":           "000000",
				"35":          "5222222222222115",
				"37":          "C0930107FSJS",
				"38":          "206420",
				"39":          "00",
				"4":           "000000001400",
				"41":          "59991234········",
				"42":          "59991234·······",
				"44":          "0002985C0000000000000568345",
				"49":          "985",
				"7":           "04020915240402091525",
				"messageType": "Auth response",
			},
			"{\"11\":\"000060\",\"12\":\"111436\",\"13\":\"0402\",\"17\":\"0402\",\"18\":\"5045\",\"3\":\"000000\",\"35\":\"5222222222222115\",\"37\":\"C0930107FSJS\",\"38\":\"206420\",\"39\":\"00\",\"4\":\"000000001400\",\"41\":\"59991234········\",\"42\":\"59991234·······\",\"44\":\"0002985C0000000000000568345\",\"49\":\"985\",\"7\":\"04020915240402091525\",\"messageType\":\"Auth response\"}",
		},
		{ // POS Cash Advance
			"000100B67AC40128E0A00A000000000010000028000000000000160000000000160004021207320402120733610000000000371406480402201704020402653601106123456165222222222222115C0930107GFK759991234········123456789123456Trevica···············Warszawa········PL985985021203000000010061601864MCC000123456SenderReceiverData",
			map[string]string{
				"10":          "61000000",
				"108":         "",
				"11":          "000037",
				"12":          "140648",
				"13":          "0402",
				"15":          "20170402",
				"17":          "0402",
				"18":          "6536",
				"22":          "011",
				"3":           "280000",
				"32":          "123456",
				"35":          "5222222222222115",
				"37":          "C0930107GFK7",
				"4":           "000000001600",
				"41":          "59991234········",
				"42":          "123456789123456",
				"43":          "Trevica···············Warszawa········PL",
				"49":          "985",
				"51":          "985",
				"6":           "000000001600",
				"61":          "203000000010061601864",
				"63":          "MCC000123456",
				"7":           "04021207320402120733",
				"messageType": "Auth request",
			},
			"{\"10\":\"61000000\",\"108\":\"\",\"11\":\"000037\",\"12\":\"140648\",\"13\":\"0402\",\"15\":\"20170402\",\"17\":\"0402\",\"18\":\"6536\",\"22\":\"011\",\"3\":\"280000\",\"32\":\"123456\",\"35\":\"5222222222222115\",\"37\":\"C0930107GFK7\",\"4\":\"000000001600\",\"41\":\"59991234········\",\"42\":\"123456789123456\",\"43\":\"Trevica···············Warszawa········PL\",\"49\":\"985\",\"51\":\"985\",\"6\":\"000000001600\",\"61\":\"203000000010061601864\",\"63\":\"MCC000123456\",\"7\":\"04021207320402120733\",\"messageType\":\"Auth request\"}",
		},
		{ // POS Cash Advance
			"0001103238C0002ED0800000000000000000140004020915240402091525000060111436040204025045165222222222222115C0930107FSJS2064200059991234········59991234·······0002985C0000000000000568345985",
			map[string]string{
				"11":          "000060",
				"12":          "111436",
				"13":          "0402",
				"17":          "0402",
				"18":          "5045",
				"3":           "000000",
				"35":          "5222222222222115",
				"37":          "C0930107FSJS",
				"38":          "206420",
				"39":          "00",
				"4":           "000000001400",
				"41":          "59991234········",
				"42":          "59991234·······",
				"44":          "0002985C0000000000000568345",
				"49":          "985",
				"7":           "04020915240402091525",
				"messageType": "Auth response",
			},
			"{\"11\":\"000060\",\"12\":\"111436\",\"13\":\"0402\",\"17\":\"0402\",\"18\":\"5045\",\"3\":\"000000\",\"35\":\"5222222222222115\",\"37\":\"C0930107FSJS\",\"38\":\"206420\",\"39\":\"00\",\"4\":\"000000001400\",\"41\":\"59991234········\",\"42\":\"59991234·······\",\"44\":\"0002985C0000000000000568345\",\"49\":\"985\",\"7\":\"04020915240402091525\",\"messageType\":\"Auth response\"}",
		},
		{ // Retail with cash back
			"000100367AC40128E1A00A09000000000000160000000000160004021207320402120733610000000000371406480402201704020402599905106123456165222222222222115C0930107GFK759991234········123456789123456Trevica···············Warszawa········PL01061005100019859850200040985D000000000300021000000000020061601864MCC000123456",
			map[string]string{
				"10":          "61000000",
				"11":          "000037",
				"12":          "140648",
				"13":          "0402",
				"15":          "20170402",
				"17":          "0402",
				"18":          "5999",
				"22":          "051",
				"3":           "090000",
				"32":          "123456",
				"35":          "5222222222222115",
				"37":          "C0930107GFK7",
				"4":           "000000001600",
				"41":          "59991234········",
				"42":          "123456789123456",
				"43":          "Trevica···············Warszawa········PL",
				"48":          "6100510001",
				"49":          "985",
				"51":          "985",
				"6":           "000000001600",
				"61":          "0040985D000000000300",
				"63":          "021000000000",
				"7":           "04021207320402120733",
				"messageType": "Auth request",
			},
			"{\"10\":\"61000000\",\"11\":\"000037\",\"12\":\"140648\",\"13\":\"0402\",\"15\":\"20170402\",\"17\":\"0402\",\"18\":\"5999\",\"22\":\"051\",\"3\":\"090000\",\"32\":\"123456\",\"35\":\"5222222222222115\",\"37\":\"C0930107GFK7\",\"4\":\"000000001600\",\"41\":\"59991234········\",\"42\":\"123456789123456\",\"43\":\"Trevica···············Warszawa········PL\",\"48\":\"6100510001\",\"49\":\"985\",\"51\":\"985\",\"6\":\"000000001600\",\"61\":\"0040985D000000000300\",\"63\":\"021000000000\",\"7\":\"04021207320402120733\",\"messageType\":\"Auth request\"}",
		},
		{ // Retail with cash back
			"0001103238C0002ED0800001000000000000200004021008270402100828000067120718040204026010165222222222222115C0930107FT4B2065600060101234········1234567891234560002985C0000000000000062315985",
			map[string]string{
				"11":          "000067",
				"12":          "120718",
				"13":          "0402",
				"17":          "0402",
				"18":          "6010",
				"3":           "010000",
				"35":          "5222222222222115",
				"37":          "C0930107FT4B",
				"38":          "206560",
				"39":          "00",
				"4":           "000000002000",
				"41":          "60101234········",
				"42":          "123456789123456",
				"44":          "0002985C0000000000000062315",
				"49":          "985",
				"7":           "04021008270402100828",
				"messageType": "Auth response",
			},
			"{\"11\":\"000067\",\"12\":\"120718\",\"13\":\"0402\",\"17\":\"0402\",\"18\":\"6010\",\"3\":\"010000\",\"35\":\"5222222222222115\",\"37\":\"C0930107FT4B\",\"38\":\"206560\",\"39\":\"00\",\"4\":\"000000002000\",\"41\":\"60101234········\",\"42\":\"123456789123456\",\"44\":\"0002985C0000000000000062315\",\"49\":\"985\",\"7\":\"04021008270402100828\",\"messageType\":\"Auth response\"}",
		},
		{ // ATM Balance Inquiry
			"000100367AC40128E1A00A30000000000000000000000000000004031009240403100925610000000000061208590403201704030403601105106123456165222222222222115C0940107GRIF60111234········123456789123456Trevica ACQ ATM·······Warszawa········PL0106100510001985985021101001000050061601864MCC000123456",
			map[string]string{
				"10":          "61000000",
				"11":          "000006",
				"12":          "120859",
				"13":          "0403",
				"15":          "20170403",
				"17":          "0403",
				"18":          "6011",
				"22":          "051",
				"3":           "300000",
				"32":          "123456",
				"35":          "5222222222222115",
				"37":          "C0940107GRIF",
				"4":           "000000000000",
				"41":          "60111234········",
				"42":          "123456789123456",
				"43":          "Trevica ACQ ATM·······Warszawa········PL",
				"48":          "6100510001",
				"49":          "985",
				"51":          "985",
				"6":           "000000000000",
				"61":          "101001000050061601864",
				"63":          "MCC000123456",
				"7":           "04031009240403100925",
				"messageType": "Auth request",
			},
			"{\"10\":\"61000000\",\"11\":\"000006\",\"12\":\"120859\",\"13\":\"0403\",\"15\":\"20170403\",\"17\":\"0403\",\"18\":\"6011\",\"22\":\"051\",\"3\":\"300000\",\"32\":\"123456\",\"35\":\"5222222222222115\",\"37\":\"C0940107GRIF\",\"4\":\"000000000000\",\"41\":\"60111234········\",\"42\":\"123456789123456\",\"43\":\"Trevica ACQ ATM·······Warszawa········PL\",\"48\":\"6100510001\",\"49\":\"985\",\"51\":\"985\",\"6\":\"000000000000\",\"61\":\"101001000050061601864\",\"63\":\"MCC000123456\",\"7\":\"04031009240403100925\",\"messageType\":\"Auth request\"}",
		},
		{ // ATM Balance Inquiry
			"0001103238C0002ED0800030000000000000000004031009240403100925000006120859040304036011165222222222222115C0940107GRIF8529630060111234········1234567891234560002985C0000000000000002285985",
			map[string]string{

				"11":          "000006",
				"12":          "120859",
				"13":          "0403",
				"17":          "0403",
				"18":          "6011",
				"3":           "300000",
				"35":          "5222222222222115",
				"37":          "C0940107GRIF",
				"38":          "852963",
				"39":          "00",
				"4":           "000000000000",
				"41":          "60111234········",
				"42":          "123456789123456",
				"44":          "0002985C0000000000000002285",
				"49":          "985",
				"7":           "04031009240403100925",
				"messageType": "Auth response",
			},
			"{\"11\":\"000006\",\"12\":\"120859\",\"13\":\"0403\",\"17\":\"0403\",\"18\":\"6011\",\"3\":\"300000\",\"35\":\"5222222222222115\",\"37\":\"C0940107GRIF\",\"38\":\"852963\",\"39\":\"00\",\"4\":\"000000000000\",\"41\":\"60111234········\",\"42\":\"123456789123456\",\"44\":\"0002985C0000000000000002285\",\"49\":\"985\",\"7\":\"04031009240403100925\",\"messageType\":\"Auth response\"}",
		},
		{ // MoneySend
			"000100B67AC40128E0A00A000000000010000028000000000000160000000000160004021207320402120733610000000000371406480402201704020402653601106123456165222222222222115C0930107GFK759991234········123456789123456Trevica···············Warszawa········PL985985021203000000010061601864MCC000123456SenderReceiverData",
			map[string]string{
				"10":          "61000000",
				"108":         "",
				"11":          "000037",
				"12":          "140648",
				"13":          "0402",
				"15":          "20170402",
				"17":          "0402",
				"18":          "6536",
				"22":          "011",
				"3":           "280000",
				"32":          "123456",
				"35":          "5222222222222115",
				"37":          "C0930107GFK7",
				"4":           "000000001600",
				"41":          "59991234········",
				"42":          "123456789123456",
				"43":          "Trevica···············Warszawa········PL",
				"49":          "985",
				"51":          "985",
				"6":           "000000001600",
				"61":          "203000000010061601864",
				"63":          "MCC000123456",
				"7":           "04021207320402120733",
				"messageType": "Auth request",
			},
			"{\"10\":\"61000000\",\"108\":\"\",\"11\":\"000037\",\"12\":\"140648\",\"13\":\"0402\",\"15\":\"20170402\",\"17\":\"0402\",\"18\":\"6536\",\"22\":\"011\",\"3\":\"280000\",\"32\":\"123456\",\"35\":\"5222222222222115\",\"37\":\"C0930107GFK7\",\"4\":\"000000001600\",\"41\":\"59991234········\",\"42\":\"123456789123456\",\"43\":\"Trevica···············Warszawa········PL\",\"49\":\"985\",\"51\":\"985\",\"6\":\"000000001600\",\"61\":\"203000000010061601864\",\"63\":\"MCC000123456\",\"7\":\"04021207320402120733\",\"messageType\":\"Auth request\"}",
		},
		{ // MoneySend
			"0001103238C0002ED0800028000000000000160004021207320402120733000037140648040204026536165222222222222115C0930107GFK73984570059991234········1234567891234560002985C0000000000000015462985",
			map[string]string{

				"11":          "000037",
				"12":          "140648",
				"13":          "0402",
				"17":          "0402",
				"18":          "6536",
				"3":           "280000",
				"35":          "5222222222222115",
				"37":          "C0930107GFK7",
				"38":          "398457",
				"39":          "00",
				"4":           "000000001600",
				"41":          "59991234········",
				"42":          "123456789123456",
				"44":          "0002985C0000000000000015462",
				"49":          "985",
				"7":           "04021207320402120733",
				"messageType": "Auth response",
			},
			"{\"11\":\"000037\",\"12\":\"140648\",\"13\":\"0402\",\"17\":\"0402\",\"18\":\"6536\",\"3\":\"280000\",\"35\":\"5222222222222115\",\"37\":\"C0930107GFK7\",\"38\":\"398457\",\"39\":\"00\",\"4\":\"000000001600\",\"41\":\"59991234········\",\"42\":\"123456789123456\",\"44\":\"0002985C0000000000000015462\",\"49\":\"985\",\"7\":\"04021207320402120733\",\"messageType\":\"Auth response\"}",
		},
		{ // Advice message
			"000120367AC4012EC1A00A00000000000000100000000000100004021018480402101849610000000000281213060402201704020402599990106123456165222222222222115C0930107FT695118700012345678········1234567891234560106100510001985985021000000000020061601864MCC000123456",
			map[string]string{
				"10":          "61000000",
				"11":          "000028",
				"12":          "121306",
				"13":          "0402",
				"15":          "20170402",
				"17":          "0402",
				"18":          "5999",
				"22":          "901",
				"3":           "000000",
				"32":          "123456",
				"35":          "5222222222222115",
				"37":          "C0930107FT69",
				"38":          "511870",
				"39":          "00",
				"4":           "000000001000",
				"41":          "12345678········",
				"42":          "123456789123456",
				"48":          "6100510001",
				"49":          "985",
				"51":          "985",
				"6":           "000000001000",
				"61":          "000000000020061601864",
				"63":          "MCC000123456",
				"7":           "04021018480402101849",
				"messageType": "Advice request",
			},
			"{\"10\":\"61000000\",\"11\":\"000028\",\"12\":\"121306\",\"13\":\"0402\",\"15\":\"20170402\",\"17\":\"0402\",\"18\":\"5999\",\"22\":\"901\",\"3\":\"000000\",\"32\":\"123456\",\"35\":\"5222222222222115\",\"37\":\"C0930107FT69\",\"38\":\"511870\",\"39\":\"00\",\"4\":\"000000001000\",\"41\":\"12345678········\",\"42\":\"123456789123456\",\"48\":\"6100510001\",\"49\":\"985\",\"51\":\"985\",\"6\":\"000000001000\",\"61\":\"000000000020061601864\",\"63\":\"MCC000123456\",\"7\":\"04021018480402101849\",\"messageType\":\"Advice request\"}",
		},
		{ // Advice message
			"000130322040002A808000000000000000001000040210184804021018490000285999165222222222222115C0930107FT690012345678········985",
			map[string]string{
				"11":          "000028",
				"18":          "5999",
				"3":           "000000",
				"35":          "5222222222222115",
				"37":          "C0930107FT69",
				"39":          "00",
				"4":           "000000001000",
				"41":          "12345678········",
				"49":          "985",
				"7":           "04021018480402101849",
				"messageType": "Advice response",
			},
			"{\"11\":\"000028\",\"18\":\"5999\",\"3\":\"000000\",\"35\":\"5222222222222115\",\"37\":\"C0930107FT69\",\"39\":\"00\",\"4\":\"000000001000\",\"41\":\"12345678········\",\"49\":\"985\",\"7\":\"04021018480402101849\",\"messageType\":\"Advice response\"}",
		},
		{ // Negative Advice message
			"000120367AC4012EC1A00A00000000000000100000000000100004021018480402101849610000000000281213060402201704020402599990106123456165222222222222115C0930107FT695118705112345678········1234567891234560106100510001985985021000000000020061601864MCC000123456",
			map[string]string{
				"10":          "61000000",
				"11":          "000028",
				"12":          "121306",
				"13":          "0402",
				"15":          "20170402",
				"17":          "0402",
				"18":          "5999",
				"22":          "901",
				"3":           "000000",
				"32":          "123456",
				"35":          "5222222222222115",
				"37":          "C0930107FT69",
				"38":          "511870",
				"39":          "51",
				"4":           "000000001000",
				"41":          "12345678········",
				"42":          "123456789123456",
				"48":          "6100510001",
				"49":          "985",
				"51":          "985",
				"6":           "000000001000",
				"61":          "000000000020061601864",
				"63":          "MCC000123456",
				"7":           "04021018480402101849",
				"messageType": "Advice request",
			},
			"{\"10\":\"61000000\",\"11\":\"000028\",\"12\":\"121306\",\"13\":\"0402\",\"15\":\"20170402\",\"17\":\"0402\",\"18\":\"5999\",\"22\":\"901\",\"3\":\"000000\",\"32\":\"123456\",\"35\":\"5222222222222115\",\"37\":\"C0930107FT69\",\"38\":\"511870\",\"39\":\"51\",\"4\":\"000000001000\",\"41\":\"12345678········\",\"42\":\"123456789123456\",\"48\":\"6100510001\",\"49\":\"985\",\"51\":\"985\",\"6\":\"000000001000\",\"61\":\"000000000020061601864\",\"63\":\"MCC000123456\",\"7\":\"04021018480402101849\",\"messageType\":\"Advice request\"}",
		},
		{ // Negative Advice message
			"000130322040002A808000000000000000001000040210184804021018490000285999165222222222222115C0930107FT690012345678········985",
			map[string]string{
				"11":          "000028",
				"18":          "5999",
				"3":           "000000",
				"35":          "5222222222222115",
				"37":          "C0930107FT69",
				"39":          "00",
				"4":           "000000001000",
				"41":          "12345678········",
				"49":          "985",
				"7":           "04021018480402101849",
				"messageType": "Advice response",
			},
			"{\"11\":\"000028\",\"18\":\"5999\",\"3\":\"000000\",\"35\":\"5222222222222115\",\"37\":\"C0930107FT69\",\"39\":\"00\",\"4\":\"000000001000\",\"41\":\"12345678········\",\"49\":\"985\",\"7\":\"04021018480402101849\",\"messageType\":\"Advice response\"}",
		},
		{ // Adjustment Advice message
			"000120367AC4012EC1A00A22000000000002000000000002000010191052481019105249610000000000361250191019201710191019599990106123456165222222222222115C0930107U4KB5118700012345678········1234567891234560106100510001985985021000000000020061601864MCC000123456",
			map[string]string{
				"10":          "61000000",
				"11":          "000036",
				"12":          "125019",
				"13":          "1019",
				"15":          "20171019",
				"17":          "1019",
				"18":          "5999",
				"22":          "901",
				"3":           "220000",
				"32":          "123456",
				"35":          "5222222222222115",
				"37":          "C0930107U4KB",
				"38":          "511870",
				"39":          "00",
				"4":           "000000020000",
				"41":          "12345678········",
				"42":          "123456789123456",
				"48":          "6100510001",
				"49":          "985",
				"51":          "985",
				"6":           "000000020000",
				"61":          "000000000020061601864",
				"63":          "MCC000123456",
				"7":           "10191052481019105249",
				"messageType": "Advice request",
			},
			"{\"10\":\"61000000\",\"11\":\"000036\",\"12\":\"125019\",\"13\":\"1019\",\"15\":\"20171019\",\"17\":\"1019\",\"18\":\"5999\",\"22\":\"901\",\"3\":\"220000\",\"32\":\"123456\",\"35\":\"5222222222222115\",\"37\":\"C0930107U4KB\",\"38\":\"511870\",\"39\":\"00\",\"4\":\"000000020000\",\"41\":\"12345678········\",\"42\":\"123456789123456\",\"48\":\"6100510001\",\"49\":\"985\",\"51\":\"985\",\"6\":\"000000020000\",\"61\":\"000000000020061601864\",\"63\":\"MCC000123456\",\"7\":\"10191052481019105249\",\"messageType\":\"Advice request\"}",
		},
		{ // Adjustment Advice message
			"000130322040002A808000220000000000020000101910524810191052490000365999165222222222222115C0930107U4KB0012345678········985",
			map[string]string{
				"11":          "000036",
				"18":          "5999",
				"3":           "220000",
				"35":          "5222222222222115",
				"37":          "C0930107U4KB",
				"39":          "00",
				"4":           "000000020000",
				"41":          "12345678········",
				"49":          "985",
				"7":           "10191052481019105249",
				"messageType": "Advice response",
			},
			"{\"11\":\"000036\",\"18\":\"5999\",\"3\":\"220000\",\"35\":\"5222222222222115\",\"37\":\"C0930107U4KB\",\"39\":\"00\",\"4\":\"000000020000\",\"41\":\"12345678········\",\"49\":\"985\",\"7\":\"10191052481019105249\",\"messageType\":\"Advice response\"}",
		},
		{ // Reversal message
			"000420367AC4012EE1A00A00000000000000100000000000100004020941100402094111610000000000251138460402201704020402599905106123456165222222222222115C0930107FSR82064640012345678········123456789123456Trevica···············Warszawa········PL02063015MCC654321  0402985985021102510000610061601864MCC000123456",
			map[string]string{
				"10":          "61000000",
				"11":          "000025",
				"12":          "113846",
				"13":          "0402",
				"15":          "20170402",
				"17":          "0402",
				"18":          "5999",
				"22":          "051",
				"3":           "000000",
				"32":          "123456",
				"35":          "5222222222222115",
				"37":          "C0930107FSR8",
				"38":          "206464",
				"39":          "00",
				"4":           "000000001000",
				"41":          "12345678········",
				"42":          "123456789123456",
				"43":          "Trevica···············Warszawa········PL",
				"48":          "63015MCC654321  0402",
				"49":          "985",
				"51":          "985",
				"6":           "000000001000",
				"61":          "102510000610061601864",
				"63":          "MCC000123456",
				"7":           "04020941100402094111",
				"messageType": "Reversal request",
			},
			"{\"10\":\"61000000\",\"11\":\"000025\",\"12\":\"113846\",\"13\":\"0402\",\"15\":\"20170402\",\"17\":\"0402\",\"18\":\"5999\",\"22\":\"051\",\"3\":\"000000\",\"32\":\"123456\",\"35\":\"5222222222222115\",\"37\":\"C0930107FSR8\",\"38\":\"206464\",\"39\":\"00\",\"4\":\"000000001000\",\"41\":\"12345678········\",\"42\":\"123456789123456\",\"43\":\"Trevica···············Warszawa········PL\",\"48\":\"63015MCC654321  0402\",\"49\":\"985\",\"51\":\"985\",\"6\":\"000000001000\",\"61\":\"102510000610061601864\",\"63\":\"MCC000123456\",\"7\":\"04020941100402094111\",\"messageType\":\"Reversal request\"}",
		},
		{ // Reversal message
			"000430322040002A808000000000000000001000040209411004020941110000255999165222222222222115C0930107FSR80012345678········985",
			map[string]string{
				"11":          "000025",
				"18":          "5999",
				"3":           "000000",
				"35":          "5222222222222115",
				"37":          "C0930107FSR8",
				"39":          "00",
				"4":           "000000001000",
				"41":          "12345678········",
				"49":          "985",
				"7":           "04020941100402094111",
				"messageType": "Reversal response",
			},
			"{\"11\":\"000025\",\"18\":\"5999\",\"3\":\"000000\",\"35\":\"5222222222222115\",\"37\":\"C0930107FSR8\",\"39\":\"00\",\"4\":\"000000001000\",\"41\":\"12345678········\",\"49\":\"985\",\"7\":\"04020941100402094111\",\"messageType\":\"Reversal response\"}",
		},
		{ // Partial Reversals message
			"000420B67AC4012EE1A00A000000020000000000000000000000150000000000150004020950370402095038610000000000271149250402201704020402599905106123456165222222222222115C0930107FSUN2065020059991234········123456789123456Trevica···············Warszawa········PL02063015MCC654321  0402985985021102510000610061601864MCC000123456 000000001200000000001200000000001200000000",
			map[string]string{
				"10":          "61000000",
				"11":          "000027",
				"12":          "114925",
				"13":          "0402",
				"15":          "20170402",
				"17":          "0402",
				"18":          "5999",
				"22":          "051",
				"3":           "000000",
				"32":          "123456",
				"35":          "5222222222222115",
				"37":          "C0930107FSUN",
				"38":          "206502",
				"39":          "00",
				"4":           "000000001500",
				"41":          "59991234········",
				"42":          "123456789123456",
				"43":          "Trevica···············Warszawa········PL",
				"48":          "63015MCC654321  0402",
				"49":          "985",
				"51":          "985",
				"6":           "000000001500",
				"61":          "102510000610061601864",
				"63":          "MCC000123456",
				"7":           "04020950370402095038",
				"95":          " 00000000120000000000120000000000120000000",
				"messageType": "Reversal request",
			},
			"{\"10\":\"61000000\",\"11\":\"000027\",\"12\":\"114925\",\"13\":\"0402\",\"15\":\"20170402\",\"17\":\"0402\",\"18\":\"5999\",\"22\":\"051\",\"3\":\"000000\",\"32\":\"123456\",\"35\":\"5222222222222115\",\"37\":\"C0930107FSUN\",\"38\":\"206502\",\"39\":\"00\",\"4\":\"000000001500\",\"41\":\"59991234········\",\"42\":\"123456789123456\",\"43\":\"Trevica···············Warszawa········PL\",\"48\":\"63015MCC654321  0402\",\"49\":\"985\",\"51\":\"985\",\"6\":\"000000001500\",\"61\":\"102510000610061601864\",\"63\":\"MCC000123456\",\"7\":\"04020950370402095038\",\"95\":\" 00000000120000000000120000000000120000000\",\"messageType\":\"Reversal request\"}",
		},
		{ // Partial Reversals message
			"000430322040002A808000000000000000001500040209503704020950380000275999165222222222222115C0930107FSR80012345678········985",
			map[string]string{
				"11":          "000027",
				"18":          "5999",
				"3":           "000000",
				"35":          "5222222222222115",
				"37":          "C0930107FSR8",
				"39":          "00",
				"4":           "000000001500",
				"41":          "12345678········",
				"49":          "985",
				"7":           "04020950370402095038",
				"messageType": "Reversal response",
			},
			"{\"11\":\"000027\",\"18\":\"5999\",\"3\":\"000000\",\"35\":\"5222222222222115\",\"37\":\"C0930107FSR8\",\"39\":\"00\",\"4\":\"000000001500\",\"41\":\"12345678········\",\"49\":\"985\",\"7\":\"04020950370402095038\",\"messageType\":\"Reversal response\"}",
		},
	}
)
