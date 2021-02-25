package type_convert

import (
	"strconv"
	"testing"
)

func TestIntToString(t *testing.T) {
	i := 12345
	str := strconv.Itoa(i)
	t.Logf("after convert: %s", str)
}

func TestIntToUint(t *testing.T) {
	i := -12345
	ui := uint(i)
	t.Logf("after convert: %d", ui)
}
