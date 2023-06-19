package crc16

import (
	"fmt"
	"github.com/yousifnimah/Cryptx/CRC32"
	"os"
	"testing"
)

func TestCRC16Result(t *testing.T) {
	input := []byte("12345")
	expectedChecksum := uint32(0x426548B8)
	checksum := CRC32.Result(input, "BZIP2")

	if checksum != expectedChecksum {
		t.Errorf("Expected checksum: %x, but got: %x", expectedChecksum, checksum)
	}
}

func TestCRC16ResultHex(t *testing.T) {
	input := []byte("12345")
	expectedChecksumHex := "0x426548B8"
	checksumHex := CRC32.ResultHex(input, "BZIP2")

	if checksumHex != expectedChecksumHex {
		t.Errorf("Expected checksumHex: %s, but got: %s", expectedChecksumHex, checksumHex)
	}
}

func TestMain(m *testing.M) {
	fmt.Println("Running CRC32 package tests...")
	result := m.Run()
	fmt.Println("Tests finished.")
	os.Exit(result)
}
