package main

import (
	"fmt"
	"os"
	"strings"
	"time"
	"unsafe"
)

// This program will try different compression algorithms:
// - Huffman,
// - Run-length encoding
// -

func main() {

	if len(os.Args) >= 3 {
		filePath := os.Args[1]
		compressor := os.Args[2]

		image := readImage(filePath)

		// Measure algorithm encoding phase
		startTime := time.Now()

		if strings.Contains(strings.ToLower(compressor), "huff") {
			fmt.Println("Compressor chosen: Huffman")
			encodedImage := huffmanCompression(image)

			//fmt.Println(encodedImage)
			fmt.Println(len(encodedImage))
			fmt.Println("SizeOfofimage: ", (unsafe.Sizeof([114976]byte{})))

			gain := float64(unsafe.Sizeof([266 * 425]uint32{})) / float64(unsafe.Sizeof([114976]byte{}))
			fmt.Println("Gain: ", gain)
		} else {
			fmt.Println("Compressor ", compressor, " does not match any existing compressor")
		}

		elapsedTime := time.Since(startTime)
		fmt.Println("Time for encoding: ", elapsedTime)
	} else {
		fmt.Println("Missing parameters. shall be: compression imageFilePath algorithm")
		fmt.Println("Note: algorithm can be: \n\t- Huffman\n\t- TODO")
	}
}
