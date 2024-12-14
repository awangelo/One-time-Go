package main

import (
	"os"
)

// FileConfig sao os parametros
type FileConfig struct {
	InputFile  string
	OutputFile string
	PadFile    string
}

// Checa se o input file foi fornecido
func (fc FileConfig) isValid() bool {
	return fc.InputFile != ""
}

// Checa se o pad file eh valido
func (fc FileConfig) isPadFileValid() bool {
	pad, err := fc.readPadFile()
	if err != nil {
		return false
	}
	input, err := fc.readInputFile()
	if err != nil {
		return false
	}
	return len(pad) == len(input)
}

// Checa se o arquivo de pad existe
func (fc FileConfig) isEncrypted() bool {
	_, err := os.Stat(fc.PadFile)
	return err == nil // true se o arquivo existe (decrypt) ou false se nao existe (encrypt)
}

func (fc FileConfig) readInputFile() ([]byte, error) {
	input, err := os.ReadFile(fc.InputFile)
	if err != nil {
		return nil, err
	}
	return input, nil
}

func (fc FileConfig) readPadFile() ([]byte, error) {
	pad, err := os.ReadFile(fc.PadFile)
	if err != nil {
		return nil, err
	}
	return pad, nil
}

func (fc FileConfig) writeFile(data []byte, target string) error {
	err := os.WriteFile(target, data, 0666)
	if err != nil {
		return err
	}
	return nil
}
