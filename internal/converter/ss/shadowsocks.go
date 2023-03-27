package ss

import (
	"encoding/base64"
	"multi-clash-subscriber/config"
	"net/url"
	"strconv"
	"strings"
)

func FromString(str string) (config.Proxy, error) {
	u, err := url.Parse(str)
	if err != nil {
		return config.Proxy{}, err
	}

	// base64 decode u.Host
	host, err := base64.RawURLEncoding.DecodeString(u.Host)
	if err != nil {
		return config.Proxy{}, err
	}

	// aes-256-gcm:eLZk38nCpogW2nHd@65.49.192.176:18322
	items := strings.Split(string(host), "@")
	if len(items) != 2 {
		return config.Proxy{}, err
	}
	// split items[0],item[1] by ":"
	items0 := strings.Split(items[0], ":")
	items1 := strings.Split(items[1], ":")

	items = append([]string{}, items0...)
	items = append(items, items1...)

	portInt, err := strconv.Atoi(items[3])
	if err != nil {
		return config.Proxy{}, err
	}
	return config.Proxy{
		Name:     u.Fragment,
		Type:     "ss",
		Server:   items[2],
		Port:     portInt,
		Cipher:   items[0],
		Password: items[1],
		UDP:      true,
	}, nil
}
