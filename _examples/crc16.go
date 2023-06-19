package main

import (
	"Cryptx/CRC16"
	"fmt"
)

func main() {
	//CRC-16/CCIT_ZERO
	Input := []byte("12345")
	checksum := CRC16.Result(Input, "CCIT_ZERO")
	checksumHex := CRC16.ResultHex(Input, "CCIT_ZERO")
	fmt.Println(checksum, checksumHex)
}
