package CRC32

func (CRC *CRC) SetValues() {
	switch CRC.Algorithm {
	case "CRC-32":
		CRC.InitialValue = 0xFFFFFFFF
		CRC.XOROUT = 0xFFFFFFFF
		CRC.InputReflected = true
		CRC.OutputReflected = true
		CRC.Polynomial = 0x4C11DB7
		break
	case "BZIP2":
		CRC.InitialValue = 0xFFFFFFFF
		CRC.XOROUT = 0xFFFFFFFF
		CRC.InputReflected = false
		CRC.OutputReflected = false
		CRC.Polynomial = 0x4C11DB7
		break
	case "C":
		CRC.InitialValue = 0xFFFFFFFF
		CRC.XOROUT = 0xFFFFFFFF
		CRC.InputReflected = true
		CRC.OutputReflected = true
		CRC.Polynomial = 0x1EDC6F41
		break
	case "D":
		CRC.InitialValue = 0xFFFFFFFF
		CRC.XOROUT = 0xFFFFFFFF
		CRC.InputReflected = true
		CRC.OutputReflected = true
		CRC.Polynomial = 0xA833982B
		break
	case "MPEG2":
		CRC.InitialValue = 0xFFFFFFFF
		CRC.XOROUT = 0x0
		CRC.InputReflected = false
		CRC.OutputReflected = false
		CRC.Polynomial = 0x4C11DB7
		break
	case "POSIX":
		CRC.InitialValue = 0x0
		CRC.XOROUT = 0xFFFFFFFF
		CRC.InputReflected = false
		CRC.OutputReflected = false
		CRC.Polynomial = 0x4C11DB7
		break
	case "Q":
		CRC.InitialValue = 0x0
		CRC.XOROUT = 0x0
		CRC.InputReflected = false
		CRC.OutputReflected = false
		CRC.Polynomial = 0x814141AB
		break
	case "JAMCRC":
		CRC.InitialValue = 0xFFFFFFFF
		CRC.XOROUT = 0x0
		CRC.InputReflected = true
		CRC.OutputReflected = true
		CRC.Polynomial = 0x4C11DB7
		break
	case "XFER":
		CRC.InitialValue = 0x0
		CRC.XOROUT = 0x0
		CRC.InputReflected = false
		CRC.OutputReflected = false
		CRC.Polynomial = 0xAF
		break
	case "KOOPM":
		CRC.InitialValue = 0xFFFFFFFF
		CRC.XOROUT = 0xFFFFFFFF
		CRC.InputReflected = true
		CRC.OutputReflected = true
		CRC.Polynomial = 0x741B8CD7
		break
	default:
		CRC.InitialValue = 0xFFFFFFFF
		CRC.XOROUT = 0xFFFFFFFF
		CRC.InputReflected = true
		CRC.OutputReflected = true
		CRC.Polynomial = 0x4C11DB7
		break
	}
}
