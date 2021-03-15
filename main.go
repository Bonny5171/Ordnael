package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

var iv = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}
var segredo string = "mother::fucker::"

func baguncaOsBytes(texto, segredo string)å ([]byte, error) {
	block, _ := aes.NewCipher([]byte(segredo))
	textoEmBytes := []byte(texto)
	cfb := cipher.NewCFBEncrypter(block, iv)
	cipherText := make([]byte, len(textoEmBytes))
	cfb.XORKeyStream(cipherText, textoEmBytes)
	return cipherText, nil
}

func reajeitarOsBytes(texto, segredo string) ([]byte, error) {
	block, _ := aes.NewCipher([]byte(segredo))
	cipherText := []byte(texto)
	cfb := cipher.NewCFBDecrypter(block, iv)
	planText := make([]byte, len(cipherText))
	cfb.XORKeyStream(planText, cipherText)
	return planText, nil
}

func main() {
	stringCriptografada, _ := baguncaOsBytes("Luis Fernando Machado Araujo", segredo)
	str := string(stringCriptografada)
	fmt.Print(str, "\n")

	stringDescriptografada, _ := reajeitarOsBytes(str, segredo)
	str2 := string(stringDescriptografada)
	fmt.Print(str2, "\n")

}
