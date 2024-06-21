package processing

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
)

func Processing(stringToHash string, typeHash string) (hash []byte, err error) {

	switch typeHash {

	case "sha256":
		h := sha256.New()
		h.Write([]byte(stringToHash))
		hash = h.Sum(nil)
		return hash, nil

	case "md5":
		h := md5.New()
		h.Write([]byte(stringToHash))
		hash = h.Sum(nil)
		return hash, nil

	default:
		err = fmt.Errorf("algorithm type is not defined")
	}
	return
}
