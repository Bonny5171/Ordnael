package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"os"
)

var iv = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}
var segredo string = "mother::fucker::::::::::::::::::"

func baguncaOsBytes(texto, segredo string) ([]byte, error) {
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

func readline() string {
	bio := bufio.NewReader(os.Stdin)
	line, _, err := bio.ReadLine()
	if err != nil {
		fmt.Println(err)
	}
	return string(line)
}

func main() {
	fmt.Println(`1 para criptografar
ou
2 para descriptografar`)
	for {
		line := readline()

		switch line {
		case "1":
			fmt.Println("Me escreva algo para q eu possa foder com esses bytes!!!")
			line := readline()

			stringCriptografada, _ := baguncaOsBytes(line, segredo)
			str := string(stringCriptografada)
			fmt.Print(str, "\n\n\n")
		case "2":
			fmt.Println("me escreva algo para q eu possa rearranjar os bytes!!!")
			line := readline()
			stringDescriptografada, _ := reajeitarOsBytes(line, segredo)
			str2 := string(stringDescriptografada)
			fmt.Print(str2, "\n\n\n")
		}
	}
	// stringCriptografada, _ := baguncaOsBytes("Luis Fernando Machado Araujo", segredo)
	// str := string(stringCriptografada)
	// fmt.Print(str, "\n")

	// stringDescriptografada, _ := reajeitarOsBytes(str, segredo)
	// str2 := string(stringDescriptografada)
	// fmt.Print(str2, "\n")
}
