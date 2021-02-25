package tm

import (
	"log"
	"strings"
	"time"
)

func GetNowTime() string {
	now := time.Now().String()

	// 2020-11-21 21:04:24.256822 +0800 CST m=+0.000653870 这个样子 切割一下
	s1 := strings.Split(now, ".")
	if len(s1) < 1 {
		log.Printf("split s1 failed, now: %s", now)
		return ""
	}

	s2 := strings.Split(s1[0], " ")
	if len(s2) != 2 {
		log.Printf("split s2 failed, now: %s", now)
		return ""
	}

	s3 := strings.Split(s2[0], "-")
	if len(s3) != 3 {
		log.Printf("split s3 failed, now: %s", now)
		return ""
	}

	return s3[0] + "年" + s3[1] + "月" + s3[2] + "日 " + s2[1]
}
