package CRC32

import (
	"Cryptx"
	"fmt"
)

type CRC struct {
	Data            []byte
	Polynomial      uint32
	InitialValue    uint32
	XOROUT          uint32
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

func Result(Data []byte, Algorithm string) uint32 {
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

func (CRC *CRC) Init() uint32 {
	CRC.SetValues()
	return CRC.Calculate()
}

func (CRC CRC) Calculate() uint32 {
	var crc uint32 = CRC.InitialValue
	Data := CRC.Data

	if CRC.InputReflected {
		// Reflect input bytes
		for i := 0; i < len(Data); i++ {
			Data[i] = Cryptx.ReflectByte(Data[i])
		}
	}

	for _, b := range Data {
		crc ^= uint32(b) << 24

		for i := 0; i < 8; i++ {
			if crc&0x80000000 != 0 {
				crc = (crc << 1) ^ CRC.Polynomial
			} else {
				crc <<= 1
			}
		}
	}

	crc ^= CRC.XOROUT

	if CRC.OutputReflected {
		crc = Cryptx.Reflect32Bits(crc)
	}
	return crc
}
