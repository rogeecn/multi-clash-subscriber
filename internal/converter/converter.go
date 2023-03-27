package converter

import (
	"errors"
	"multi-clash-subscriber/config"
	"multi-clash-subscriber/internal/converter/ss"
	"multi-clash-subscriber/internal/converter/vmess"
	"strings"
)

func FromString(str string) (config.Proxy, error) {
	if strings.HasPrefix(str, "vmess://") {
		return vmess.FromString(str)
	} else if strings.HasPrefix(str, "ss://") {
		return ss.FromString(str)
		// } else if strings.HasPrefix(str, "trojan://") {
		// 	return trojan.FromString(str)
		// } else if strings.HasPrefix(str, "ssr://") {
		// 	return ssr.FromString(str)
		// } else if strings.HasPrefix(str, "http://") || strings.HasPrefix(str, "https://") {
		// 	return http.FromString(str)
		// } else if strings.HasPrefix(str, "socks5://") {
		// 	return socks5.FromString(str)
		// } else if strings.HasPrefix(str, "custom://") {
		// 	return custom.FromString(str)
	}
	return config.Proxy{}, errors.New("no match converter")
}
