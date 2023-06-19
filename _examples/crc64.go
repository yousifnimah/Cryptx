package main

import (
	"Cryptx/CRC64"
	"fmt"
)

func main() {
	//CRC-64/ECMA
	Input := []byte("12345")
	checksum := CRC64.Result(Input, "ECMA")
	checksumHex := CRC64.ResultHex(Input, "ECMA")
	fmt.Println(checksum, checksumHex)
}
