package L_time

import (
	"fmt"
	"testing"
	"time"
)

var (
	USA *time.Location
	err error
)

func init() {
	USA, err = time.LoadLocation("America/New_York")
	if err != nil {
		panic(err)
	}
}

func TestLocation(t *testing.T) {
	// time 包提供了 Location 的两个实例：Local 和 UTC
	// Local 代表当前系统本地时区
	// UTC 代表通用协调时间，也就是零时区
	// time 包默认（为显示提供时区）使用 UTC 时区
	fmt.Println(time.Local)
	fmt.Println(time.UTC)
	// Unix 系统以标准格式将时区信息存于文件中，这些文件位于 /usr/share/zoneinfo
	// 而本地时区可以通过 /etc/localtime 获取，这是一个符号链接，指向 /usr/share/zoneinfo 中某一个时区
	// 因此，在初始化 Local 时，通过读取 /etc/localtime 可以获取到系统本地时区。
	// 当然，如果设置了环境变量 TZ，则会优先使用它。
}

func TestLoadLocation(t *testing.T) {
	fmt.Println(USA)
}
