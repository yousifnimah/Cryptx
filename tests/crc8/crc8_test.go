package crc8

import (
	"Cryptx/CRC8"
	"fmt"
	"os"
	"testing"
)

func TestCRC8Result(t *testing.T) {
	input := []byte("12345")
	expectedChecksum := byte(0x9E)
	checksum := CRC8.Result(input, "ITU")

	if checksum != expectedChecksum {
		t.Errorf("Expected checksum: %x, but got: %x", expectedChecksum, checksum)
	}
}

func TestCRC8ResultHex(t *testing.T) {
	input := []byte("12345")
	expectedChecksumHex := "0x9E"
	checksumHex := CRC8.ResultHex(input, "ITU")

	if checksumHex != expectedChecksumHex {
		t.Errorf("Expected checksumHex: %s, but got: %s", expectedChecksumHex, checksumHex)
	}
}

func TestMain(m *testing.M) {
	fmt.Println("Running CRC8 package tests...")
	result := m.Run()
	fmt.Println("Tests finished.")
	os.Exit(result)
}
