package network

import (
	"regexp"
	"strings"
)

var PrivacyReg []*regexp.Regexp

func hidePrivacy(msg string) string {
	for _, r := range PrivacyReg {
		msg = r.ReplaceAllStringFunc(msg, func(s string) string {
			return strings.Repeat("*", len(s))
		})
	}
	return msg
}
