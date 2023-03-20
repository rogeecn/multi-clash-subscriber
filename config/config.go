package config

type Config struct {
	App struct {
		Port int
	}
	Sources []Source
}

type Source struct {
	Name         string
	SubscribeURL string
	IgnoreNodes  []string
	FilterChars  []string
}

type Clash struct {
	Port               int           `yaml:"port"`
	SocksPort          int           `yaml:"socks-port"`
	AllowLan           bool          `yaml:"allow-lan"`
	Mode               string        `yaml:"mode"`
	LogLevel           string        `yaml:"log-level"`
	ExternalController string        `yaml:"external-controller"`
	Experimental       Experimental  `yaml:"experimental"`
	DNS                DNS           `yaml:"dns"`
	ProxyGroups        []ProxyGroups `yaml:"proxy-groups"`
	Rules              []string      `yaml:"rules"`
	Proxies            []Proxy       `yaml:"proxies"`
}

type ProxyGroups struct {
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
	Enable       bool     `yaml:"enable"`
	Ipv6         bool     `yaml:"ipv6"`
	EnhancedMode string   `yaml:"enhanced-mode"`
	Nameserver   []string `yaml:"nameserver"`
	Fallback     []string `yaml:"fallback"`
}
type Experimental struct {
	IgnoreResolveFail bool `yaml:"ignore-resolve-fail"`
}
