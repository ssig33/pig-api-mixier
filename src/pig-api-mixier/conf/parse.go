package conf

import (
	"gopkg.in/yaml.v2"
	"log"
)

func parse() map[interface{}]interface{} {
	bs, recoverAssetError := Asset("conf.yaml")
	if recoverAssetError != nil {
		log.Fatal(recoverAssetError)
	}

	m := make(map[interface{}]interface{})
	parseError := yaml.Unmarshal([]byte(bs), &m)

	if parseError != nil {
		log.Fatalf("error: %v", parseError)
	}
	return m
}
