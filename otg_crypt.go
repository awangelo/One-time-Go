package main

import "fmt"

// encrypt criptografa os dados de entrada usando o metodo de criptografia One-time pad.
// Este metodo requer uma chave que seja tao longa (ou maior) quanto os dados de entrada.
// A funcao le o conteudo do arquivo, cria a chave (pad), realiza a criptografia e retorna um erro se ocorrer algum problema.
func (d Data) encrypt() error {
	input, err := d.getInputContent()
	if err != nil {
		return fmt.Errorf("Error reading input file: %w", err)
	}
}

// decrypt descriptografa os dados criptografados usando o metodo de criptografia One-time pad.
// Este metodo requer uma chave (pad) que seja tao longa (ou maior) quanto os dados criptografados.
// A funcao le o conteudo do arquivo da chave (pad), realiza a descriptografia e retorna um erro se ocorrer algum problema.
func (d Data) decrypt() error {
	pad, err := d.getPadContent()
	if err != nil {
		return fmt.Errorf("Error reading pad file: %w", err)
	}
}
