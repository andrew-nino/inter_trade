package processing

import "encoding/base64"

func encodingBase64(b []byte) string {

	return base64.StdEncoding.EncodeToString(b)
}
