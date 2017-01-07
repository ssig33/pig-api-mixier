package conf

import ()

type UrlInfo struct {
	Url   string
	Rule  []string
	Basic []string
}

func UrlBases() []string {
	config := parse()

	result := []string{}

	for _, ep := range config["external_pigs"].([]interface{}) {
		result = append(result, ep.(map[interface{}]interface{})["url_base"].(string))
	}
	return result
}

func urlBasesByCase(c string) []UrlInfo {
	config := parse()

	result := []UrlInfo{}

	for _, ep := range config["external_pigs"].([]interface{}) {
		pig := ep.(map[interface{}]interface{})
		slice := pig["method"].([]interface{})
		included := false
		for _, method := range slice {
			if !included {
				included = method.(string) == c
			}
		}
		if included {
			pig := ep.(map[interface{}]interface{})
			url := pig["url_base"].(string)
			rule := []string{}
			for _, r := range pig["rule"].([]interface{}) {
				rule = append(rule, r.(string))
			}
			basic := []string{}
			for _, b := range pig["basic"].([]interface{}) {
				basic = append(basic, b.(string))
			}
			result = append(result, UrlInfo{Url: url, Rule: rule, Basic: basic})
		}
	}
	return result
}

func LatestUrlBases() []UrlInfo {
	return urlBasesByCase("latest")
}

func SearchUrlBases() []UrlInfo {
	return urlBasesByCase("search")
}
