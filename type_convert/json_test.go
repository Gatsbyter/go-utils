package type_convert

import (
	"crypto/sha256"
	"encoding/base64"
	"testing"
)

func TestJSON(t *testing.T) {
	bytes := []byte{
		0x1,
		0x23,
		0x34,
	}
	str := base64.StdEncoding.EncodeToString(bytes)
	t.Logf("%s", str)

	hasher := sha256.New()
	hasher.Write(bytes)
	t.Logf("%x", hasher.Sum(nil))
}
