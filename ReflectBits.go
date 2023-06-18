package Cryptx

func Reflect64Bits(data uint64) uint64 {
	const width = 64
	reflected := uint64(0)
	for i := 0; i < width; i++ {
		reflected <<= 1
		if data&1 != 0 {
			reflected |= 1
		}
		data >>= 1
	}
	return reflected
}

func Reflect32Bits(data uint32) uint32 {
	const width = 32
	var reflected uint32
	for i := 0; i < width; i++ {
		if data&1 != 0 {
			reflected |= 1 << (width - 1 - i)
		}
		data >>= 1
	}
	return reflected
}

func Reflect16Bits(data uint16) uint16 {
	const width = 16
	var reflected uint16
	for i := 0; i < width; i++ {
		if data&1 != 0 {
			reflected |= 1 << (width - 1 - i)
		}
		data >>= 1
	}
	return reflected
}
