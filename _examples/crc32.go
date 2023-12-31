package main

import (
	"fmt"
	"github.com/yousifnimah/Cryptx/CRC32"
)

func main() {
	//CRC-32/BZIP2
	Input := []byte("12345")
	checksum := CRC32.Result(Input, "BZIP2")
	checksumHex := CRC32.ResultHex(Input, "BZIP2")
	fmt.Println(checksum, checksumHex)
}
