package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// type Object struct {
// 	obj string `json:"obj"`
// }

// func HandleGET(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-type", "application/json;charset=UTF-8")
// 	w.WriteHeader(http.StatusOK)

// 	obj := Object{obj: "pouet"}
// 	json.NewEncoder(w).Encode(obj)
// }

func Hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello woorld!")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", Hello)
	http.Handle("/", r)
	fmt.Println("Starting up on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

	// if len(os.Args) >= 3 {
	// 	filePath := os.Args[1]
	// 	compressor := os.Args[2]

	// 	image := readImage(filePath)

	// 	// Measure algorithm encoding phase
	// 	startTime := time.Now()

	// 	if strings.Contains(strings.ToLower(compressor), "huff") {
	// 		fmt.Println("Compressor chosen: Huffman")
	// 		encodedImage := huffmanCompression(image)

	// 		//fmt.Println(encodedImage)
	// 		fmt.Println(len(encodedImage))
	// 		fmt.Println("SizeOfofimage: ", (unsafe.Sizeof([114976]byte{})))

	// 		gain := float64(unsafe.Sizeof([266 * 425]uint32{})) / float64(unsafe.Sizeof([114976]byte{}))
	// 		fmt.Println("Gain: ", gain)
	// 	} else {
	// 		fmt.Println("Compressor ", compressor, " does not match any existing compressor")
	// 	}

	// 	elapsedTime := time.Since(startTime)
	// 	fmt.Println("Time for encoding: ", elapsedTime)
	// } else {
	// 	fmt.Println("Missing parameters. shall be: compression imageFilePath algorithm")
	// 	fmt.Println("Note: algorithm can be: \n\t- Huffman\n\t- TODO")
	// }
}
