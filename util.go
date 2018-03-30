package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"io"
	"log"
	"strings"
	"time"
)

//Md5 return md5 hash string
func Md5(input string) string {
	hasher := md5.New()
	hasher.Write([]byte(input))
	return hex.EncodeToString(hasher.Sum(nil))
}

//DistinctUint return only unique uint values
func DistinctUint(input []uint) []uint {
	u := make([]uint, 0, len(input))
	m := make(map[uint]bool)

	for _, val := range input {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}
	return u
}

//DistinctInt64 return only unique uint values
func DistinctInt64(input []int64) []int64 {
	u := make([]int64, 0, len(input))
	m := make(map[int64]bool)

	for _, val := range input {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}
	return u
}

//DistinctString return only unique string values.
func DistinctString(input []string) []string {
	u := make([]string, 0, len(input))
	m := make(map[string]bool)

	for _, val := range input {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}
	return u
}

//DistinctLowerCase return only unique lowercase string values.
//For example input = ["Test1","tEst1"] result will be ["test1"]
func DistinctLowerCase(input []string) []string {
	u := make([]string, 0, len(input))
	m := make(map[string]bool)

	for _, val := range input {
		lowerCaseVal := strings.ToLower(val)
		if _, ok := m[lowerCaseVal]; !ok {
			m[lowerCaseVal] = true
			u = append(u, lowerCaseVal)
		}
	}
	return u
}

//ContainsString return true when string value contains in slice []string.
//In other case return false.
//Func is case sensetive
func ContainsString(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

//TimeTrack loggin excecution time
func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s,%s\n", name, elapsed)
}

//Encrypt plain text by specified key
func Encrypt(plaintext []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, plaintext, nil), nil
}

//Decrypt ciphertext by specified key
func Decrypt(ciphertext []byte, key []byte) (plaintext []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < gcm.NonceSize() {
		return nil, errors.New("malformed ciphertext")
	}

	return gcm.Open(nil,
		ciphertext[:gcm.NonceSize()],
		ciphertext[gcm.NonceSize():],
		nil,
	)
}

//EncryptBase64 encrypt plain text by specified key and encode to base64 string
func EncryptBase64(plaintext []byte, key []byte) (string, error) {

	ciphertext, err := Encrypt(plaintext, key)
	if err != nil {
		return "", err
	}

	// convert to base64
	return base64.RawURLEncoding.EncodeToString(ciphertext), nil
}

//DecryptBase64 decrypt by key base64 encoded ciphertext to plain text
func DecryptBase64(base64string string, key []byte) ([]byte, error) {
	ciphertext, _ := base64.RawURLEncoding.DecodeString(base64string)
	return Decrypt(ciphertext, key)
}