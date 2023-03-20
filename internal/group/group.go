package group

import "multi-clash-subscriber/config"

type Group struct {
	proxies []config.Proxy
	groups  config.Groups
}

func New(config *config.Config, proxies []config.Proxy) *Group {
	return &Group{proxies: proxies}
}

// Generate groups
func (g Group) Generate() {
}
