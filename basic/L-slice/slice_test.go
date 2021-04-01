package L_slice

import (
	"fmt"
	"testing"
)

// ⚠️ copy 是开辟新空间的
// ⚠️ append 到了cap上限是开辟新空间的 没到就不开 引用相同的底层数组
// copy函数将返回成功复制的元素的个数，等于两个slice中较小的长度，所以我们不用担心覆盖会超出目标slice的范围。

func TestSlice(t *testing.T) {
	// slice本质是array元素的引用，官方把slice叫做underlying array
	// 切片定义
	var s1 []int // nil
	s2 := []int{}
	s3 := make([]int, 0)
	s4 := []int(nil) // nil

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)

	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)
	fmt.Println(s3 == nil)
	fmt.Println(s4 == nil)

	// ⚠️ 如果你需要测试一个slice是否是空的，使用len(s) == 0来判断，而不应该用s == nil来判断
}

func TestSliceInit(t *testing.T) {
	// 二维slice 初始化
	var s [][]int

	for i := 0; i < 10; i++ {
		s = append(s, []int{})
		for j := 0; j < 8; j++ {
			s[i] = append(s[i], -1)
		}
	}

	fmt.Println(s)
}

func TestMake2Slice(t *testing.T) {
	// 第二个参数是slice长度，初始len和cap相同
	s := make([]int, 5)
	fmt.Println(len(s))
	fmt.Println(cap(s))

	// 如果在len==cap时 开始append len+1 cap*2
	s = append(s, 3)
	fmt.Println(len(s))
	fmt.Println(cap(s))
}

func TestMake3Slice(t *testing.T) {
	// 第二个参数是slice长度，第三个是容量
	s := make([]int, 5, 10)
	fmt.Println(len(s))
	fmt.Println(cap(s))
	fmt.Println(s)

	// 这里 如果直接s[5] 会报错

	// 往s后面添加五个元素
	// 在不超过容量的情况下 每添加一个 长度加一 容量不变
	s = append(append(append(append(append(s, 3), 3), 3), 3), 3)
	fmt.Println(len(s))
	fmt.Println(cap(s))
	fmt.Println(s)

	// 当长度超过容量的时候 容量乘2
	s = append(s, 3)
	fmt.Println(len(s))
	fmt.Println(cap(s))
}

func TestOverSlice(t *testing.T) {
	// 定义一个slice
	s := []int{2, 3, 5, 7, 11, 13}
	// 对slice切片 得到新slice [3,5,7]
	sli := s[1:4]
	fmt.Println(sli)
	// 此时的长度就是真实的3  cap沿用前一个切片的 去头不去尾
	fmt.Println(len(sli))
	fmt.Println(cap(sli))

	// slice扩容 最大可以扩到当前cap
	fmt.Println(sli[:cap(sli)])

	// 扩容超过cap 直接报错
	// fmt.Println(sli[:8])
}

func TestSliceRef(t *testing.T) {
	// 切片的直接赋值相当于传指针
	s1 := []int{1, 2, 3, 4}
	s2 := s1
	s2[0] = 10

	// 这种对下标赋值 改s2 s1会变 改s1 s2也会变
	fmt.Println(s1)

	// 使用copy才是传值 但是copy前dst要先make好
	s3 := []int{1, 2, 3, 4}
	s4 := make([]int, 4)
	copy(s4, s3)
	s4[0] = 10

	// 这种赋值 不管怎么改 都不影响另外一个
	fmt.Println(s3)

	// 这种 要注意了 不影响了
	s5 := []int{1, 2, 3, 4}
	s6 := s5[1:3]
	s5 = []int{10, 20, 30, 40} // 因为这里等号右边相当于new了一下

	fmt.Println(s6)
}

// 两个切片引用同一个数组时，修改一个 另外一个也会变
func TestRef(t *testing.T) {
	arr := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	a1 := arr[2:5]
	a2 := arr[3:6]
	fmt.Println(a1)
	fmt.Println(a2)

	a1[2] = 1000

	fmt.Println(a1)
	fmt.Println(a2)
}

func TestReverse(t *testing.T) {
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:])
	fmt.Println(a)
}

// 倒序
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
