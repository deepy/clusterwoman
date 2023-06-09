package cfg

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io"
	"log"
	"os"
)

type Conf struct {
	Nodes map[string]string `yaml:"nodes"`
}

func (c *Conf) GetConfFile(filename string) (*Conf, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return c.GetConf(file)
}

func (c *Conf) GetConf(reader io.Reader) (*Conf, error) {
	yamlFile, err := io.ReadAll(reader)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
		return nil, err
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
		return nil, err
	}

	if hasDupes(c.Nodes) {
		return nil, errors.New("multiple nodes may not share the same MAC Address")
	}

	return c, nil
}

func hasDupes(m map[string]string) bool {
	x := make(map[string]struct{})

	for _, v := range m {
		if _, has := x[v]; has {
			return true
		}
		x[v] = struct{}{}
	}

	return false
}
