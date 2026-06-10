package crypto

import (
	"bytes"
	"crypto/des"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"strings"
)

// salt y iteraciones idénticos a Encrypter_old.java (PBEWithMD5AndDES)
// Java bytes son signed: {-87,-101,-56,50,86,53,-29,3} == unsigned {0xA9,0x9B,0xC8,0x32,0x56,0x35,0xE3,0x03}
var salt = []byte{0xA9, 0x9B, 0xC8, 0x32, 0x56, 0x35, 0xE3, 0x03}

const iterations = 19

// Encrypter replica exactamente org.aston.services.Encrypter("104F") de Java.
type Encrypter struct {
	passPhrase []byte
}

func NewEncrypter(passPhrase string) *Encrypter {
	return &Encrypter{passPhrase: []byte(passPhrase)}
}

// Encrypt cifra con PBEWithMD5AndDES → Base64, igual que Java encrypt().
func (e *Encrypter) Encrypt(text string) string {
	if text == "" {
		return text
	}
	key, iv := e.deriveKeyAndIV()
	ciphertext, err := encryptDESCBC([]byte(text), key, iv)
	if err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(ciphertext)
}

// Decrypt descifra un string Base64 cifrado con PBEWithMD5AndDES, igual que Java decrypt().
func (e *Encrypter) Decrypt(text string) string {
	text = strings.ReplaceAll(text, " ", "")
	if text == "" {
		return text
	}
	decoded, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return ""
	}
	key, iv := e.deriveKeyAndIV()
	plaintext, err := decryptDESCBC(decoded, key, iv)
	if err != nil {
		return ""
	}
	return string(plaintext)
}

// deriveKeyAndIV implementa PKCS#5 PBKDF1 con MD5.
// T_1 = MD5(password || salt)
// T_i = MD5(T_{i-1})  para i = 2..iterations
// key = T_iterations[0:8], iv = T_iterations[8:16]
func (e *Encrypter) deriveKeyAndIV() (key, iv []byte) {
	d := append(e.passPhrase, salt...)
	var h [16]byte
	for i := 0; i < iterations; i++ {
		h = md5.Sum(d)
		d = h[:]
	}
	return d[0:8], d[8:16]
}

func encryptDESCBC(plaintext, key, iv []byte) ([]byte, error) {
	padded := pkcs5Pad(plaintext, des.BlockSize)
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	ciphertext := make([]byte, len(padded))
	cipher.NewCBCEncrypter(block, iv).CryptBlocks(ciphertext, padded)
	return ciphertext, nil
}

func decryptDESCBC(ciphertext, key, iv []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	plaintext := make([]byte, len(ciphertext))
	cipher.NewCBCDecrypter(block, iv).CryptBlocks(plaintext, ciphertext)
	return pkcs5Unpad(plaintext)
}

func pkcs5Pad(data []byte, blockSize int) []byte {
	padding := blockSize - (len(data) % blockSize)
	return append(data, bytes.Repeat([]byte{byte(padding)}, padding)...)
}

func pkcs5Unpad(data []byte) ([]byte, error) {
	if len(data) == 0 {
		return data, nil
	}
	padding := int(data[len(data)-1])
	return data[:len(data)-padding], nil
}
