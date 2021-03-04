//给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
//给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
// https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/

package leecode

import (
	"github.com/Gatsbyter/go-utils/data_struct/stack"
	"testing"
)

func Test_9(t *testing.T) {
	t.Log(letterCombinations("23"))
}

var ref = map[byte][]byte{
	'2': {'a', 'b', 'c'},
	'3': {'d', 'e', 'f'},
	'4': {'g', 'h', 'i'},
	'5': {'j', 'k', 'l'},
	'6': {'m', 'n', 'o'},
	'7': {'p', 'q', 'r', 's'},
	'8': {'t', 'u', 'v'},
	'9': {'w', 'x', 'y', 'z'},
}

type OO struct {
	ch    byte
	index int
}

func Go(oo []interface{}) string {
	var res []byte
	for _, o := range oo {
		res = append(res, ref[o.(OO).ch][o.(OO).index])
	}

	return string(res)
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}

	var res []string
	arr := []byte(digits)
	length := len(arr)
	up := false

	s := stack.NewStack()
	s.Push(OO{
		ch:    arr[0],
		index: 0,
	})

	for s.Size() > 0 {
		if s.Size() == length {
			res = append(res, Go(s.ToSlice()))

			o := s.Pop().(OO)
			if o.index < len(ref[o.ch])-1 {
				s.Push(OO{
					ch:    o.ch,
					index: o.index + 1,
				})
			} else {
				up = true
			}

		} else {
			if up {
				o := s.Pop().(OO)
				if o.index < len(ref[o.ch])-1 {
					s.Push(OO{
						ch:    o.ch,
						index: o.index + 1,
					})

					up = false
				} else {
					up = true
				}
			} else {
				s.Push(OO{
					ch:    arr[s.Size()],
					index: 0,
				})
				up = false
			}
		}
	}

	return res
}
