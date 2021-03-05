package reflect

import (
	"fmt"
	"reflect"
)

type CPosition struct {
	Typ           string `json:"type"`
	Leverage      string `json:"leverage"`
	EnterPrice    string `json:"enterPrice"`
	Margin        string `json:"margin"`
	NotionalValue string `json:"notionalValue"`
	LiqPrice      string `json:"liqPrice"`
	Funding       string `json:"funding"`
	Pnl           string `json:"pnl"`
	CloseFee      string `json:"closeFee"`
	CloseType     string `json:"closeType"`
	LiqFee        string `json:"liqFee,omitempty"`
	TxHash        string `json:"txHash"`
}

func printTag(o interface{}) {
	t := reflect.TypeOf(o)

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Printf("%s: %s\n", f.Name, f.Tag.Get("json"))
	}

}
