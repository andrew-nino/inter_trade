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

// We get a hash from a string taking into account the type of algorithm.
func Processing(stringToHash string, typeHash string) (string, error) {

	var err error
	var salt []byte
	var once sync.Once

	once.Do(func() {
		salt = generateRandomSalt(saltSize)
	})

	hashBytes, err := hashString(stringToHash, typeHash, salt)

	return encodingBase64(hashBytes), err
}

func hashString(inputString string, typeHash string, salt []byte) ([]byte, error) {

	var err error
	stringBytes := []byte(inputString)

	switch typeHash {

	case "sha256":
		sha256Hasher := sha256.New()
		stringBytes = append(stringBytes, salt...)
		_, err = sha256Hasher.Write(stringBytes)
		return sha256Hasher.Sum(nil), err

	case "md5":
		md5Hasher := md5.New()
		stringBytes = append(stringBytes, salt...)
		_, err = md5Hasher.Write(stringBytes)
		return md5Hasher.Sum(nil), err

	default:
		err = fmt.Errorf("algorithm type is not defined")
	}
	return stringBytes, err
}
