package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"os"
)

// MAGIC 	: 3 bytes, 'F' or 'C', 'W', 'S', a C on the first byte states compression
// VERSION : 1 byte, i.e. 5 , 6, 7
// SIZE 		: 4 bytes, uncompressed size in LSB-MSB format ([0]+[1]<<8+[2]<<16+[3]<<32)
// If file is compressed, a GZ stream starts here (0x78 indicating DEFLATE)

// SwfHeader flash file header
type SwfHeader struct {
	Signature  [3]byte
	Version    uint8
	FileLength uint32
}

func (swfHeader SwfHeader) String() string {
	compression := GetCompression(swfHeader.Signature)
	return fmt.Sprintf("Signature: %s\nCompression: %s\nSWF Version: %d\nFile Size: %d\n", swfHeader.Signature, compression, swfHeader.Version, swfHeader.FileLength)
}

// https://wwwimages2.adobe.com/content/dam/acom/en/devnet/pdf/swf-file-format-spec.pdf
func main() {

	fmt.Println(len(os.Args))
	if len(os.Args) <= 1 {
		panic("Usage main.go file.swf")
	}

	argsWithoutProg := os.Args[1:]
	swfPath := argsWithoutProg[0]

	file, err := os.Open(swfPath)
	if err != nil {
		log.Fatal("Error while opening file", err)
	}

	defer file.Close()

	swfHeader := SwfHeader{}
	binary.Read(file, binary.LittleEndian, &swfHeader)
	fmt.Println(swfHeader)

}

// GetCompression Type of compression
func GetCompression(magic [3]byte) string {
	// magic[0] == F
	out := "uncompressed"
	switch magic[0] {
	case 'C':
		out = "ZLIB"
	case 'Z':
		out = "LZMA"
	}
	return out
}
