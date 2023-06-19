package main

import (
	"fmt"
	"github.com/yousifnimah/Cryptx/CRC8"
)

func main() {
	//CRC-8/ITU
	Input := []byte("12345")
	checksum := CRC8.Result(Input, "ITU")
	checksumHex := CRC8.ResultHex(Input, "ITU")
	fmt.Println(checksum, checksumHex)
}
