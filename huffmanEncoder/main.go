package main

import (
	"fmt"
	"os"
	"strings"
	"time"
	"unsafe"
)

func main() {
	if len(os.Args) >= 3 {
		filePath := os.Args[1]
		compressor := os.Args[2]

		image, status := readImage(filePath)

		if status == nil {
			// Measure algorithm encoding phase
			startTime := time.Now()

			if strings.Contains(strings.ToLower(compressor), "huff") {
				fmt.Println("Compressor chosen: Huffman")
				encodedImage := huffmanCompression(image)

				fmt.Println(len(encodedImage))
				fmt.Println("SizeOf of image: ", (unsafe.Sizeof([114976]byte{})))

				gain := float64(unsafe.Sizeof([266 * 425]uint32{})) / float64(unsafe.Sizeof([114976]byte{}))
				fmt.Println("Gain: ", gain)
			} else {
				fmt.Println("Compressor ", compressor, " does not match any existing compressor")
			}

			elapsedTime := time.Since(startTime)
			fmt.Println("Time for encoding: ", elapsedTime)
		} else {
			fmt.Println(status)
		}
	} else {
		fmt.Println("Missing parameters. shall be: compression imageFilePath algorithm")
		fmt.Println("Note: algorithm can be: \n\t- Huffman\n\t- TODO")
	}
}
