package self_ip

import (
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetSelfIP() (string, error) {
	res, err := http.Get("https://ipinfo.io/ip")
	if err != nil {
		return "", errors.Wrap(err, "get ipinfo.io/ip failed")
	}

	defer res.Body.Close()
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", errors.Wrap(err, "read from res body failed")
	}

	return strings.Replace(string(content), "\n", "", -1), nil
}
