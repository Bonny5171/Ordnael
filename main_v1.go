package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	// "io/ioutil"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

func varrediretorio(root string) {
	var files []string
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	for _, file := range files {
		fmt.Println(file)
	}
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

	i := flag.Int("age", -1, "your age")
	n := flag.String("name", "", "your first name")
	b := flag.Bool("married", false, "are you married?")

	flag.Parse()

	fmt.Println("Name: ", *n)
	fmt.Println("Age: ", *i)
	fmt.Println("Married: ", *b)

	// log.Println(`
	// .
	// Ola! eu sou Ordnael
	// .
	// ¯\_( ͡° ͜ʖ ͡°)_/¯`)

	// for {
	// 	line := readline()

	// 	switch line {
	// 	case "encryp":
	// 		fmt.Println("Me escreva algo para q eu possa foder com esses bytes!!!")
	// 		line := readline()

	// 		key := []byte("a very very very very secret key") // 32 bytes
	// 		plaintext := []byte(line)
	// 		ciphertext := encrypt(key, plaintext)
	// 		fmt.Printf("%x\n", ciphertext)
	// 		result := decrypt(key, ciphertext)
	// 		fmt.Printf("%s\n", string(result))
	// 	case "decryp":
	// 		fmt.Println("me escreva algo para q eu possa foder com esses bytes!!!")
	// 		line := readline()
	// 		ciphertext := []byte(line)

	// 		key := []byte("a very very very very secret key") // 32 bytes
	// 		fmt.Printf("%x\n", line)
	// 		result := decrypt(key, ciphertext)
	// 		fmt.Printf("%s\n", string(result))
	// 	}
	// }
}

// See recommended IV creation from ciphertext below
//var iv = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

func encodeBase64(b []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(b))
}

func decodeBase64(b []byte) []byte {
	data, err := base64.StdEncoding.DecodeString(string(b))
	if err != nil {
		panic(err)
	}
	return data
}

func encrypt(key, text []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	b := encodeBase64(text)
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], b)
	return ciphertext
}

func decrypt(key, text []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	if len(text) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	return decodeBase64(text)
}
