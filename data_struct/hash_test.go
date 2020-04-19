package data_struct

import (
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestHash(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 10; i++ {
		s := strconv.Itoa(rand.Intn(10000))
		hash := Hash(s)
		t.Logf("round %d : str=%s, hash=%s", i, s, hash)
	}
}
