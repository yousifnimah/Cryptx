package crc16

import (
	"fmt"
	"github.com/yousifnimah/Cryptx/CRC16"
	"os"
	"testing"
)

func TestCRC16Result(t *testing.T) {
	input := []byte("12345")
	expectedChecksum := uint16(0x546C)
	checksum := CRC16.Result(input, "CCIT_ZERO")

	if checksum != expectedChecksum {
		t.Errorf("Expected checksum: %x, but got: %x", expectedChecksum, checksum)
	}
}

func TestCRC16ResultHex(t *testing.T) {
	input := []byte("12345")
	expectedChecksumHex := "0x546C"
	checksumHex := CRC16.ResultHex(input, "CCIT_ZERO")

	if checksumHex != expectedChecksumHex {
		t.Errorf("Expected checksumHex: %s, but got: %s", expectedChecksumHex, checksumHex)
	}
}

func TestMain(m *testing.M) {
	fmt.Println("Running CRC16 package tests...")
	result := m.Run()
	fmt.Println("Tests finished.")
	os.Exit(result)
}
