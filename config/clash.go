package config

import "gopkg.in/yaml.v2"

type Clash struct {
	Port               int          `yaml:"port"`
	SocksPort          int          `yaml:"socks-port,omitempty"`
	AllowLan           bool         `yaml:"allow-lan,omitempty"`
	Mode               string       `yaml:"mode"`
	LogLevel           string       `yaml:"log-level"`
	ExternalController string       `yaml:"external-controller,omitempty"`
	Experimental       Experimental `yaml:"experimental,omitempty"`
	DNS                DNS          `yaml:"dns,omitempty"`
	ProxyGroups        []ProxyGroup `yaml:"proxy-groups"`
	Rules              []string     `yaml:"rules"`
	Proxies            []Proxy      `yaml:"proxies"`
}

type ProxyGroup struct {
	Name     string   `yaml:"name"`
	Type     string   `yaml:"type"`
	Proxies  []string `yaml:"proxies"`
	URL      string   `yaml:"url"`
	Interval int      `yaml:"interval"`
}

type Proxy struct {
	Name     string `yaml:"name"`
	Type     string `yaml:"type"`
	Server   string `yaml:"server"`
	Port     int    `yaml:"port"`
	Cipher   string `yaml:"cipher"`
	Password string `yaml:"password"`
	UDP      bool   `yaml:"udp"`
}

type DNS struct {
	Enable       bool     `yaml:"enable,omitempty"`
	Ipv6         bool     `yaml:"ipv6,omitempty"`
	EnhancedMode string   `yaml:"enhanced-mode,omitempty"`
	Nameserver   []string `yaml:"nameserver,omitempty"`
	Fallback     []string `yaml:"fallback,omitempty"`
}
type Experimental struct {
	IgnoreResolveFail bool `yaml:"ignore-resolve-fail,omitempty"`
}

func (c *Clash) String() (string, error) {
	out, err := c.Bytes()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func (c *Clash) Bytes() ([]byte, error) {
	return yaml.Marshal(c)
}
