package cfg

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Conf struct {
	Nodes map[string]string `yaml:"nodes"`
}

func (c *Conf) GetConf() (*Conf, error) {

	yamlFile, err := ioutil.ReadFile("conf.yaml")
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
