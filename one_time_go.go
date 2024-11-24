package main

import (
	"flag"
	"fmt"
)

var data Data

func init() {
	flag.StringVar(&data.InputFile, "input", "", "Input file (required)")
	flag.StringVar(&data.OutputFile, "output", "encrypted.out", "Encrypted output file name")
	flag.StringVar(&data.PadFile, "pad", "pad.out", "Pad file name")
	flag.StringVar(&data.InputFile, "i", "", "Alias for input")
	flag.StringVar(&data.OutputFile, "o", "encrypted.out", "Alias for output")
	flag.StringVar(&data.PadFile, "p", "pad.out", "Alias for pad")

	flag.Parse()
}

func main() {
	// Mostrar ajuda se nao for passado input
	if data.isValid() {
		flag.Usage()
		return
	}

	if data.isEncrypted() {
		err := decrypt()
		if err != nil {
			panic(err)
		}
		fmt.Println("File decrypted successfully: ", data.OutputFile)
	} else {
		err := encrypt()
		if err != nil {
			panic(err)
		}
		fmt.Printf("File encrypted successfully: %s, %s\n", data.OutputFile, data.PadFile)
	}
}
