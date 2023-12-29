package network

import (
	"regexp"
	"testing"
)

func TestHidePrivacy(t *testing.T) {
	reg, err := regexp.Compile("(?U)access_token=(.*)&")
	if err != nil {
		t.Log(err)
		return
	}

	PrivacyReg = []*regexp.Regexp{reg}
	msg := `Get "https://pan.baidu.com/rest/2.0/xpan/file?access_token=121.d1f66e95acfa40274920079396a51c48.Y2aP2vQDq90hLBE3PAbVije59uTcn7GiWUfw8LCM_olw&dir=%2F&limit=200&method=list&order=name&start=0&web=web " : net/http: TLS handshake timeout`
	res := hidePrivacy(msg)
	t.Log(res)
}
