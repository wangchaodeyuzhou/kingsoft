package ip

import (
	"git.shiyou.kingsoft.com/go/errors"
	"net"
	"net/http"
	"strings"
)

var ipv4Addr10, ipv4Addr172, ipv4Addr192 string

// 获取内网 ipv4 地址
func GetIPInternal() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, a := range addrs {
		if ipNet, ok := a.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				switch ipNet.IP.To4()[0] {
				case 10:
					ipv4Addr10 = ipNet.IP.String()
				case 172:
					ipv4Addr172 = ipNet.IP.To4().String()
				case 192:
					ipv4Addr192 = ipNet.IP.To4().String()
				}
			}
		}
	}

	if ipv4Addr10 != "" {
		return ipv4Addr10, nil
	}

	if ipv4Addr192 != "" {
		return ipv4Addr192, nil
	}

	if ipv4Addr172 != "" {
		return ipv4Addr172, nil
	}

	return "", errors.New("there is no network interface")
}

// GetSrcIP 通过 request 获得 源 ip
func GetSrcIP(r *http.Request) (string, error) {
	if r == nil {
		return "", errors.New("param illegal")
	}
	// x-forwarded-for 解决通过代理的请求， remoteaddr 并非源 ip
	// 参考 RFC 7239（Forwarded HTTP Extension）
	ip := r.Header.Get("x-forwarded-for")
	if len(ip) != 0 {
		arrIp := strings.Split(ip, ",")
		if len(arrIp) > 0 {
			return arrIp[0], nil
		}
	}
	index := strings.LastIndex(r.RemoteAddr, ":")
	return r.RemoteAddr[:index], nil
}
