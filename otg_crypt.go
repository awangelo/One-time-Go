package main

import (
	"crypto/rand"
	"fmt"
)

// encrypt criptografa os dados de entrada usando o metodo de criptografia One-time pad.
// Este metodo requer uma chave que seja tao longa (ou maior) quanto os dados de entrada.
// A funcao le o conteudo do arquivo, cria a chave (pad), realiza a criptografia e retorna um erro se ocorrer algum problema.
func (d Data) encrypt() error {
	input, err := d.readInputFile()
	if err != nil {
		return fmt.Errorf("Error reading input file: %w", err)
	}

	padLen := len(input) // padLen tem o tamanho do pad que sera usado
	pad := make([]byte, padLen, padLen)
	// verifica se o tamanho do pad esta correto
	if pad == nil || len(pad) != padLen {
		return fmt.Errorf("Error allocating pad size: %w", err)
	}
	// pad eh preenchido com valores verdadeiramente aleatorios
	_, err = rand.Read(pad)
	if err != nil {
		return fmt.Errorf("Error creating pad: %w", err)
	}

	// O arquivo encriptado e criado e preenchido com os dados criptografados
	encLen := padLen
	encrypted := make([]byte, encLen, encLen)
	// verifica se o tamanho do arquivo encriptado esta correto
	if encrypted == nil || len(encrypted) != encLen {
		return fmt.Errorf("Error allocating encrypted size: %w", err)
	}
	// O arquivo eh encriptado usando Vernam cipher (XOR)
	for i, v := range input {
		encrypted[i] = v ^ pad[i]
	}

	err = d.writeEncryptedFile(encrypted)
	if err != nil {
		return fmt.Errorf("Error writing encrypted file: %w", err)
	}

	return nil
}

// decrypt descriptografa os dados criptografados usando o metodo de criptografia One-time pad.
// Este metodo requer uma chave (pad) que seja tao longa (ou maior) quanto os dados criptografados.
// A funcao le o conteudo do arquivo da chave (pad), realiza a descriptografia e retorna um erro se ocorrer algum problema.
func (d Data) decrypt() error {
	pad, err := d.readPadFile()
	if err != nil {
		return fmt.Errorf("Error reading pad file: %w", err)
	}
}
