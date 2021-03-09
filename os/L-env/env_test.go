package L_env

import (
	"fmt"
	"os"
	"testing"
)

// 操作环境变量
func TestEnv(t *testing.T) {
	envs := os.Environ()

	for _, env := range envs {
		fmt.Println(env)
	}

	// 还有很多 用的时候看看 很简单
}
