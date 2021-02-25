package tree

import "testing"

var dic = []string{
	"abc",
	"aaa",
	"add",
	"bcd",
	"dca",
	"aac",
	"ddc",
	"bca",
}

func TestBuildTree(t *testing.T) {
	root := BuildTree(dic)
	t.Logf("%v", Match(root, "efd"))
}
