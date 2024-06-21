package processing

import "encoding/hex"

func encodingBase64(b []byte) string {

	return hex.EncodeToString(b)
}
