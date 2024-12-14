package main

import (
	"crypto/rand"
	"fmt"
)

// encrypt criptografa os dados de entrada usando o metodo de criptografia One-time pad.
// Este metodo requer uma chave que seja tao longa (ou maior) quanto os dados de entrada.
// A funcao le o conteudo do arquivo, cria a chave (pad), realiza a criptografia e retorna um erro se ocorrer algum problema.
func (d FileConfig) encrypt() error {
	input, err := d.readInputFile()
	if err != nil {
		return fmt.Errorf("error reading input file: %w", err)
	}

	padLen := len(input) // padLen tem o tamanho do pad que sera usado
	pad := make([]byte, padLen)
	// verifica se o tamanho do pad esta correto
	if len(pad) != padLen {
		return fmt.Errorf("error allocating pad size: %w", err)
	}
	// pad eh preenchido com valores verdadeiramente aleatorios
	_, err = rand.Read(pad)
	if err != nil {
		return fmt.Errorf("error creating pad: %w", err)
	}

	// O arquivo encriptado e criado e preenchido com os dados criptografados
	encLen := padLen
	encrypted := make([]byte, encLen)
	// verifica se o tamanho do arquivo encriptado esta correto
	if len(encrypted) != encLen {
		return fmt.Errorf("error allocating encrypted size: %w", err)
	}
	// O arquivo eh encriptado usando Vernam cipher (XOR)
	for i, v := range input {
		encrypted[i] = v ^ pad[i]
	}

	// Escreve o arquivo encriptados e o pad
	err = d.writeEncryptedFile(encrypted)
	if err != nil {
		return fmt.Errorf("error writing encrypted file: %w", err)
	}
	err = d.writePadFile(pad)
	if err != nil {
		return fmt.Errorf("error writing pad file: %w", err)
	}

	return nil
}

// decrypt descriptografa os dados criptografados usando o metodo de criptografia One-time pad.
// Este metodo requer uma chave (pad) que seja tao longa (ou maior) quanto os dados criptografados.
// A funcao le o conteudo do arquivo da chave (pad), realiza a descriptografia e retorna um erro se ocorrer algum problema.
func (d FileConfig) decrypt() error {
	_, err := d.readPadFile()
	if err != nil {
		return fmt.Errorf("error reading pad file: %w", err)
	}

	return nil
}
