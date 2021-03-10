package L_time

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	// Time类里面就三个 纳秒 秒 时区

	// time.Now 的具体实现在runtime里
	// 调用系统调用 clock_gettime 获取时钟值（这是 POSIX 时钟）
	// 其中 clockid_t 时钟类型是 CLOCK_REALTIME，也就是可设定的系统级实时时钟 （最多到纳秒）
	// 如果 clock_gettime 不存在，则使用精度差些的系统调用 gettimeofday（最多到微秒）
	t1 := time.Now() // ⚠️ time.Now 默认用的是Local时区
	fmt.Printf("date      --------  year: %d | month: %d | day: %d | hour: %d | minute: %d | second: %d \n",
		t1.Year(), t1.Month(), t1.Day(), t1.Hour(), t1.Minute(), t1.Second())
	fmt.Println("year day  -------- ", t1.YearDay())
	fmt.Println("local     -------- ", t1)
	fmt.Println("UTC       -------- ", t1.UTC())
	fmt.Println("USA       -------- ", t1.In(USA))
	fmt.Println("unix      -------- ", t1.Unix())
	fmt.Println("nano      -------- ", t1.UnixNano())
	fmt.Println("location  -------- ", t1.Location())
	fmt.Println("weekday   -------- ", t1.Weekday())
}

// 将任何形式的时间字符串 格式化成Time类型
func TestFormat(t *testing.T) {
	t1, err := time.Parse("2006 1 2 15-04-05", "2016 6 13 09-14-00")
	if err != nil {
		panic(err)
	}

	fmt.Println(t1)
	fmt.Println(t1.Format("2006/01/02 15/04/05"))

	// 这个还可以按时区
	t2, err := time.ParseInLocation("2006 1 2 15-04-05", "2016 6 13 09-14-00", time.Local)
	if err != nil {
		panic(err)
	}

	fmt.Println(t2)
}

// 取整
func TestTruncate(t *testing.T) {
	// 第一种方法
	t1, _ := time.ParseInLocation("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:00:00"), time.Local)
	fmt.Println(t1)

	// 第二种方法
	t2 := time.Now()
	// 整点（向下取整）
	fmt.Println(t2.Truncate(1 * time.Hour))
	// 整点（最接近）
	fmt.Println(t2.Round(1 * time.Hour))
	// 整分（向下取整）
	fmt.Println(t2.Truncate(1 * time.Minute))
	// 整分（最接近）
	fmt.Println(t2.Round(1 * time.Minute))
}
