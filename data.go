package main

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
