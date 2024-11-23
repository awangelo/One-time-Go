package main

import (
	"flag"
)

func main() {
	// Argumentos longos
	inputFile := flag.String("input", "", "Input file (required)")
	outputFile := flag.String("output", "encrypted.out", "Encrypted output file name")
	padFile := flag.String("pad", "pad.out", "Pad file name")
	// Argumentos curtos
	flag.StringVar(inputFile, "i", "", "Alias for input")
	flag.StringVar(outputFile, "o", "encrypted.out", "Alias for output")
	flag.StringVar(padFile, "p", "pad.out", "Alias for pad")

	flag.Parse()

	// Mostrar ajuda se nao for passado input
	if *inputFile == "" {
		flag.Usage()
		return
	}

}
