package ip

import (
	"github.com/stretchr/testify/assert"
	"net"
	"testing"
)

func TestExtract(t *testing.T) {
	tests := []struct {
		addr   string
		expect string
	}{
		{"127.0.0.1:80", "127.0.0.1:80"},
		{"10.0.0.1:80", "10.0.0.1:80"},
		{"172.16.0.1:80", "172.16.0.1:80"},
		{"192.168.1.1:80", "192.168.1.1:80"},
		{"0.0.0.0:80", ""},
		{"[::]:80", ""},
		{":80", ""},
	}
	for _, test := range tests {
		t.Run(test.addr, func(t *testing.T) {
			res, err := Extract(test.addr, nil)
			if err != nil {
				t.Fatal(err)
			}
			if res != test.expect && (test.expect == "" && test.addr == test.expect) {
				t.Fatalf("expected %s got %s", test.expect, res)
			}
		})
	}
	lis, err := net.Listen("tcp", ":12345")
	assert.NoError(t, err)
	res, err := Extract("", lis)
	assert.NoError(t, err)
	expect, err := Extract(lis.Addr().String(), nil)
	assert.NoError(t, err)
	assert.Equal(t, expect, res)
}
