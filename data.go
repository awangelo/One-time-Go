package main

import "os"

// Data sao os parametros
type Data struct {
	InputFile  string
	OutputFile string
	PadFile    string
}

// Checa se o input file foi fornecido
func (d Data) isValid() bool {
	return d.InputFile == ""
}

// Checa se o arquivo de pad existe
func (d Data) isEncrypted() bool {
	_, err := os.Stat(d.PadFile)
	return err == nil // retorna true se o arquivo existe (decrypt) ou false se nao existe (encrypt)
}

func (d Data) readInputFile() ([]byte, error) {
	input, err := os.ReadFile(d.InputFile)
	if err != nil {
		return nil, err
	}
	return input, nil
}

func (d Data) readPadFile() ([]byte, error) {
	pad, err := os.ReadFile(d.PadFile)
	if err != nil {
		return nil, err
	}
	return pad, nil
}

func (d Data) writeEncryptedFile(data []byte) error {
	err := os.WriteFile(d.OutputFile, data, 0666)
	if err != nil {
		return err
	}
	return nil
}

func (d Data) writePadFile(data []byte) error {
	err := os.WriteFile(d.PadFile, data, 0666)
	if err != nil {
		return err
	}
	return nil
}
