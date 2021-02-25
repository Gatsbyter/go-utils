package self_ip

import "testing"

func TestIP(t *testing.T) {
	ip, err := GetSelfIP()
	if err != nil {
		t.Fatalf("get self IP failed, err : %s", err.Error())
	}
	t.Logf("get self IP succed, IP: %s", ip)
}
