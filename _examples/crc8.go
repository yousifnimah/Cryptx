package main

import (
	"Cryptx/CRC8"
	"fmt"
)

func main() {
	//CRC-8/ITU
	Input := []byte("12345")
	checksum := CRC8.Result(Input, "ITU")
	checksumHex := CRC8.ResultHex(Input, "ITU")
	fmt.Println(checksum, checksumHex)
}
