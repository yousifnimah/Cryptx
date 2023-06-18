package CRC64

func (CRC *CRC) SetValues() {
	switch CRC.Algorithm {
	case "ECMA":
		CRC.InitialValue = 0xFFFFFFFFFFFFFFFF
		CRC.XOROUT = 0xFFFFFFFFFFFFFFFF
		CRC.InputReflected = false
		CRC.OutputReflected = false
		CRC.Polynomial = 0x42F0E1EBA9EA3693
		break
	case "WE":
		CRC.InitialValue = 0xFFFFFFFFFFFFFFFF
		CRC.XOROUT = 0xFFFFFFFFFFFFFFFF
		CRC.InputReflected = false
		CRC.OutputReflected = false
		CRC.Polynomial = 0x42F0E1EBA9EA3693
		break
	default:
		CRC.InitialValue = 0xFFFFFFFFFFFFFFFF
		CRC.XOROUT = 0xFFFFFFFFFFFFFFFF
		CRC.InputReflected = false
		CRC.OutputReflected = false
		CRC.Polynomial = 0x42F0E1EBA9EA3693
		break
	}
}
