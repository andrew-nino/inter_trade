package processing

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"sync"
)

const saltSize = 16

func generateRandomSalt(saltSize int) []byte {

	var salt = make([]byte, saltSize)

	_, err := rand.Read(salt[:])
	if err != nil {
		panic(err)
	}
	return salt
}

func Processing(stringToHash string, typeHash string) (string, string, error) {

	var err error
	var salt []byte
	var once sync.Once

	once.Do(func() {
		salt = generateRandomSalt(saltSize)
	})

	switch typeHash {

	case "sha256":
		return encodingBase64(hashPassword_sha256(stringToHash, salt)), typeHash, nil
	case "md5":
		return encodingBase64(hashPassword_MD5(stringToHash, salt)), typeHash, nil

	default:
		err = fmt.Errorf("algorithm type is not defined")
	}
	return "", "", err
}

func hashPassword_sha256(password string, salt []byte) []byte {

	passwordBytes := []byte(password)
	sha256Hasher := sha256.New()
	passwordBytes = append(passwordBytes, salt...)
	sha256Hasher.Write(passwordBytes)

	return sha256Hasher.Sum(nil)
}

func hashPassword_MD5(password string, salt []byte) []byte {

	passwordBytes := []byte(password)
	md5Hasher := md5.New()
	passwordBytes = append(passwordBytes, salt...)
	md5Hasher.Write(passwordBytes)

	return md5Hasher.Sum(nil)
}
