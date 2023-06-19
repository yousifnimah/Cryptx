package CRC16

import (
	"fmt"
	"github.com/yousifnimah/Cryptx"
)

type CRC struct {
	Data            []byte
	Polynomial      uint16
	InitialValue    uint16
	XOROUT          uint16
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

func Result(Data []byte, Algorithm string) uint16 {
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

func (CRC *CRC) Init() uint16 {
	CRC.SetValues()
	return CRC.Calculate()
}

func (CRC CRC) Calculate() uint16 {
	var crc uint16 = CRC.InitialValue
	Data := CRC.Data

	if CRC.InputReflected {
		// Reflect input bytes
		for i := 0; i < len(Data); i++ {
			Data[i] = Cryptx.ReflectByte(Data[i])
		}
	}

	for _, b := range Data {
		crc ^= uint16(b) << 8

		for i := 0; i < 8; i++ {
			if crc&0x8000 != 0 {
				crc = (crc << 1) ^ CRC.Polynomial
			} else {
				crc <<= 1
			}
		}
	}

	crc ^= CRC.XOROUT

	if CRC.OutputReflected {
		crc = Cryptx.Reflect16Bits(crc)
	}
	return crc
}
