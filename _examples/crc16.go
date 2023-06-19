package main

import (
	"fmt"
	"github.com/yousifnimah/Cryptx/CRC16"
)

func main() {
	//CRC-16/CCIT_ZERO
	Input := []byte("12345")
	checksum := CRC16.Result(Input, "CCIT_ZERO")
	checksumHex := CRC16.ResultHex(Input, "CCIT_ZERO")
	fmt.Println(checksum, checksumHex)
}
