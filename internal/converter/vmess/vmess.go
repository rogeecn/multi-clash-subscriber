package vmess

import (
	"encoding/base64"
	"encoding/json"
	"multi-clash-subscriber/config"
	"net/url"
	"strconv"
)

type Vmess struct {
	Add  string      `json:"add"`
	Aid  interface{} `json:"aid"`
	Host string      `json:"host"`
	ID   string      `json:"id"`
	Net  string      `json:"net"`
	Path string      `json:"path"`
	Port string      `json:"port"`
	Ps   string      `json:"ps"`
	TLS  string      `json:"tls"`
	Type string      `json:"type"`
	V    string      `json:"v"`
}

func (v *Vmess) IntPort() int {
	portInt, _ := strconv.Atoi(v.Port)
	return portInt
}
func (v *Vmess) IntAid() int {
	switch (v.Aid).(type) {
	case string:
		portInt, _ := strconv.Atoi(v.Aid.(string))
		return portInt
	case int:
		portInt := int(v.Aid.(int))
		return portInt
	case float64:
		portInt := int(v.Aid.(float64))
		return portInt
	default:
		return 0
	}
}

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

	var data Vmess
	err = json.Unmarshal(host, &data)
	if err != nil {
		return config.Proxy{}, err
	}

	return config.Proxy{
		Name:    data.Ps,
		Type:    "vmess",
		Server:  data.Add,
		Port:    data.IntPort(),
		UUID:    data.ID,
		AlterId: data.IntAid(),
		Cipher:  "auto",
		UDP:     true,
	}, nil
}
