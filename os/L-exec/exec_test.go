package L_exec

import (
	"fmt"
	"log"
	"os/exec"
	"testing"
)

// 用来执行系统命令的
// https://colobu.com/2017/06/19/advanced-command-execution-in-Go-with-os-exec/
func TestExec(t *testing.T) {
	cmd := exec.Command("ls", "-lah", "../")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	fmt.Printf("combined out:\n%s\n", string(out))
}
