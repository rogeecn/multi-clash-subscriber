package rule

import (
	"log"
	"multi-clash-subscriber/config"
	"os"
	"strings"

	"github.com/imroc/req/v3"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

func Generate(output string) error {
	_ = os.Remove(output)
	fd, err := os.Create(output)
	if err != nil {
		return errors.Wrapf(err, "create file %s failed", output)
	}

	for _, group := range config.C.Groups {
		log.Println("group: ", group.Name)
		for _, rule := range group.Rules {
			log.Println("> rule: ", rule)
			for s := range parse(rule, group.Name) {
				log.Println(s)
				_, _ = fd.WriteString(s + "\n")
			}
		}
	}

	_, _ = fd.WriteString("GEOIP,CN,直连" + "\n")
	_, _ = fd.WriteString("MATCH,漏网之鱼" + "\n")
	return nil
}

func parse(rule, name string) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		rules := strings.Split(rule, "|")
		if len(rules) != 3 {
			log.Println("err: wrong rule format, ", rule, ", got: ", len(rules))
			return
		}

		prefix, url, suffix := rules[0], rules[1], rules[2]

		data, err := download(url)
		if err != nil {
			log.Println("err: ", err)
			return
		}

		var payload payload
		err = yaml.Unmarshal(data, &payload)
		if err != nil {
			log.Println("err: yaml unmarshal failed, ", err)
			return
		}
		log.Println(">> payload size:", len(payload.Payload))

		for _, r := range payload.Payload {
			data := []string{prefix, r, name, suffix}
			result := strings.Trim(strings.Join(data, ","), ",")
			log.Println(">> GOT: ", result)
			ch <- result
		}
	}()
	return ch
}

func download(url string) ([]byte, error) {
	resp, err := req.C().R().Get(url)
	if err != nil {
		return nil, errors.Wrapf(err, "download %s failed", url)
	}

	return resp.Bytes(), nil
}

type payload struct {
	Payload []string `yaml:"payload"`
}
