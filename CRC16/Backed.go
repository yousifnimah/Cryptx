package CRC16

func (CRC *CRC) SetValues() {
	switch CRC.Algorithm {
	case "CCIT_ZERO":
		CRC.InitialValue = 0x0000
		CRC.XOROUT = 0x0000
		CRC.InputReflected = false
		CRC.OutputReflected = false
		CRC.Polynomial = 0x1021
		break
	case "ARC":
		CRC.InitialValue = 0x0000
		CRC.XOROUT = 0x0000
		CRC.InputReflected = true
		CRC.OutputReflected = true
		CRC.Polynomial = 0x8005
		break
	case "AUG_CCITT":
		CRC.InitialValue = 0x1D0F
		CRC.XOROUT = 0x0000
		CRC.InputReflected = false
		CRC.OutputReflected = false
		CRC.Polynomial = 0x1021
		break
	case "BUYPASS":
		CRC.InitialValue = 0x0000
		CRC.XOROUT = 0x0000
		CRC.InputReflected = false
		CRC.OutputReflected = false
		CRC.Polynomial = 0x8005
		break
	case "CCITT_FALSE":
		CRC.InitialValue = 0xFFFF
		CRC.XOROUT = 0x0000
		CRC.InputReflected = false
		CRC.OutputReflected = false
		CRC.Polynomial = 0x1021
		break
	case "CDMA2000":
		CRC.InitialValue = 0xFFFF
		CRC.XOROUT = 0x0000
		CRC.InputReflected = false
		CRC.OutputReflected = false
		CRC.Polynomial = 0xC867
		break
	case "DDS_110":
		CRC.InitialValue = 0x800D
		CRC.XOROUT = 0x0000
		CRC.InputReflected = false
		CRC.OutputReflected = false
		CRC.Polynomial = 0x8005
		break
	case "DECT_R":
		CRC.InitialValue = 0x0000
		CRC.XOROUT = 0x1
		CRC.InputReflected = false
		CRC.OutputReflected = false
		CRC.Polynomial = 0x589
		break
	case "DECT_X":
		CRC.InitialValue = 0x0000
		CRC.XOROUT = 0x0000
		CRC.InputReflected = false
		CRC.OutputReflected = false
		CRC.Polynomial = 0x589
		break
	case "DNP":
		CRC.InitialValue = 0x0000
		CRC.XOROUT = 0xFFFF
		CRC.InputReflected = true
		CRC.OutputReflected = true
		CRC.Polynomial = 0x3D65
		break
	case "EN_13757":
		CRC.InitialValue = 0x0000
		CRC.XOROUT = 0xFFFF
		CRC.InputReflected = false
		CRC.OutputReflected = false
		CRC.Polynomial = 0x3D65
		break
	case "GENIBUS":
		CRC.InitialValue = 0xFFFF
		CRC.XOROUT = 0xFFFF
		CRC.InputReflected = false
		CRC.OutputReflected = false
		CRC.Polynomial = 0x1021
		break
	case "MAXIM":
		CRC.InitialValue = 0x0
		CRC.XOROUT = 0xFFFF
		CRC.InputReflected = true
		CRC.OutputReflected = true
		CRC.Polynomial = 0x8005
		break
	case "MCRF4XX":
		CRC.InitialValue = 0xFFFF
		CRC.XOROUT = 0x0
		CRC.InputReflected = true
		CRC.OutputReflected = true
		CRC.Polynomial = 0x1021
		break
	case "RIELLO":
		CRC.InitialValue = 0xB2AA
		CRC.XOROUT = 0x0
		CRC.InputReflected = true
		CRC.OutputReflected = true
		CRC.Polynomial = 0x1021
		break
	case "T10_DIF":
		CRC.InitialValue = 0x0
		CRC.XOROUT = 0x0
		CRC.InputReflected = false
		CRC.OutputReflected = false
		CRC.Polynomial = 0x8BB7
		break
	case "TELEDISK":
		CRC.InitialValue = 0x0
		CRC.XOROUT = 0x0
		CRC.InputReflected = false
		CRC.OutputReflected = false
		CRC.Polynomial = 0xA097
		break
	case "TMS37157":
		CRC.InitialValue = 0x89EC
		CRC.XOROUT = 0x0
		CRC.InputReflected = true
		CRC.OutputReflected = true
		CRC.Polynomial = 0x1021
		break
	case "USB":
		CRC.InitialValue = 0xFFFF
		CRC.XOROUT = 0xFFFF
		CRC.InputReflected = true
		CRC.OutputReflected = true
		CRC.Polynomial = 0x8005
		break
	case "A":
		CRC.InitialValue = 0xC6C6
		CRC.XOROUT = 0x0
		CRC.InputReflected = true
		CRC.OutputReflected = true
		CRC.Polynomial = 0x1021
		break
	case "KERMIT":
		CRC.InitialValue = 0x0
		CRC.XOROUT = 0x0
		CRC.InputReflected = true
		CRC.OutputReflected = true
		CRC.Polynomial = 0x1021
		break
	case "MODBUS":
		CRC.InitialValue = 0xFFFF
		CRC.XOROUT = 0x0
		CRC.InputReflected = true
		CRC.OutputReflected = true
		CRC.Polynomial = 0x8005
		break
	case "X_25":
		CRC.InitialValue = 0xFFFF
		CRC.XOROUT = 0xFFFF
		CRC.InputReflected = true
		CRC.OutputReflected = true
		CRC.Polynomial = 0x1021
		break
	case "XMODEM":
		CRC.InitialValue = 0x0
		CRC.XOROUT = 0x0
		CRC.InputReflected = false
		CRC.OutputReflected = false
		CRC.Polynomial = 0x1021
		break
	case "IBM":
		CRC.InitialValue = 0x0
		CRC.XOROUT = 0x0
		CRC.InputReflected = true
		CRC.OutputReflected = true
		CRC.Polynomial = 0x8005
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
