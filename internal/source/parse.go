package source

import (
	"multi-clash-subscriber/config"
	"multi-clash-subscriber/utils"
	"strings"

	"github.com/imroc/req/v3"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type Source struct {
	source *config.Source
}

func New(source *config.Source) *Source {
	return &Source{source: source}
}

func (c Source) Proxies() ([]config.Proxy, error) {

	clash, err := c.download()
	if err != nil {
		return nil, err
	}

	return c.filterProxies(clash.Proxies), nil
}

// Download the file from the url
func (c Source) download() (*config.Clash, error) {
	client := req.C()
	// client.DevMode()

	req, err := client.R().Get(c.source.SubscribeURL)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to download the file: %s", c.source.SubscribeURL)
	}

	var result config.Clash
	err = yaml.Unmarshal(req.Bytes(), &result)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to unmarshal the config: %s", c.source.SubscribeURL)
	}

	return &result, nil

}

// filter proxies by config rule
func (c Source) filterProxies(proxies []config.Proxy) []config.Proxy {

	newProxies := []config.Proxy{}
	for _, proxy := range proxies {
		for _, filterChar := range c.source.FilterChars {
			switch filterChar {
			case "{emoji}":
				proxy.Name = utils.FilterEmojis(proxy.Name)
			case "{space}":
				proxy.Name = utils.FilterSpaces(proxy.Name)
			default:
				proxy.Name = strings.ReplaceAll(proxy.Name, filterChar, "")
			}
		}

		match := false
		for _, ignoreNode := range c.source.IgnoreNodes {
			if proxy.Name == ignoreNode {
				match = true
				break
			}
		}

		if !match {
			newProxies = append(newProxies, proxy)
		}

	}
	return newProxies
}
