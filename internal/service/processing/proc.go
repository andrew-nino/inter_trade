package processing

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
)

func Processing(stringToHash string, typeHash string) (string, string, error) {

	var err error

	switch typeHash {

	case "sha256":
		h := sha256.New()
		h.Write([]byte(stringToHash))

		return encodingBase64(h.Sum(nil)), typeHash, nil

	case "md5":
		h := md5.New()
		h.Write([]byte(stringToHash))
		return encodingBase64(h.Sum(nil)), typeHash, nil

	default:
		err = fmt.Errorf("algorithm type is not defined")
	}
	return "","", err
}
