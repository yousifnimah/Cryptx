## Golang CRC (Cyclic Redundancy Check)

CRC (Cyclic Redundancy Check) is an error detection algorithm that generates a checksum based on the data being
transmitted. In Go (Golang), you can implement CRC using various algorithms, such as CRC-8, CRC-16, CRC-32 or CRC-64.

## Table of Contents

- [Installation](#Installation)
- [Get started](#Get-Started)
    - [CRC-8 Usage](#CRC-8-Usage)
    - [CRC-16 Usage](#CRC-16-Usage)
    - [CRC-32 Usage](#CRC-32-Usage)
    - [CRC-64 Usage](#CRC-64-Usage)

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
      AlgorithmName := "ITU"   //CRC-8 algorithm name from supported table
      checksumHex := CRC8.ResultHex(Input, AlgorithmName)
      fmt.Println("Output:", checksumHex)
  }
    ```

  Result:

    ```
    Output: 0x9E
    ```

  #### You can get output in byte:

    ```go
    Input := []byte("12345")
    checksumByte := CRC8.Result(Input, "ITU")
    fmt.Println("Output:", checksumByte)
    ```

</br>

- ## CRC-16 Usage

  1 - Importing package

    ```go
    import ("Cryptx/CRC16")
    ```

  2 - Use algorithm name from the table below:

  | Algorithm   | Polynomial | InitialValue | XOROUT |
  |-------------|------------|--------------|--------|
  | CCIT_ZERO   | 0x1021     | 0x0000       | 0x0000 |
  | ARC         | 0x8005     | 0x0000       | 0x0000 |
  | AUG_CCITT   | 0x1021     | 0x1D0F       | 0x0000 |
  | BUYPASS     | 0x8005     | 0x0000       | 0x0000 |
  | CCITT_FALSE | 0x1021     | 0xFFFF       | 0x0000 |
  | CDMA2000    | 0xC867     | 0xFFFF       | 0xFFFF |
  | DDS_110     | 0x8005     | 0x800D       | 0x0000 |
  | DECT_R      | 0x589      | 0x0000       | 0x0001 |
  | DECT_X      | 0x589      | 0x0000       | 0x0000 |
  | DNP         | 0x3D65     | 0x0000       | 0xFFFF |
  | EN_13757    | 0x3D65     | 0x0000       | 0xFFFF |
  | GENIBUS     | 0x1021     | 0xFFFF       | 0xFFFF |
  | MAXIM       | 0x8005     | 0x0000       | 0xFFFF |
  | MCRF4XX     | 0x1021     | 0xFFFF       | 0x0000 |
  | RIELLO      | 0x1021     | 0xB2AA       | 0x0000 |
  | T10_DIF     | 0x8BB7     | 0x0000       | 0x0000 |
  | TELEDISK    | 0xA097     | 0x0000       | 0x0000 |
  | TMS37157    | 0x1021     | 0x89EC       | 0x0000 |
  | USB         | 0x8005     | 0xFFFF       | 0xFFFF |
  | A           | 0x1021     | 0xC6C6       | 0x0000 |
  | KERMIT      | 0x1021     | 0x0000       | 0x0000 |
  | MODBUS      | 0x8005     | 0xFFFF       | 0x0000 |
  | X_25        | 0x1021     | 0xFFFF       | 0xFFFF |
  | XMODEM      | 0x1021     | 0x0000       | 0x0000 |
  | IBM         | 0x8005     | 0x0000       | 0x0000 |

  In main.go:

    ```go
    package main
  
  import (
      "Cryptx/CRC16"
      "fmt"
  )
  
  func main() {
      Input := []byte("12345")     //string to slice of bytes
      AlgorithmName := "CCIT_ZERO" //Algorithm name from supported table
      checksumHex := CRC16.ResultHex(Input, AlgorithmName)
      fmt.Println("Output:", checksumHex)
  }
    ```

  Result:

    ```
    Output: 0x546C
    ```

- ## CRC-32 Usage

  1 - Importing package

    ```go
    import ("Cryptx/CRC32")
    ```

  2 - Use algorithm name from the table below:

  | Algorithm | Polynomial | InitialValue |   XOROUT   |
  |-----------|------------|--------------|------------|
  | CRC-32    | 0x4C11DB7  | 0xFFFFFFFF   | 0xFFFFFFFF |
  | BZIP2     | 0x4C11DB7  | 0xFFFFFFFF   | 0xFFFFFFFF |
  | C         | 0x1EDC6F41 | 0xFFFFFFFF   | 0xFFFFFFFF |
  | D         | 0xA833982B | 0xFFFFFFFF   | 0xFFFFFFFF |
  | MPEG2     | 0x4C11DB7  | 0xFFFFFFFF   | 0x0        |
  | POSIX     | 0x4C11DB7  | 0x0          | 0xFFFFFFFF |
  | Q         | 0x814141AB | 0x0          | 0x0        |
  | JAMCRC    | 0x4C11DB7  | 0xFFFFFFFF   | 0x0        |
  | XFER      | 0xAF       | 0x0          | 0x0        |
  | KOOPM     | 0x741B8CD7 | 0xFFFFFFFF   | 0xFFFFFFFF |

  In main.go:

    ```go
    package main
    
    import (
       "Cryptx/CRC32"
       "fmt"
    )
    
    func main() {
          Input := []byte("12345") //string to slice of bytes
          AlgorithmName := "BZIP2" //Algorithm name from supported table
          checksumHex := CRC32.ResultHex(Input, AlgorithmName)
          fmt.Println("Output:", checksumHex)
    }
    ```

  Result:
  ```
  Output: 0x426548B8
  ```

- ## CRC-64 Usage

  1 - Importing package

    ```go
    import ("Cryptx/CRC64")
    ```

  2 - Use algorithm name from the table below:

  | Algorithm |     Polynomial      |    InitialValue     |       XOROUT        |
  |-----------|---------------------|---------------------|---------------------|
  | ECMA      | 0x42F0E1EBA9EA3693  | 0xFFFFFFFFFFFFFFFF  | 0xFFFFFFFFFFFFFFFF  |
  | WE        | 0x42F0E1EBA9EA3693  | 0xFFFFFFFFFFFFFFFF  | 0xFFFFFFFFFFFFFFFF  |

  In main.go:

    ```go
    package main
    
    import (
       "Cryptx/CRC64"
       "fmt"
    )
    
    func main() {
          Input := []byte("12345") //string to slice of bytes
          AlgorithmName := "ECMA" //Algorithm name from supported table
          checksumHex := CRC64.ResultHex(Input, AlgorithmName)
          fmt.Println("Output:", checksumHex)
    }
    ```

  Result:

  ```
  Output: 0xDDEDF5CB80FA64E9
  ```