package main

import "os"

// Data sao os parametros
type Data struct {
	InputFile  string
	OutputFile string
	PadFile    string
}

func (d Data) isValid() bool {
	return d.InputFile != ""
}

func (d Data) isEncrypted() bool {
	return d.PadFile != ""
}

func (d Data) getInputContent() ([]byte, error) {
	input, err := os.ReadFile(d.InputFile)
	if err != nil {
		return nil, err
	}
	return input, nil
}

func (d Data) getPadContent() ([]byte, error) {
	pad, err := os.ReadFile(d.PadFile)
	if err != nil {
		return nil, err
	}
	return pad, nil
}
