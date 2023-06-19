package CRC8

import (
	fmt "fmt"
	"github.com/yousifnimah/Cryptx"
)

type CRC struct {
	Data            []byte
	Polynomial      byte
	InitialValue    byte
	XOROUT          byte
	InputReflected  bool
	OutputReflected bool
	Algorithm       string
}

func New(Data []byte, Algorithm string) CRC {
	C := CRC{
		Data:      Data,
		Algorithm: Algorithm,
	}
	return C
}

func Result(Data []byte, Algorithm string) byte {
	C := &CRC{
		Data:      Data,
		Algorithm: Algorithm,
	}
	return C.Init()
}

func ResultHex(Data []byte, Algorithm string) string {
	C := &CRC{
		Data:      Data,
		Algorithm: Algorithm,
	}
	Checksum := C.Init()
	return fmt.Sprintf("0x%X", Checksum)
}

func (CRC *CRC) Init() byte {
	CRC.SetValues()
	return CRC.Calculate()
}

func (CRC CRC) Calculate() byte {
	var crc byte = CRC.InitialValue
	Data := CRC.Data

	if CRC.InputReflected {
		// Reflect input bytes
		for i := 0; i < len(Data); i++ {
			Data[i] = Cryptx.ReflectByte(Data[i])
		}
	}

	for _, b := range Data {
		crc ^= b

		for i := 0; i < 8; i++ {
			if crc&0x80 != 0 {
				crc = (crc << 1) ^ CRC.Polynomial
			} else {
				crc <<= 1
			}
		}
	}

	crc ^= CRC.XOROUT

	if CRC.OutputReflected {
		crc = Cryptx.ReflectByte(crc)
	}
	return crc
}
