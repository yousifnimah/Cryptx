## Golang CRC (Cyclic Redundancy Check)

CRC (Cyclic Redundancy Check) is an error detection algorithm that generates a checksum based on the data being
transmitted. In Go (Golang), you can implement CRC using various algorithms, such as CRC-8, CRC-16, CRC-32 or CRC-64.

## Table of Contents

- [Installation](#Installation)
- [Get started](#get-started-for-free)
  - [CRC-8 Usage](#CRC-8 Usage)
- [Supported Algorithms](#get-started-for-free)

## Installation

```go
go install github.com/yousifnimah/Golang-CRC
```

## Get Started

- ### CRC-8 Usage

  1 - Importing package into main package

    ```go
    import ("Cryptx/CRC8")
    ```

  2 - Use algorithm name from the table below:

  | Algorithm         | Polynomial | Initial Value | XOROUT |
  |:-----------------:|:----------:|:-------------:|:------:|
  | CRC-8             |    0x07    |     0x00      |  0x00  |
  | ITU               |    0x07    |     0x00      |  0x55  |
  | ROHC              |    0x07    |     0xFF      |  0x00  |
  | SAE-J1850         |    0x1D    |     0xFF      |  0xFF  |
  | SAE-J1850_ZERO    |    0x1D    |     0x00      |  0x00  |
  | 8H2F              |    0x2F    |     0xFF      |  0xFF  |
  | CDMA2000          |    0x9B    |     0xFF      |  0x00  |
  | DARC              |    0x39    |     0x00      |  0x00  |
  | DVB_S2            |    0xD5    |     0x00      |  0x00  |
  | EBU               |    0x1D    |     0xFF      |  0x00  |
  | ICODE             |    0x1D    |     0xFD      |  0x00  |
  | MAXIM             |    0x31    |     0x00      |  0x00  |
  | WCDMA             |    0x9B    |     0x00      |  0x00  |


  In main.go:

  ```go
  package main

  import (
     "Cryptx/CRC8"
     "fmt"
  )

  func main() {
        Input := []byte("12345") //string to slice of bytes
        AlgorithmName := "ITU" //CRC-8 algorithm name from supported table
        checksumHex := CRC8.ResultHex(Input, AlgorithmName)
        fmt.Println("Output:", checksumHex)
  }
  ```

  Results:
  ```
  Output: 0x9E
  ```
  
  ### You can get output in byte:

  ```go
    Input := []byte("12345")
    checksumByte := CRC8.Result(Input, "ITU")
    fmt.Println("Output:", checksumByte)
  ```