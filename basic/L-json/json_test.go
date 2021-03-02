package L_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSON(t *testing.T) {
	type JJJ struct {
		Year  int  `json:"released"`
		Color bool `json:"color,omitempty"` // omitempty: 这个参数为空或零值 就不显示
	}

	j := JJJ{
		Year:  2019,
		Color: false,
	}

	data, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))

	// 缩进展示
	dataI, err := json.MarshalIndent(j, "", "    ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(dataI))
}
