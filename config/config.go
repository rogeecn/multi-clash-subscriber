package config

import (
	"bytes"
	"os"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
)

type Config struct {
	App struct {
		Port int
		Rule string
	}
	Common struct {
		Port               int
		SocksPort          int
		AllowLan           bool
		Mode               string
		LogLevel           string
		ExternalController string
		Experimental       Experimental
	}
	Subscribes []Subscribe
	Groups     []Group
}

type Group struct {
	Name        string
	Type        string
	Proxies     []string
	TestURL     string
	Interval    int
	AppendNodes bool
}

type Subscribe struct {
	Name        string
	URL         string
	IgnoreChars []string
	FilterChars []string
	Proxies     []Proxy
	UserInfo    struct {
		Upload   int
		Download int
		Total    int
		Expire   int
	}
}

func (c *Config) Generate() (*Clash, error) {
	rules, err := os.ReadFile(c.App.Rule)
	if err != nil {
		return nil, errors.Wrap(err, "read rule failed")
	}
	lines := bytes.Split(rules, []byte("\n"))

	clash := &Clash{}
	if err := copier.Copy(clash, &c.Common); err != nil {
		return nil, errors.Wrap(err, "copy fields failed")
	}

	for _, s := range c.Subscribes {
		clash.Proxies = append(clash.Proxies, s.Proxies...)
	}

	proxyNames := []string{}
	for _, proxy := range clash.Proxies {
		proxyNames = append(proxyNames, proxy.Name)
	}

	for _, group := range c.Groups {
		proxyGroup := ProxyGroup{
			Name: group.Name,
			Type: group.Type,
		}

		if group.AppendNodes {
			proxyGroup.Proxies = append(group.Proxies, proxyNames...)
		}

		if group.Type == "url-test" {
			proxyGroup.URL = group.TestURL
			proxyGroup.Interval = group.Interval
		}

		clash.ProxyGroups = append(clash.ProxyGroups, proxyGroup)
	}

	for _, line := range lines {
		clash.Rules = append(clash.Rules, string(line))
	}

	return clash, nil
}
