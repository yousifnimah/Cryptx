package crc16

import (
	"Cryptx/CRC64"
	"fmt"
	"os"
	"testing"
)

func TestCRC16Result(t *testing.T) {
	input := []byte("12345")
	expectedChecksum := uint64(0xDDEDF5CB80FA64E9)
	checksum := CRC64.Result(input, "ECMA")

	if checksum != expectedChecksum {
		t.Errorf("Expected checksum: %x, but got: %x", expectedChecksum, checksum)
	}
}

func TestCRC16ResultHex(t *testing.T) {
	input := []byte("12345")
	expectedChecksumHex := "0xDDEDF5CB80FA64E9"
	checksumHex := CRC64.ResultHex(input, "ECMA")

	if checksumHex != expectedChecksumHex {
		t.Errorf("Expected checksumHex: %s, but got: %s", expectedChecksumHex, checksumHex)
	}
}

func TestMain(m *testing.M) {
	fmt.Println("Running CRC64 package tests...")
	result := m.Run()
	fmt.Println("Tests finished.")
	os.Exit(result)
}
