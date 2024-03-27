package conf

import (
	"log"
	"os"

	"multi-clash-subscriber/config"
	"multi-clash-subscriber/internal/source"

	"golang.org/x/sync/errgroup"
)

func GetYaml() ([]byte, error) {
	var eg errgroup.Group
	for _, s := range config.C.Subscribes {
		s := s
		eg.Go(func() error {
			log.Println("downloading source ", s.URL)
			if err := source.New(s).Parse(); err != nil {
				log.Println("err: ", err)
				return err
			}
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		log.Println("get yaml has error: ", err)
	}

	clash, err := config.C.Generate()
	if err != nil {
		log.Println("err: ", err)
		return nil, err
	}
	return clash.Bytes()
}

func Generate(output string) error {
	b, err := GetYaml()
	if err != nil {
		return err
	}

	return os.WriteFile(output, b, 0o644)
}
