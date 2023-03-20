package source

import (
	"log"
	"multi-clash-subscriber/config"
	"multi-clash-subscriber/utils"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/imroc/req/v3"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type Subscribe struct {
	subscribe *config.Subscribe
}

func New(subscribe *config.Subscribe) *Subscribe {
	return &Subscribe{subscribe: subscribe}
}

func (c *Subscribe) Parse() error {

	clash, err := c.download()
	if err != nil {
		return err
	}

	c.subscribe.Proxies = c.filterProxies(clash.Proxies)
	return nil
}

// Download the file from the url
func (c *Subscribe) download() (*config.Clash, error) {
	client := req.C()
	// client.DevMode()

	req, err := client.R().Get(c.subscribe.URL)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to download the file: %s", c.subscribe.URL)
	}

	userInfo := req.Header.Get("subscription-userinfo")
	if userInfo != "" {
		c.parseUserInfo(userInfo)
		if c.subscribe.UserInfo.Expire < int(time.Now().Unix()) {
			return nil, errors.New("expired subscribe: " + c.subscribe.Name)
		}
	}

	var result config.Clash
	err = yaml.Unmarshal(req.Bytes(), &result)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to unmarshal the config: %s", c.subscribe.URL)
	}

	return &result, nil

}

// parseUserInfo parse userinfo from header upload=4576337548; download=6582828335; total=107374182400; expire=1680492220
func (c *Subscribe) parseUserInfo(value string) {
	value = utils.FilterSpaces(value)
	value = strings.ReplaceAll(value, ";", "&")
	val, err := url.ParseQuery(value)
	if err != nil {
		log.Println("ERR: ", err)
		return
	}

	if item, ok := val["upload"]; ok && len(item) > 1 {
		c.subscribe.UserInfo.Upload, _ = strconv.Atoi(item[0])
	}

	if item, ok := val["download"]; ok && len(item) > 1 {
		c.subscribe.UserInfo.Download, _ = strconv.Atoi(item[0])
	}

	if item, ok := val["total"]; ok && len(item) > 1 {
		c.subscribe.UserInfo.Total, _ = strconv.Atoi(item[0])
	}

	if item, ok := val["expire"]; ok && len(item) > 1 {
		c.subscribe.UserInfo.Expire, _ = strconv.Atoi(item[0])
	}
}

// filter proxies by config rule
func (c *Subscribe) filterProxies(proxies []config.Proxy) []config.Proxy {
	newProxies := []config.Proxy{}
	for _, proxy := range proxies {
		for _, filterChar := range c.subscribe.FilterChars {
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
		for _, ignoreChars := range c.subscribe.IgnoreChars {
			if strings.Contains(proxy.Name, ignoreChars) {
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
