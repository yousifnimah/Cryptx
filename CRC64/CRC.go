package CRC64

import "fmt"

type CRC struct {
	Data            []byte
	Polynomial      uint64
	InitialValue    uint64
	XOROUT          uint64
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

func Result(Data []byte, Algorithm string) uint64 {
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

func (CRC *CRC) Init() uint64 {
	CRC.SetValues()
	return CRC.Calculate()
}

func (CRC CRC) Calculate() uint64 {
	var crc uint64 = CRC.InitialValue
	Data := CRC.Data

	for _, b := range Data {
		crc = (crc << 8) ^ uint64(b)
		for i := 0; i < 8; i++ {
			if (crc & 1) != 0 {
				crc = (crc >> 1) ^ CRC.Polynomial
			} else {
				crc >>= 1
			}
		}
	}

	crc ^= CRC.XOROUT
	return crc
}
