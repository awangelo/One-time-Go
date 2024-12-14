package main

import (
	"flag"
	"fmt"
	"os"
)

var file FileConfig

func init() {
	flag.StringVar(&file.InputFile, "input", "", "Input file (required)")
	flag.StringVar(&file.OutputFile, "output", "encrypted.out", "Encrypted output file name")
	flag.StringVar(&file.PadFile, "pad", "pad.out", "Pad file name")
	flag.StringVar(&file.InputFile, "i", "", "Alias for input (required)")
	flag.StringVar(&file.OutputFile, "o", "encrypted.out", "Alias for output")
	flag.StringVar(&file.PadFile, "p", "pad.out", "Alias for pad")

	flag.Parse()
}

func main() {
	// Mostrar ajuda se nao for passado input
	if file.isValid() {
		flag.Usage()
		return
	}

	if file.isEncrypted() {
		err := file.decrypt()
		if err != nil {
			fmt.Println("Error decrypting file: ", err)
			os.Exit(1)
		}
		fmt.Println("File decrypted successfully: ", file.OutputFile)
	} else {
		err := file.encrypt()
		if err != nil {
			fmt.Println("Error encrypting file: ", err)
			os.Exit(1)
		}
		fmt.Printf("File encrypted successfully: %s, %s\n", file.OutputFile, file.PadFile)
	}
}
