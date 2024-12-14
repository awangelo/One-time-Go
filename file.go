package main

import "os"

// FileConfig sao os parametros
type FileConfig struct {
	InputFile  string
	OutputFile string
	PadFile    string
}

// Checa se o input file foi fornecido
func (d FileConfig) isValid() bool {
	return d.InputFile == ""
}

// Checa se o arquivo de pad existe
func (d FileConfig) isEncrypted() bool {
	_, err := os.Stat(d.PadFile)
	return err == nil // retorna true se o arquivo existe (decrypt) ou false se nao existe (encrypt)
}

func (d FileConfig) readInputFile() ([]byte, error) {
	input, err := os.ReadFile(d.InputFile)
	if err != nil {
		return nil, err
	}
	return input, nil
}

func (d FileConfig) readPadFile() ([]byte, error) {
	pad, err := os.ReadFile(d.PadFile)
	if err != nil {
		return nil, err
	}
	return pad, nil
}

func (d FileConfig) writeEncryptedFile(data []byte) error {
	err := os.WriteFile(d.OutputFile, data, 0666)
	if err != nil {
		return err
	}
	return nil
}

func (d FileConfig) writePadFile(data []byte) error {
	err := os.WriteFile(d.PadFile, data, 0666)
	if err != nil {
		return err
	}
	return nil
}
