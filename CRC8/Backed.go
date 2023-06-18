package CRC8

func (CRC *CRC) SetValues() {
	switch CRC.Algorithm {
	case "CRC-8":
		CRC.InitialValue = 0x00
		CRC.XOROUT = 0x00
		CRC.InputReflected = false
		CRC.OutputReflected = false
		CRC.Polynomial = 0x07
		break
	case "ITU":
		CRC.InitialValue = 0x0
		CRC.XOROUT = 0x55
		CRC.InputReflected = false
		CRC.OutputReflected = false
		CRC.Polynomial = 0x07
		break
	case "ROHC":
		CRC.InitialValue = 0xFF
		CRC.XOROUT = 0x00
		CRC.InputReflected = true
		CRC.OutputReflected = true
		CRC.Polynomial = 0x07
		break
	case "SAE-J1850":
		CRC.InitialValue = 0xFF
		CRC.XOROUT = 0xFF
		CRC.InputReflected = false
		CRC.OutputReflected = false
		CRC.Polynomial = 0x1D
		break
	case "SAE-J1850_ZERO":
		CRC.InitialValue = 0x0
		CRC.XOROUT = 0x0
		CRC.InputReflected = false
		CRC.OutputReflected = false
		CRC.Polynomial = 0x1D
		break
	case "8H2F":
		CRC.InitialValue = 0xFF
		CRC.XOROUT = 0xFF
		CRC.InputReflected = false
		CRC.OutputReflected = false
		CRC.Polynomial = 0x2F
		break
	case "CDMA2000":
		CRC.InitialValue = 0xFF
		CRC.XOROUT = 0x0
		CRC.InputReflected = false
		CRC.OutputReflected = false
		CRC.Polynomial = 0x9B
		break
	case "DARC":
		CRC.InitialValue = 0x00
		CRC.XOROUT = 0x0
		CRC.InputReflected = false
		CRC.OutputReflected = false
		CRC.Polynomial = 0x39
		break
	case "DVB_S2":
		CRC.InitialValue = 0x0
		CRC.XOROUT = 0x0
		CRC.InputReflected = false
		CRC.OutputReflected = false
		CRC.Polynomial = 0xD5
		break
	case "EBU":
		CRC.InitialValue = 0xFF
		CRC.XOROUT = 0x0
		CRC.InputReflected = true
		CRC.OutputReflected = true
		CRC.Polynomial = 0x1D
		break
	case "ICODE":
		CRC.InitialValue = 0xFD
		CRC.XOROUT = 0x0
		CRC.InputReflected = false
		CRC.OutputReflected = false
		CRC.Polynomial = 0x1D
		break
	case "MAXIM":
		CRC.InitialValue = 0x0
		CRC.XOROUT = 0x0
		CRC.InputReflected = true
		CRC.OutputReflected = true
		CRC.Polynomial = 0x31
		break
	case "WCDMA":
		CRC.InitialValue = 0x0
		CRC.XOROUT = 0x0
		CRC.InputReflected = true
		CRC.OutputReflected = true
		CRC.Polynomial = 0x9B
		break
	default:
		CRC.InitialValue = 0x00
		CRC.XOROUT = 0x00
		CRC.InputReflected = false
		CRC.OutputReflected = false
		CRC.Polynomial = 0x07
		break
	}
}
