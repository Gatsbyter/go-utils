package http

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"strings"
)

// r传引用
func Get(url string, r interface{}) error {
	rsp, err := http.Get(url)
	if err != nil {
		return errors.Wrap(err, "http get req failed")
	}

	defer rsp.Body.Close()
	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return errors.Wrap(err, "read http rsp failed")
	}

	err = json.Unmarshal(body, r)
	if err != nil {
		return errors.Wrap(err, "decode http rsp failed")
	}

	return nil
}

// re post的请求参数 传值就可以
// rs 要传引用
func Post(url string, re, rs interface{}) error {
	data, err := json.Marshal(re)
	if err != nil {
		return errors.Wrap(err, "http post encode req failed")
	}

	rsp, err := http.Post(url, "application/json;charset=UTF-8", strings.NewReader(string(data)))
	if err != nil {
		return errors.Wrap(err, "http post req failed")
	}
	defer rsp.Body.Close()

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return errors.Wrap(err, "http post read rsp failed")
	}

	err = json.Unmarshal(body, rs)
	if err != nil {
		return errors.Wrap(err, "http post decode rsp failed")
	}

	return nil
}
