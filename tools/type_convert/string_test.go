package type_convert

import (
	"encoding/json"
	"strconv"
	"testing"
)

func TestStringToInt(t *testing.T) {
	str := "99999999999"
	i, err := strconv.Atoi(str)
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("after convert: %d", i)
	}
}

func TestStringToUint(t *testing.T) {
	str := "99999999999"
	i, err := strconv.Atoi(str)
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("after convert: %d", uint(i))
	}
}

func TestStringToInt64(t *testing.T) {
	str := "123"
	i64, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("after convert: %d", i64)
	}
}

func TestStringToFloat64(t *testing.T) {
	str := "123.456"
	i64, err := strconv.ParseFloat(str, 10)
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("after convert: %f", i64)
	}
}

func TestStringToMap(t *testing.T) {
	str := `{"IP": "192.168.11.22", "name": "SKY"}`
	m := make(map[string]string)

	err := json.Unmarshal([]byte(str), &m)
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("after convert: %v", m)
	}

	for k, v := range m {
		t.Log(k, ":", v)
	}
}
