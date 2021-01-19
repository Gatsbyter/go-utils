package os

import (
	"path/filepath"
	"testing"
)

func TestDir(t *testing.T) {
	dir := filepath.Dir(filepath.Dir("/Users/gatsbyter/learning/go-utils/net/tcp"))
	t.Log(dir)
}